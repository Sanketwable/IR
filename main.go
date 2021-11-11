package main

import (
	"fmt"
	"ir/configs"
	"ir/controllers"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	configs.ReadConfig()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// port := "8080"
	port := os.Getenv("PORT")

	e.GET("/query", controllers.GetWord)
	e.GET("/addstr", controllers.Addstr)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}
