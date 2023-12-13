package main

import (
	"todo-htmx/model"
	"todo-htmx/routes"
)

func main() {
	model.Setup()
	routes.SetupAndRun()
}
