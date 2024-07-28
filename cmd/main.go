package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"

	"example.com/go-practicing/cmd/api"
	configs "example.com/go-practicing/cmd/config"
	"example.com/go-practicing/cmd/db"

)

func main() {
	cfg := mysql.Config{
		User:                 configs.Envs.DBUser,
		Passwd:               configs.Envs.DBPassword,
		Addr:                 configs.Envs.DBAddress,
		DBName:               configs.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	//creating a connection to database-SQL Container
	db, err := db.NewMySQLStorage(cfg)
	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)
	fmt.Println("the database is this:", db)

	//creating and running the server
	server := api.NewAPIServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

// to run the db create by NewSQLStorage
func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB Connected Successfully !!")
}

// package main

// import (
// 	"database/sql"
// 	"fmt"
// 	_ "github.com/go-sql-driver/mysql"
// 	"log"
// )

// func initStorage() (*sql.DB, error) {
// 	dsn := "root:my-secret-pwd@tcp(127.0.0.1:3306)/goPracticing?checkConnLiveness=false&parseTime=true&maxAllowedPacket=0"
// 	db, err := sql.Open("mysql", dsn)
// 	if err != nil {
// 		return nil, fmt.Errorf("error opening database: %w", err)
// 	}

// 	// Check if the database is alive
// 	if err := db.Ping(); err != nil {
// 		return nil, fmt.Errorf("error pinging database: %w", err)
// 	}

// 	log.Println("DB Connected Successfully !!")
// 	return db, nil
// }

// func main() {
// 	db, err := initStorage()
// 	if err != nil {
// 		log.Fatalf("Failed to connect to the database: %v", err)
// 	}
// 	defer db.Close()

// 	log.Println("The database is this:", db)

// 	// Perform a simple query to verify connection and data retrieval
// 	var version string
// 	err = db.QueryRow("SELECT VERSION()").Scan(&version)
// 	if err != nil {
// 		log.Fatalf("Failed to query database version: %v", err)
// 	}
// 	log.Println("Database version:", version)

// 	// Start your server or other application logic
// 	log.Println("Listening on :8080")
// 	// Example: http.ListenAndServe(":8080", nil)
// }
