package main

import (
	"github.com/milanlopatin/todo-app"
	"github.com/milanlopatin/todo-app/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("Error has occured while running server: %s", err.Error())
	}
}
