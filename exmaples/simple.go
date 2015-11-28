package main

import (
	"github.com/goforgery/favicon"
	"github.com/goforgery/forgery2"
)

func main() {
	app := f.CreateApp()
	app.Use(favicon.Create())
	app.Listen(3000)
}
