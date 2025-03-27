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

func createTable(db *sql.DB) error {
	createTableSQL := `
    CREATE TABLE IF NOT EXISTS snippets (
        id INT AUTO_INCREMENT PRIMARY KEY,
        title VARCHAR(255) NOT NULL,
        content TEXT NOT NULL,
        created DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
        expires DATETIME NOT NULL
    );`
	_, err := db.Exec(createTableSQL)
	return err
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbSSLCA := os.Getenv("DB_SSL_CA")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)

	caCert, err := os.ReadFile(dbSSLCA)
	if err != nil {
		log.Fatal("Unable to read CA file:", err)
	}

	tlsConfig := &tls.Config{
		InsecureSkipVerify: false,
		RootCAs:            x509.NewCertPool(),
	}
	tlsConfig.RootCAs.AppendCertsFromPEM(caCert)

	mysql.RegisterTLSConfig("custom", tlsConfig)
	dsn = fmt.Sprintf("%s&tls=custom", dsn)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	err = createTable(db)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Table created successfully!")
}
