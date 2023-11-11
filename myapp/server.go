package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

var todos = make(map[int]*Todo)
var id = 0

func main() {
	e := echo.New()

	// Routes for CRUD operations
	e.POST("/todos", createTodo)
	e.GET("/todos", getTodos)
	e.GET("/todos/:id", getTodo)
	e.PUT("/todos/:id", updateTodo)
	e.DELETE("/todos/:id", deleteTodo)

	e.Logger.Fatal(e.Start(":8080"))
}

func createTodo(c echo.Context) error {
	t := new(Todo)
	if err := c.Bind(t); err != nil {
		return err
	}
	id++
	t.ID = id
	todos[id] = t
	println("Create ToDo called")
	return c.JSON(http.StatusCreated, t)
}

func getTodos(c echo.Context) error {
	var todoList []*Todo
	for _, t := range todos {
		todoList = append(todoList, t)
	}
	return c.JSON(http.StatusOK, todoList)
}

func getTodo(c echo.Context) error {
	id := c.Param("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID")
	}
	t, ok := todos[i]
	if !ok {
		return c.JSON(http.StatusNotFound, "Todo not found")
	}
	return c.JSON(http.StatusOK, t)
}

func updateTodo(c echo.Context) error {
	id := c.Param("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID")
	}
	t, ok := todos[i]
	if !ok {
		return c.JSON(http.StatusNotFound, "Todo not found")
	}
	if ok {
		return c.JSON(http.StatusNotFound, "Todo Deleted")
	}
	updatedTodo := new(Todo)
	if err := c.Bind(updatedTodo); err != nil {
		return err
	}
	t.Title = updatedTodo.Title
	return c.JSON(http.StatusOK, t)
}

func deleteTodo(c echo.Context) error {
	id := c.Param("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID")
	}
	_, ok := todos[i]
	if !ok {
		return c.JSON(http.StatusNotFound, "Todo not found")
	}
	if ok {
		return c.JSON(http.StatusNotFound, "Todo Deleted")
	}

	delete(todos, i)
	return c.NoContent(http.StatusNoContent)
}
