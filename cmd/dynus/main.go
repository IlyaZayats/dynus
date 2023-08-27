package main

import (
	"github.com/IlyaZayats/dynus/internal/api/handlers/v1"
	"github.com/IlyaZayats/dynus/internal/api/server"
	"github.com/IlyaZayats/dynus/internal/db/postgres"
)

func main() {
	database := postgres.OpenConnection()
	route := v1.NewHandlers(&database)
	r := server.InitNewServer(route)
	r.Run()
}
