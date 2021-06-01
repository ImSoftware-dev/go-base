package main

import (
	"go-trendy-wash-backend/controllers"
	"go-trendy-wash-backend/db"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var conn = db.ConnectDB()

func main() {

	e := echo.New()
	e.Use(middleware.CORS())
	e.Static("/api/v1/image", "image")
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"Access-Control-Allow-Origin", "*"},
		AllowHeaders: []string{"authorization", "Content-Type"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete, http.MethodOptions},
	}))

	_publicAPI := e.Group("/api/v1")

	controllers.BannerDock(_publicAPI)

	e.Logger.Fatal(e.Start(":9001"))

}
