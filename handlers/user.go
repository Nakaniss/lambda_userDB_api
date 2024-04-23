package handlers

import (
	"encoding/json"
	"fmt"
	"lmd-func/models"
	"lmd-func/repository"

	"github.com/aws/aws-lambda-go/events"
)

type UserHandler struct {
	repo *repository.UserRepository
}

func NewUserHandler(repo *repository.UserRepository) *UserHandler {
	return &UserHandler{repo: repo}
}

func (h *UserHandler) GetHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	id := request.QueryStringParameters["id"]
	if id != "" {
		user, err := h.repo.GetUser(id)
		if err != nil {
			return NotFound(fmt.Sprintf("User with ID %s not found", id))
		}
		responseBody, _ := json.Marshal(user)
		return Ok(string(responseBody))
	}

	users, err := h.repo.GetUsers()
	if err != nil {
		return BadRequest(err.Error())
	}

	responseBody, _ := json.Marshal(users)
	return Ok(string(responseBody))
}

func (h *UserHandler) PostHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	id := request.QueryStringParameters["id"]
	if id != "" {
		return MethodNotAllowed()
	}

	var user models.User
	if err := json.Unmarshal([]byte(request.Body), &user); err != nil {
		return BadRequest(err.Error())
	}

	if err := h.repo.CreateUser(&user); err != nil {
		return BadRequest(err.Error())
	}

	responseBody, _ := json.Marshal(user)
	return Ok(string(responseBody))
}

func (h *UserHandler) PutHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	id := request.QueryStringParameters["id"]
	if id == "" {
		return BadRequest("ID is required")
	}

	_, err := h.repo.GetUser(id)
	if err != nil {
		return NotFound(fmt.Sprintf("User with ID %s not found", id))
	}

	var updatedUser models.User
	if err := json.Unmarshal([]byte(request.Body), &updatedUser); err != nil {
		return BadRequest(err.Error())
	}

	updatedUser.ID = id
	if err := h.repo.UpdateUser(&updatedUser); err != nil {
		return BadRequest(err.Error())
	}

	responseBody, _ := json.Marshal(updatedUser)
	return Ok(string(responseBody))
}

func (h *UserHandler) DeleteHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	id := request.QueryStringParameters["id"]
	if id == "" {
		return BadRequest("ID is required")
	}

	if err := h.repo.DeleteUser(id); err != nil {
		return NotFound(fmt.Sprintf("User with ID %s not found", id))
	}
	return NoContent()
}
