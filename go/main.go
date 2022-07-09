package main

import (
	"api/controllers"
	"api/database"
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// インスタンスの作成
	e := echo.New()

	// ミドルウェアの登録
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ルーティング
	createRoutes(e)

	// DBと接続
	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_PASSWORD")
	dbName := os.Getenv("MYSQL_DATABASE")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	pool, err := database.Establish(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, host, port, dbName))
	if err != nil {
		e.Logger.Fatal(err)
	}

	pool.SetMaxOpenConns(10)
	pool.SetMaxIdleConns(10)

	// 起動
	e.Logger.Fatal(e.Start(":8080"))
}

func createRoutes(e *echo.Echo) {
	api := e.Group("/api")

	api.GET("/timetables", controllers.GetAllTimetables)

	api.GET("/timetables/:id", controllers.GetOneTimetable)
	api.GET("/timetables/:id/data", controllers.GetTimetableData)
}
