package main

import (
	"crypto/tls"
	"crypto/x509"
	"database/sql"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/go-playground/form/v4"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"snippetbox.alexedwards.net/internal/models"

	"github.com/alexedwards/scs/mysqlstore"
	"github.com/alexedwards/scs/v2"
)

type application struct {
	errorLog       *log.Logger
	infoLog        *log.Logger
	db             *sql.DB
	snippets       *models.SnippetModel
	templateCache  map[string]*template.Template
	formDecoder    *form.Decoder
	sessionManager *scs.SessionManager
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Lỗi khi load file .env")
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbSSLCA := os.Getenv("DB_SSL_CA")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := openDB(dsn, dbSSLCA)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Kiểm tra kết nối đến database
	err = db.Ping()
	if err != nil {
		log.Fatal("Không thể kết nối đến database: ", err)
	}

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// In thông báo xác nhận kết nối thành công
	infoLog.Println("Đã kết nối thành công đến database!")

	// Xóa bảng snippets nếu muốn (tùy chọn)
	// err = removeTable(db)
	// if err != nil {
	// 	errorLog.Fatal("Lỗi khi xóa bảng snippets: ", err)
	// }
	// infoLog.Println("Bảng snippets đã được xóa thành công hoặc không tồn tại!")

	// Tạo bảng snippets nếu chưa tồn tại
	// err = createTable(db)
	// if err != nil {
	// 	errorLog.Fatal("Lỗi khi tạo bảng snippets: ", err)
	// }
	// infoLog.Println("Bảng snippets đã được tạo thành công hoặc đã tồn tại!")

	// Initialize a new template cache...
	templateCache, err := newTemplateCache()
	if err != nil {
		errorLog.Fatal(err)
	}

	// Initialize a decoder instance...
	formDecoder := form.NewDecoder()

	sessionManager := scs.New()
	sessionManager.Store = mysqlstore.New(db)
	sessionManager.Lifetime = 12 * time.Hour

	app := &application{
		errorLog:       errorLog,
		infoLog:        infoLog,
		db:             db,
		snippets:       &models.SnippetModel{DB: db},
		templateCache:  templateCache,
		formDecoder:    formDecoder,
		sessionManager: sessionManager,
	}

	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Println("Starting server on :4000")
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

func openDB(dsn string, dbSSLCA string) (*sql.DB, error) {
	caCert, err := os.ReadFile(dbSSLCA)
	if err != nil {
		return nil, fmt.Errorf("unable to read CA file: %v", err)
	}

	tlsConfig := &tls.Config{
		RootCAs: x509.NewCertPool(),
	}

	ok := tlsConfig.RootCAs.AppendCertsFromPEM(caCert)
	if !ok {
		return nil, fmt.Errorf("failed to append CA cert to the root pool")
	}

	err = mysql.RegisterTLSConfig("custom", tlsConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to register TLS config: %v", err)
	}

	if strings.Contains(dsn, "?") {
		dsn = fmt.Sprintf("%s&tls=custom", dsn)
	} else {
		dsn = fmt.Sprintf("%s?tls=custom", dsn)
	}

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

// // Hàm tạo bảng snippets
// func createTable(db *sql.DB) error {
// 	createTableSQL := `
//     CREATE TABLE IF NOT EXISTS snippets (
//         id INT AUTO_INCREMENT PRIMARY KEY,
//         title VARCHAR(255) NOT NULL,
//         content TEXT NOT NULL,
//         created DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
//         expires DATETIME NOT NULL
//     );`
// 	_, err := db.Exec(createTableSQL)
// 	return err
// }

// // Hàm xóa bảng snippets
// func removeTable(db *sql.DB) error {
// 	dropTableSQL := `DROP TABLE IF EXISTS snippets;`
// 	_, err := db.Exec(dropTableSQL)
// 	return err
// }
