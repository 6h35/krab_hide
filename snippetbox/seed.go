package main

import (
	"crypto/tls"
	"crypto/x509"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	// Đọc cờ dòng lệnh
	// dropTable := flag.Bool("drop", false, "Drop the snippets table before seeding")
	// flag.Parse()

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

	err = db.Ping()
	if err != nil {
		log.Fatal("Không thể kết nối đến database: ", err)
	}

	fmt.Println("Đã kết nối thành công đến database!")

	// Tạo bảng nếu chưa tồn tại
	// err = createTable(db)
	// if err != nil {
	// 	log.Fatal("Lỗi khi tạo bảng snippets: ", err)
	// }
	// fmt.Println("Bảng snippets đã được tạo thành công hoặc đã tồn tại!")

	// if *dropTable {
	// 	err = removeTable(db)
	// 	if err != nil {
	// 		log.Fatal("Lỗi khi xóa bảng snippets: ", err)
	// 	}
	// 	fmt.Println("Bảng snippets đã được xóa!")
	// }

	// Chèn dữ liệu mẫu
	// err = insertSampleData(db)
	// if err != nil {
	// 	log.Fatal("Lỗi khi chèn dữ liệu mẫu: ", err)
	// }
	// fmt.Println("Dữ liệu mẫu đã được chèn thành công!")

	// Gọi hàm tạo bảng
	if err := createSessionsTable(db); err != nil {
		log.Fatal(err)
	}

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

// func removeTable(db *sql.DB) error {
// 	dropTableSQL := `DROP TABLE IF EXISTS snippets;`
// 	_, err := db.Exec(dropTableSQL)
// 	return err
// }

// func insertSampleData(db *sql.DB) error {
// 	// Kiểm tra xem bảng đã có dữ liệu chưa
// 	var count int
// 	err := db.QueryRow("SELECT COUNT(*) FROM snippets").Scan(&count)
// 	if err != nil {
// 		return err
// 	}
// 	if count > 0 {
// 		fmt.Println("Bảng snippets đã có dữ liệu, bỏ qua chèn dữ liệu mẫu.")
// 		return nil
// 	}

// 	// Dữ liệu mẫu
// 	insertSQL := `
//     INSERT INTO snippets (title, content, expires) VALUES
//         ('Morning Thoughts', 'A beautiful sunrise inspires new ideas every day.', DATE_ADD(UTC_TIMESTAMP(), INTERVAL 7 DAY)),
//         ('Project Notes', 'Meeting notes for the upcoming sprint planning session.', DATE_ADD(UTC_TIMESTAMP(), INTERVAL 14 DAY)),
//         ('Poetry Draft', 'Roses are red, violets are blue, coding is fun, and so are you!', DATE_ADD(UTC_TIMESTAMP(), INTERVAL 30 DAY)),
//         ('Todo List', '1. Finish coding\n2. Test application\n3. Deploy to server', DATE_ADD(UTC_TIMESTAMP(), INTERVAL 3 DAY)),
//         ('Random Idea', 'What if we could automate this process with AI?', DATE_ADD(UTC_TIMESTAMP(), INTERVAL 60 DAY));
//     `
// 	_, err = db.Exec(insertSQL)
// 	return err
// }

// Câu lệnh tạo bảng sessions
func createSessionsTable(db *sql.DB) error {
	// Tạo bảng sessions
	createTableQuery := `CREATE TABLE IF NOT EXISTS sessions (
		token CHAR(43) PRIMARY KEY,
		data BLOB NOT NULL,
		expiry TIMESTAMP(6) NOT NULL
	);`
	_, err := db.Exec(createTableQuery)
	if err != nil {
		return fmt.Errorf(" Lỗi khi tạo bảng sessions: %v", err)
	}

	// Kiểm tra xem index đã tồn tại chưa
	var indexName string
	err = db.QueryRow("SHOW INDEX FROM sessions WHERE Key_name = 'sessions_expiry_idx'").Scan(&indexName)
	if err != nil {
		// Nếu index chưa tồn tại, tạo mới
		if err == sql.ErrNoRows {
			_, err = db.Exec("CREATE INDEX sessions_expiry_idx ON sessions (expiry);")
			if err != nil {
				return fmt.Errorf(" Lỗi khi tạo index sessions_expiry_idx: %v", err)
			}
		} else {
			return fmt.Errorf(" Lỗi khi kiểm tra index: %v", err)
		}
	}

	fmt.Println("Bảng 'sessions' đã được tạo")
	return nil
}
