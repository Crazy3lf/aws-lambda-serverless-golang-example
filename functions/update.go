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
			"X-MyCompany-Func-Reply": "update-handler",
		},
	}

	err = json.Unmarshal([]byte(request.Body), &person)
	if err != nil {
		return Response{StatusCode: http.StatusBadRequest}, err
	}

	if person.Id > 0 && person.Name != "" {
		if err := connectDB(); err != nil {
			return Response{StatusCode: http.StatusInternalServerError}, err
		}

		_, err := db.Exec(fmt.Sprintf(`UPDATE Persons SET name = '%s' WHERE PersonID = %d`, person.Name, person.Id))
		if err != nil {
			return Response{StatusCode: http.StatusInternalServerError}, err
		}

		message = fmt.Sprintf("Edited %d.", person.Id)
	} else {
		resp.StatusCode = http.StatusBadRequest
		message = "parameters are invalid"
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
