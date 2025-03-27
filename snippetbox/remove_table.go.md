package main

import (
	"crypto/tls"
	"crypto/x509"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func removeTable(db *sql.DB) error {
	// Câu lệnh SQL để xóa bảng snippets
	dropTableSQL := `DROP TABLE IF EXISTS snippets;`
	_, err := db.Exec(dropTableSQL)
	return err

}

func main() {
	// Tải biến môi trường từ file .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Lấy các biến môi trường
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbSSLCA := os.Getenv("DB_SSL_CA")

	// Tạo chuỗi kết nối DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)

	// Tải chứng chỉ CA từ tệp
	caCert, err := os.ReadFile(dbSSLCA)
	if err != nil {
		log.Fatal("Unable to read CA file:", err)
	}

	// Cấu hình TLS
	tlsConfig := &tls.Config{
		InsecureSkipVerify: false, // Đặt thành false để đảm bảo kết nối an toàn
		RootCAs:            x509.NewCertPool(),
	}
	tlsConfig.RootCAs.AppendCertsFromPEM(caCert)

	// Đăng ký cấu hình TLS với MySQL driver
	mysql.RegisterTLSConfig("custom", tlsConfig)
	dsn = fmt.Sprintf("%s&tls=custom", dsn)

	// Mở kết nối đến cơ sở dữ liệu
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Kiểm tra kết nối
	if err = db.Ping(); err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	// Gọi hàm xóa bảng
	err = removeTable(db)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Table snippets removed successfully!")
}
