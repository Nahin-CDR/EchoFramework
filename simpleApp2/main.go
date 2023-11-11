package main

import (
	"html/template"
	"io"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Task struct {
	ID   int
	Name string
}

var tasks = []Task{
	{1, "Learn Golang"},
	{2, "Build a web app"},
	{3, "Deploy the app"},
}

func render(c echo.Context, code int, name string, data interface{}) error {
	return c.Render(code, name, data)
}

func getTasks(c echo.Context) error {
	return render(c, http.StatusOK, "index.html", tasks)
}

func addTask(c echo.Context) error {
	name := c.FormValue("name")
	if name != "" {
		taskID := len(tasks) + 1
		newTask := Task{taskID, name}
		tasks = append(tasks, newTask)
	}
	return c.Redirect(http.StatusFound, "/")
}

func deleteTask(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid task ID")
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			break
		}
	}

	return c.Redirect(http.StatusFound, "/")
}

func updateTaskForm(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid task ID")
	}

	taskToUpdate := Task{}
	for _, task := range tasks {
		if task.ID == id {
			taskToUpdate = task
			break
		}
	}

	return render(c, http.StatusOK, "update.html", taskToUpdate)
}

func updateTask(c echo.Context) error {
	id, err := strconv.Atoi(c.FormValue("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid task ID")
	}

	newName := c.FormValue("name")
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Name = newName
			break
		}
	}

	return c.Redirect(http.StatusFound, "/")
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/static", "static")

	e.GET("/", getTasks)
	e.POST("/add", addTask)
	e.GET("/delete/:id", deleteTask)
	e.GET("/update/:id", updateTaskForm)
	e.POST("/update", updateTask)

	e.Renderer = &Template{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}

	e.Start(":8080")
}

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
