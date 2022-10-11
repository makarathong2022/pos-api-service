package main

import (
	"database/sql"
	"log"
	"github.com/boincompany/pos_api_service/api"
	db "github.com/boincompany/pos_api_service/db/sqlc"
	"github.com/boincompany/pos_api_service/utils"
	_ "github.com/lib/pq"
)

/*
- Pass "." here, which means the current folder, cuz our config file is in the same location whit this main.go file
*/
func main() {
	config, err := utils.LoadConfig(".")

	if err != nil {
		log.Fatal("Cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)

	if err != nil {
		log.Fatal("cannot create server:", err)

	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot connect to server")
	}

}
