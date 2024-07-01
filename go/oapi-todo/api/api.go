package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

var db *sql.DB

// SetDB sets the database instance for the api package
func SetDB(database *sql.DB) {
	db = database
}

type Server struct{}

func NewServer() Server {
	return Server{}
}

// (GET /todos)
func (Server) GetTodos(ctx echo.Context) error {
	var todos []Todo
	rows, err := db.Query("SELECT * FROM todo")
	if err != nil {
		return fmt.Errorf("could not query todo: %w", err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var td Todo
		if err := rows.Scan(&td.Id, &td.Text, &td.Done); err != nil {
			return fmt.Errorf("error scanning todo column values: %v", err)
		}
		todos = append(todos, td)
	}
	if err := rows.Err(); err != nil {
		return fmt.Errorf("GetTodos: %v", err)
	}

	return ctx.JSON(http.StatusOK, todos)
}

// (GET /todos/:id)
func (Server) GetTodosId(ctx echo.Context, id int) error {
	var td Todo
	row := db.QueryRow("SELECT * FROM todo WHERE id = ?", id)
	if err := row.Scan(&td.Id, &td.Text, &td.Done); err != nil {
		if err == sql.ErrNoRows {
			return ctx.NoContent(http.StatusNotFound)
		}
		return fmt.Errorf("could not query todo: %w", err)
	}
	return ctx.JSON(http.StatusOK, td)
}

// (POST /todos)
func (Server) PostTodos(ctx echo.Context) error {
	var td Todo
	if err := ctx.Bind(&td); err != nil {
		return fmt.Errorf("error binding todo: %w", err)
	}
	result, err := db.Exec("INSERT INTO todo (text, done) VALUES (?, ?)", td.Text, td.Done)
	if err != nil {
		return fmt.Errorf("failed to insert todo: %w", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("could not get last insert id: %w", err)
	}
	td.Id = int(id)
	return ctx.JSON(http.StatusCreated, td)
}

// (PUT /todos/:id)
func (Server) PutTodosId(ctx echo.Context, id int) error {
	var td Todo
	if err := ctx.Bind(&td); err != nil {
		return fmt.Errorf("error binding todo: %w", err)
	}
	_, err := db.Exec("UPDATE todo SET text = ?, done = ? WHERE id = ?", td.Text, td.Done, id)
	if err != nil {
		return fmt.Errorf("failed to update todo: %w", err)
	}
	return ctx.JSON(http.StatusOK, td)
}

// (DELETE /todos/:id)
func (Server) DeleteTodosId(ctx echo.Context, id int) error {
	_, err := db.Exec("DELETE FROM todo WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("failed to delete todo: %w", err)
	}
	return ctx.NoContent(http.StatusNoContent)
}
