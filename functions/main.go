package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// DB Models
type Person struct {
	Id   int64  `json:"id,omitempty"`
	Name string `json:"name"`
}

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

func connectDB() (err error) {
	db, err = sql.Open("mysql", "sysdba:masterkey@tcp(database-1.ci6oumsei337.us-east-1.rds.amazonaws.com)/testDB")
	if err != nil {
		return err
	}
	err = db.Ping()
	return err
}

func main() {
	lambda.Start(Handler)
}
