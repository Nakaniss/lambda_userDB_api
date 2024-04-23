package routes

import (
	"lmd-func/handlers"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

type Router struct {
	userHandler *handlers.UserHandler
}

func NewRouter(userHandler *handlers.UserHandler) *Router {
	return &Router{userHandler: userHandler}
}

func (r *Router) Route(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	switch request.HTTPMethod {
	case http.MethodPost:
		return r.userHandler.PostHandler(request)
	case http.MethodGet:
		return r.userHandler.GetHandler(request)
	case http.MethodDelete:
		return r.userHandler.DeleteHandler(request)
	case http.MethodPut:
		return r.userHandler.PutHandler(request)
	default:
		return handlers.MethodNotAllowed()
	}
}
