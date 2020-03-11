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
	con, err := sql.Open("postgres", config.ProxyConfig.Database.Url)
	if err != nil {
		log.Fatal("Could not open DB: ", err)
	}
	// max_connection RDS is set to ~700
	con.SetMaxOpenConns(600)

	err = con.Ping()
	if err != nil {
		log.Fatal("Could not open DB: ", err)
	}

	log.Println("DB initialized")
	Connection = con
}
