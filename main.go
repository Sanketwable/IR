package main

import (
	"IR/config"
	"IR/controllers"
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	config.ReadConfig()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	port := os.Getenv("PORT")

	e.POST("/substr", controllers.Substr)
	e.POST("/prefixstr", controllers.PrefixStr)
	e.POST("/suffixstr", controllers.Suffixstr)
	e.POST("/findstr", controllers.Findstr)
	e.POST("/addstr", controllers.Addstr)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}
