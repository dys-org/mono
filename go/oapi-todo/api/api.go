package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Server struct{}

func NewServer() Server {
	return Server{}
}

// (GET /todos)
func (Server) GetTodos(ctx echo.Context) error {
	resp := []Todo{{Id: 1, Text: "Do the laundry", Done: false}, {Id: 2, Text: "Do the laundry again", Done: false}}
	return ctx.JSON(http.StatusOK, resp)
}

// (GET /todos/:id)
func (Server) GetTodosId(ctx echo.Context, id int) error {
	resp := Todo{Id: 1, Text: "Do the laundry", Done: false}
	// id := ctx.Param("id")
	return ctx.JSON(http.StatusOK, resp)
}

// (POST /todos)
func (Server) PostTodos(ctx echo.Context) error {
	t := new(Todo)
	if err := ctx.Bind(t); err != nil {
		return err
	}
	return ctx.JSON(http.StatusCreated, t)
}

// (PUT /todos/:id)
func (Server) PutTodosId(ctx echo.Context, id int) error {
	return ctx.JSON(http.StatusOK, id)
}

// (DELETE /todos/:id)
func (Server) DeleteTodosId(ctx echo.Context, id int) error {
	return ctx.NoContent(http.StatusNoContent)
}
