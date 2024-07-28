package api

import (
	"database/sql"
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
	rows, err := db.Query("SELECT * FROM todo")
	if err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "Failed to query todos"}

	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	var todos []Todo
	for rows.Next() {
		var td Todo
		if err := rows.Scan(&td.Id, &td.Text, &td.Done); err != nil {
			return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "Failed to scan todo"}
		}
		todos = append(todos, td)
	}

	return ctx.JSON(http.StatusOK, todos)
}

// (GET /todos/:id)
func (Server) GetTodosId(ctx echo.Context, id int) error {
	var td Todo
	row := db.QueryRow("SELECT * FROM todo WHERE id = ?", id)
	if err := row.Scan(&td.Id, &td.Text, &td.Done); err != nil {
		if err == sql.ErrNoRows {
			return &echo.HTTPError{Code: http.StatusNotFound, Message: "Todo not found"}
		}
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "Failed to query todo"}
	}
	return ctx.JSON(http.StatusOK, td)
}

// (POST /todos)
func (Server) PostTodos(ctx echo.Context) error {
	var td Todo
	if err := ctx.Bind(&td); err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "Invalid request"}

	}
	result, err := db.Exec("INSERT INTO todo (text, done) VALUES (?, ?)", td.Text, td.Done)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "Failed to insert todo"}
	}
	id, err := result.LastInsertId()
	if err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "Failed to get last insert id"}
	}
	i := int(id)
	td.Id = &i
	return ctx.JSON(http.StatusCreated, td)
}

// (PUT /todos/:id)
func (Server) PutTodosId(ctx echo.Context, id int) error {
	var td Todo
	if err := ctx.Bind(&td); err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "Invalid request"}

	}
	// TODO check if todo exists first, if not return 404
	_, err := db.Exec("UPDATE todo SET text = ?, done = ? WHERE id = ?", td.Text, td.Done, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return &echo.HTTPError{Code: http.StatusNotFound, Message: "Todo not found"}
		}
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "Failed to update todo"}

	}
	return ctx.JSON(http.StatusOK, td)
}

// (DELETE /todos/:id)
func (Server) DeleteTodosId(ctx echo.Context, id int) error {
	// TODO check if todo exists first, if not return 404
	_, err := db.Exec("DELETE FROM todo WHERE id = ?", id)
	if err != nil {
		if err == sql.ErrNoRows {
			return &echo.HTTPError{Code: http.StatusNotFound, Message: "Todo not found"}
		}
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "Failed to delete todo"}
	}
	return ctx.NoContent(http.StatusNoContent)
}
