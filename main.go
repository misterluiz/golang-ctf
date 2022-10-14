package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/misterluiz/golang-ctf/api"
	db "github.com/misterluiz/golang-ctf/db/sqlc"
)

const (
	dbDriver      = "Postgres"
	dbSource      = "postgresql://postgres:password@localhost:5432/go-ctf?sslmode=disable"
	serverAddress = "0.0.0.0:8000"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("NÃO FOI POSsivel conectar", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)

	if err != nil {
		log.Fatal("Deu erro :)", err)
	}

}
