provider "aws" {
  region = "us-east-1"
}

resource "aws_iam_role" "lambda_exec_role" {
  name = "lambda_exec_role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action    = "sts:AssumeRole"
        Principal = {
          Service = "lambda.amazonaws.com"
        }
        Effect    = "Allow"
        Sid       = ""
      },
    ]
  })
}

resource "aws_iam_role_policy" "lambda_policy" {
  name = "lambda_policy"
  role = aws_iam_role.lambda_exec_role.id

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action   = [
          "logs:CreateLogGroup",
          "logs:CreateLogStream",
          "logs:PutLogEvents"
        ]
        Resource = "*"
        Effect   = "Allow"
      },
    ]
  })
}

resource "aws_lambda_function" "auth_lambda" {
  function_name = "auth_lambda"
  role          = aws_iam_role.lambda_exec_role.arn
  handler       = "main"
  runtime       = "provided.al2"

  filename      = "function.zip"
  source_code_hash = filebase64sha256("function.zip")
}

resource "aws_api_gateway_rest_api" "auth_api" {
  name        = "auth-api"
  description = "API Gateway para autenticação com Lambda"
}

resource "aws_api_gateway_resource" "auth_resource" {
  rest_api_id = aws_api_gateway_rest_api.auth_api.id
  parent_id   = aws_api_gateway_rest_api.auth_api.root_resource_id
  path_part   = "auth"
}

resource "aws_api_gateway_method" "auth_method" {
  rest_api_id   = aws_api_gateway_rest_api.auth_api.id
  resource_id   = aws_api_gateway_resource.auth_resource.id
  http_method   = "POST"
  authorization = "NONE"
}

resource "aws_api_gateway_integration" "auth_integration" {
  rest_api_id = aws_api_gateway_rest_api.auth_api.id
  resource_id = aws_api_gateway_resource.auth_resource.id
  http_method = aws_api_gateway_method.auth_method.http_method
  integration_http_method = "POST"
  type       = "AWS_PROXY"
  uri        = "arn:aws:apigateway:us-east-1:lambda:path/2015-03-31/functions/${aws_lambda_function.auth_lambda.arn}/invocations"
}

resource "aws_lambda_permission" "allow_api_gateway" {
  statement_id  = "AllowExecutionFromAPIGateway"
  action        = "lambda:InvokeFunction"
  principal     = "apigateway.amazonaws.com"
  function_name = aws_lambda_function.auth_lambda.function_name
}