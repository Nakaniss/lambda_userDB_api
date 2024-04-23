package handlers

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

// Ok は200 OKのレスポンスを返します
func Ok(body string) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       body,
	}, nil
}

// NoContent は204 No Contentのレスポンスを返します
func NoContent() (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusNoContent,
	}, nil
}

// NotFound は404 Not Foundのレスポンスを返します
func NotFound(message string) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusNotFound,
		Body:       message,
	}, nil
}

// BadRequest は400 Bad Requestのレスポンスを返します
func BadRequest(message string) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusBadRequest,
		Body:       message,
	}, nil
}

// MethodNotAllowed は405 Method Not Allowedのレスポンスを返します
func MethodNotAllowed() (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusMethodNotAllowed,
		Body:       http.StatusText(http.StatusMethodNotAllowed),
	}, nil
}
