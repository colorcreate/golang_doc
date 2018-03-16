package main

import (
	"log"

	"golang_doc/go-resty/api"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/test", test.GetAll)
	e.GET("/test/:id", test.GetByID)
	e.POST("/test", test.Create)
	e.GET("/hello", test.Hello)
	e.DELETE("/test/:id", test.Remove)
	e.PUT("/test/:id", test.Update)
	log.Fatal(e.Start(":8000"))
}
