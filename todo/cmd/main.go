package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}

// e.GET("/users/:id", getUser)
func getUser(c echo.Context) error {
	// User ID from path 'users/:id'
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

// e.GET("/tasks", getTasks)
// タスク一覧を取得する。

// e.POST("/tasks", postTasks)
// タスクを追加する。

// e.PUT("/tasks/:id", putTasks)
// タスクを完了にする。

// e.DELETE("/tasks/:id", deleteTasks)
// タスクを削除する。
