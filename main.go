package main

import (
	"database/sql"
	"log"

	"github.com/NovikovAndrew/SimpleBank/api"
	db "github.com/NovikovAndrew/SimpleBank/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:root@localhost:5432/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("cannot to connect to database, error: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)
	if err := server.Start(serverAddress); err != nil {
		log.Fatal("cannot start server, error: ", err)
	}
}
