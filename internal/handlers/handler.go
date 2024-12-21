package handlers

import (
	"github.com/Pos-tech-FIAP-GO-HORSE/lambda-authorization/internal/core/usecases"
	"github.com/aws/aws-lambda-go/events"
	"net/http"
)

type AuthenticationHandler struct {
	ValidationUserUseCase *usecases.AuthorizerUseCase
}

func NewAuthenticationHandler(validationUserUseCase *usecases.AuthorizerUseCase) *AuthenticationHandler {
	return &AuthenticationHandler{
		ValidationUserUseCase: validationUserUseCase,
	}
}

func (h *AuthenticationHandler) Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	cpf := request.Headers["cpf"]
	if cpf == "" {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       "CPF header is missing",
		}, nil
	}
	authenticationResult, err := h.ValidationUserUseCase.AuthenticateUser(cpf)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}, err
	}
	return events.APIGatewayProxyResponse{
		StatusCode:      http.StatusOK,
		Body:            authenticationResult,
		IsBase64Encoded: false,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}, nil
}
