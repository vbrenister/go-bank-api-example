package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/vbrenister/go-bank-api-example/api"
	db "github.com/vbrenister/go-bank-api-example/db/sqlc"
	"github.com/vbrenister/go-bank-api-example/util"
)



func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal(err)
	}
}