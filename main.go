package main

import (
	"echo-user/config"
	"echo-user/routes"
	m "echo-user/middlewares"
)

func main() {
	config.InitDB()
	e := routes.New()
	m.LogMiddlewares(e)
	e.Logger.Fatal(e.Start(":8000"))
}