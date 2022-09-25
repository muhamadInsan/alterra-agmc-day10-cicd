package main

import (
	"day4/tugas-crud-dinamis/config"
	"day4/tugas-crud-dinamis/middlewares"
	"day4/tugas-crud-dinamis/routes"
	// "github.com/iswanulumam/go-training-restful/routes"
)

func main() {

	config.InitDB()
	e := routes.New()

	middlewares.LogMiddlewares(e)
	// middlewares.LogMiddleware(e)

	e.Logger.Fatal(e.Start(":8080"))
}
