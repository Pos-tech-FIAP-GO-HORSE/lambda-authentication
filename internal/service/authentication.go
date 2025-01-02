package service

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

type AuthenticationService struct {
	CognitoClient *cognitoidentityprovider.CognitoIdentityProvider
	UserPoolID    string
}

func NewAuthenticationService(cognitoClient *cognitoidentityprovider.CognitoIdentityProvider, userPoolID string) *AuthenticationService {
	return &AuthenticationService{
		CognitoClient: cognitoClient,
		UserPoolID:    userPoolID,
	}
}

func (s *AuthenticationService) CheckUserExists(cpf string) (bool, error) {
	input := &cognitoidentityprovider.ListUsersInput{
		Filter:     aws.String(fmt.Sprintf("username = \"%s\"", cpf)),
		UserPoolId: aws.String(s.UserPoolID),
	}

	result, err := s.CognitoClient.ListUsers(input)
	if err != nil {
		return false, err
	}

	return len(result.Users) > 0, nil
}

func (s *AuthenticationService) CreateUser(cpf string) error {
	input := &cognitoidentityprovider.AdminCreateUserInput{
		UserPoolId: aws.String(s.UserPoolID),
		Username:   aws.String(cpf),
		UserAttributes: []*cognitoidentityprovider.AttributeType{
			{Name: aws.String("username"), Value: aws.String(cpf)},
		},
	}
	_, err := s.CognitoClient.AdminCreateUser(input)
	if err != nil {
		return fmt.Errorf("failed to create user in Cognito: %w", err)
	}
	return nil
}
