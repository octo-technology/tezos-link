package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	pkgdatabase "github.com/octo-technology/tezos-link/backend/pkg/infrastructure/database"
	"log"
	"os"
)

func main() {
	lambda.Start(HandleRequest)
}

// HandleRequest execute the logic
func HandleRequest(ctx context.Context) (string, error) {
	log.Print("metrics clean starting")

	dnsStr := fmt.Sprintf("postgres://%s:%s@%s/%s", os.Getenv("DATABASE_USERNAME"), os.Getenv("DATABASE_PASSWORD"), os.Getenv("DATABASE_URL"), os.Getenv("DATABASE_TABLE"))

	var Connection *sql.DB
	con, err := sql.Open("postgres", dnsStr)
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

	metricsRepo := pkgdatabase.NewPostgresMetricsRepository(Connection)

	err2 := metricsRepo.Remove3MonthOldMetrics()

	log.Print(err2)

	return "metric clean started.", nil
}

