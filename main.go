package main

import (
	"github.com/allanassis/todo-list/models"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func main() {
	e := echo.New()

	e.GET("/task/:id", func(c echo.Context) error {
		id := c.Param("id")
		t := models.Task{
			Id: id,
		}
		t, err := t.Get()
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusOK, t)
	})

	e.POST("/task", func(c echo.Context) error {
		t := &models.Task{
			Id:          uuid.New().String(),
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
