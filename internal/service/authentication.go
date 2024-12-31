package service

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"log"
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
	log.Println("Chegou na function de check")
	input := &cognitoidentityprovider.ListUsersInput{
		Filter:     aws.String(fmt.Sprintf("cpf = \"%s\"", cpf)),
		UserPoolId: aws.String(s.UserPoolID),
	}

	log.Println("Chamando o listUsers")
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
			{Name: aws.String("cpf"), Value: aws.String(cpf)},
		},
	}
	_, err := s.CognitoClient.AdminCreateUser(input)
	return err
}
