package main

import (
	"context"
	"log"
	"simplebank/api"
	db "simplebank/db/sqlc"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	dbSource      = "postgresql://root:secret@127.0.0.1:5432/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := pgxpool.New(context.Background(), dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}

}
