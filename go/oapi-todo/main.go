package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/dys-org/oapi-todo/api"
	_ "modernc.org/sqlite"
)

var db *sql.DB

func initDB() {
	var err error
	db, err = sql.Open("sqlite", "./db_data/oapi-todo.db")
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}
	fmt.Println("Connected!")

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS todo (
			id      INTEGER PRIMARY KEY AUTOINCREMENT,
			text    TEXT NOT NULL,
			done    INTEGER NOT NULL DEFAULT 0
	)`)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}

}

func main() {
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
	e.Logger.Fatal(e.Start(":6970"))
}
