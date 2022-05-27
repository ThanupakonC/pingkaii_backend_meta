package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	api "github.com/thanupakonc/pingkaii_meta/api"
	db "github.com/thanupakonc/pingkaii_meta/db/sqlc"
	"log"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://doadmin:AVNS_F1H-ygWnfaSfvba@apartment-db-postgresql-sgp-1-do-user-10407607-0.b.db.ondigitalocean.com:25060/thailand_area_db?sslmode=require"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	fmt.Println(server)
	//
	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
