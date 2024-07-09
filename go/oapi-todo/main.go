package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/dys-org/oapi-todo/api"
)

var db *sql.DB

func initDB() {
	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "oapi_todo",
	}

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected!")

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS todo (
			id      INT AUTO_INCREMENT NOT NULL,
			text    VARCHAR(255) NOT NULL,
			done    BOOLEAN NOT NULL DEFAULT false,
			PRIMARY KEY (id)
	)`)
	if err != nil {
		log.Fatal(err)
	}

}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	initDB()

	// Pass the db instance to the api package
	api.SetDB(db)

	// create a type that satisfies the `api.ServerInterface`, which contains an implementation of every operation from the generated code
	server := api.NewServer()
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	api.RegisterHandlersWithBaseURL(e, server, "/api")

	// And we serve HTTP until the world ends.
	e.Logger.Fatal(e.Start(":6969"))
}
