package main

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"

	"example.com/go-practicing/cmd/api"
	"example.com/go-practicing/cmd/config"
	"example.com/go-practicing/cmd/db"

)

func main() {
	//creating a connection to database-SQL Container
	db, err := db.NewSQLStorage(&mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)

	//creating and running the server
	server := api.NewAPIServer(":8080", nil)
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
