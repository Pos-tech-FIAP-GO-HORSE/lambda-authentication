terraform:
	terraform init
	terraform plan
	terraform apply

deploy:
	cd cmd && GOOS=linux GOARCH=amd64 go build -o bootstrap main.go
	cd cmd && zip function.zip bootstrap
	cd cmd && aws lambda update-function-code --function-name auth_lambda --zip-file fileb://function.zip

