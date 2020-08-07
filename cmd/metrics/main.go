package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	pkgdatabase "github.com/octo-technology/tezos-link/backend/pkg/infrastructure/database"
)

func main() {
	lambda.Start(HandleRequest)
}

// HandleRequest execute the logic
func HandleRequest(ctx context.Context) (string, error) {
	log.Print("retrieving database secret on secretmanager")

	rdsPassword := retrieveRDSSecretOnSecretManager()

	log.Print("metrics-cleaner: clean task is starting.")

	dnsStr := fmt.Sprintf("postgres://%s:%s@%s/%s",
		os.Getenv("DATABASE_USERNAME"),
		*rdsPassword,
		os.Getenv("DATABASE_URL"),
		os.Getenv("DATABASE_NAME"))

	log.Println(fmt.Sprintf("metrics-cleaner: Initialize connection with database %s", os.Getenv("DATABASE_URL")))
	sess, err := sql.Open("postgres", dnsStr)
	if err != nil {
		panic(fmt.Sprintf("Error: Could not open connection with the DB: %v", err))
	}
	// max_connection RDS is set to ~700
	sess.SetMaxOpenConns(600)

	err = sess.Ping()
	if err != nil {
		panic(fmt.Sprintf("Could not join the DB: %v", err))
	}

	log.Println("metrics-cleaner: DB connection initialized")

	metricsRepo := pkgdatabase.NewPostgresMetricsRepository(sess)

	err = metricsRepo.RemoveThreeMonthsOldMetrics()
	if err != nil {
		panic(fmt.Sprintf("Error: An error occure when running RemoveThreeMonthsOldMetrics() method: %v", err))
	}

	log.Println("metrics-cleaner: clean task is complete.")
	return "metrics-cleaner: clean task is complete.", nil
}

func retrieveRDSSecretOnSecretManager() *string {
	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("S3_REGION"))},
	)

	secretName := os.Getenv("DATABASE_PASSWORD_SECRET_ARN")

	svc := secretsmanager.New(sess)

	params := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	}

	res, err := svc.GetSecretValue(params)

	if err != nil {
		panic(fmt.Sprintf("Error when trying to retrieve the secret %s : %v", secretName, err))
	}

	rdsPassword := res.SecretString

	return rdsPassword
}
