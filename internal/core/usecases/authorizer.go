package usecases

import (
	"fmt"
	"github.com/Pos-tech-FIAP-GO-HORSE/lambda-authorization/internal/service"
)

type AuthorizerUseCase struct {
	AuthenticationService *service.AuthenticationService
}

func NewAuthorizerUseCase(authenticationService *service.AuthenticationService) *AuthorizerUseCase {
	return &AuthorizerUseCase{
		AuthenticationService: authenticationService,
	}
}
func (uc *AuthorizerUseCase) AuthenticateUser(cpf string) (string, error) {
	userAlreadyExists, err := uc.AuthenticationService.CheckUserExists(cpf)
	if err != nil {
		return "", fmt.Errorf("failed to check if user exists: %w", err)
	}

	if userAlreadyExists {
		return "User authenticated successfully", nil
	}

	err = uc.AuthenticationService.CreateUser(cpf)
	if err != nil {
		return "", fmt.Errorf("failed to create user: %w", err)
	}

	return "User created and authenticated successfully", nil
}
