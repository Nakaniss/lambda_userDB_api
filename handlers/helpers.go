package handlers

import (
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
)

// addCorsHeaders は、レスポンスにCORSヘッダーを追加します
func addCorsHeaders(resp *events.APIGatewayProxyResponse) {
	resp.Headers = map[string]string{
		"Access-Control-Allow-Origin":  os.Getenv("ORIGIN_URL"),
		"Access-Control-Allow-Headers": "Content-Type, X-Amz-Date, Authorization, X-Api-Key, X-Amz-Security-Token",
		"Access-Control-Allow-Methods": "GET, POST, PUT, DELETE",
	}
}

// Ok は200 OKのレスポンスを返します
func Ok(body string) (events.APIGatewayProxyResponse, error) {
	resp := events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       body,
	}
	addCorsHeaders(&resp)
	return resp, nil
}

// NoContent は204 No Contentのレスポンスを返します
func NoContent() (events.APIGatewayProxyResponse, error) {
	resp := events.APIGatewayProxyResponse{
		StatusCode: http.StatusNoContent,
	}
	addCorsHeaders(&resp)
	return resp, nil
}

// NotFound は404 Not Foundのレスポンスを返します
func NotFound(message string) (events.APIGatewayProxyResponse, error) {
	resp := events.APIGatewayProxyResponse{
		StatusCode: http.StatusNotFound,
		Body:       message,
	}
	addCorsHeaders(&resp)
	return resp, nil
}

// BadRequest は400 Bad Requestのレスポンスを返します
func BadRequest(message string) (events.APIGatewayProxyResponse, error) {
	resp := events.APIGatewayProxyResponse{
		StatusCode: http.StatusBadRequest,
		Body:       message,
	}
	addCorsHeaders(&resp)
	return resp, nil
}

// MethodNotAllowed は405 Method Not Allowedのレスポンスを返します
func MethodNotAllowed() (events.APIGatewayProxyResponse, error) {
	resp := events.APIGatewayProxyResponse{
		StatusCode: http.StatusMethodNotAllowed,
		Body:       http.StatusText(http.StatusMethodNotAllowed),
	}
	addCorsHeaders(&resp)
	return resp, nil
}
