package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"net/http"
)

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(request events.APIGatewayProxyRequest) (resp Response, err error) {
	var buf bytes.Buffer
	var message string
	var person Person
	resp = Response{
		StatusCode:      http.StatusOK,
		IsBase64Encoded: false,
		Headers: map[string]string{
			"Content-Type":           "application/json",
			"X-MyCompany-Func-Reply": "create-handler",
		},
	}

	err = json.Unmarshal([]byte(request.Body), &person)
	if err != nil {
		return Response{StatusCode: http.StatusBadRequest}, err
	}

	if person.Name == "" {
		resp.StatusCode = http.StatusBadRequest
		message = "parameter 'name' is invalid"
	}else{
		if err:=connectDB();err!=nil{
			return Response{StatusCode: http.StatusInternalServerError}, err
		}

		_,err:=db.Exec(fmt.Sprintf(`INSERT INTO Persons(Name) VALUES('%s')`,person.Name))
		if err!=nil{
			return Response{StatusCode: http.StatusInternalServerError}, err
		}

		message = fmt.Sprintf("Added %s.",person.Name)
	}

	body, err := json.Marshal(map[string]interface{}{
		"message": message,
	})
	if err != nil {
		return Response{StatusCode: http.StatusInternalServerError}, err
	}

	json.HTMLEscape(&buf, body)

	resp.Body = buf.String()

	return resp, nil
}
