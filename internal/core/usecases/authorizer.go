package usecases

import (
	"fmt"
	"github.com/Pos-tech-FIAP-GO-HORSE/lambda-authorization/internal/service"
	"regexp"
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
	regex := regexp.MustCompile(`^\d{3}\.\d{3}\.\d{3}-\d{2}$`)
	if !regex.MatchString(cpf) {
		return "", fmt.Errorf("invalid CPF format")
	}
	userAlreadyExists, err := uc.AuthenticationService.CheckUserExists(cpf)
	if err != nil {
		return "Error check if user already exists", err
	}
	if userAlreadyExists {
		return "User authenticated successfully", nil
	}
	err = uc.AuthenticationService.CreateUser(cpf)
	if err != nil {
		return "Error create user", err
	}
	return "User created and authenticated", nil
}
