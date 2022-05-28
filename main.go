package main

import (
	"bank/api"
	db "bank/db/sqlc"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

const (
	dbDriver     = "postgres"
	dbSource     = "postgresql://root:secret@localhost:5432/bank?sslmode=disable"
	serverAdress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAdress)
	if err != nil {
		log.Fatal("Cannot start server", err)
	}
}
