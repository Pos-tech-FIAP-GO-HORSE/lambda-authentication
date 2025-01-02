package main

import (
	"github.com/Pos-tech-FIAP-GO-HORSE/lambda-authorization/internal/handlers"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
	"os"

	"github.com/Pos-tech-FIAP-GO-HORSE/lambda-authorization/internal/core/usecases"
	"github.com/Pos-tech-FIAP-GO-HORSE/lambda-authorization/internal/service"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

func main() {
	region := os.Getenv("AWS_REGION")
	if region == "" {
		region = "us-east-1"
	}

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})
	if err != nil {
		log.Fatalf("Error creating AWS session: %v", err)
	}

	cognitoClient := cognitoidentityprovider.New(sess)

	userPoolID := os.Getenv("COGNITO_USER_POOL_ID")
	if userPoolID == "" {
		log.Fatal("COGNITO_USER_POOL_ID environment variable is not set")
	}

	//service
	authService := service.NewAuthenticationService(cognitoClient, userPoolID)
	//use case
	usecase := usecases.NewAuthorizerUseCase(authService)
	//handler
	handler := handlers.NewAuthenticationHandler(usecase)

	lambda.Start(handler.Handler)
}
