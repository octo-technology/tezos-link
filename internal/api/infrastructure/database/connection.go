package database

import (
	"database/sql"
	"github.com/octo-technology/tezos-link/backend/config"
	"log"
)

// Connection represents the database connection
var Connection *sql.DB

// Configure setup the database connection
func Configure() {
	con, err := sql.Open("postgres", config.BackendConfig.Db.Url)
	if err != nil {
		log.Fatal("Could not open DB: ", err)
	}

	err = con.Ping()
	if err != nil {
		log.Fatal("Could not open DB: ", err)
	}

	log.Println("DB initialized")
	Connection = con
}
