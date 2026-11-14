package main

import (
	"fmt"

	"github.com/go-chi/chi/v5"
	"github.com/JoaoVitor615/hook-flow/internal/api"
)

func main() {
	server := api.NewServer()
	server.InitHandlers()
	
}