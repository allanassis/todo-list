package main

import (
	"github.com/allanassis/todo-list/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/task", func(c echo.Context) error {
		t := &models.Task{
			CreatedDate: time.Now(),
		}
		if err := c.Bind(t); err != nil {
			return err
		}
		t.Save()
		return c.JSON(http.StatusCreated, t)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
