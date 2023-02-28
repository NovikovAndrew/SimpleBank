package main

import (
	"database/sql"
	"log"

	"github.com/NovikovAndrew/SimpleBank/api"
	db "github.com/NovikovAndrew/SimpleBank/db/sqlc"
	"github.com/NovikovAndrew/SimpleBank/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")

	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("cannot to connect to database, error: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)
	if err := server.Start(config.ServerAddress); err != nil {
		log.Fatal("cannot start server, error: ", err)
	}
}
