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
        Action   = "logs:*"
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
