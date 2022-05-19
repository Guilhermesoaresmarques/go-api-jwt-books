package main

import (
	"github.com/guilhermesoaresmarques/go-api-jwt-books/database"
	"github.com/guilhermesoaresmarques/go-api-jwt-books/server"
)

func main() {
	database.StartDB()
	s := server.NewServer()

	s.Run()
}
