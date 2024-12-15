# Easy Lamb Go

Parse, Build and Generate Terraform code for deploying AWS Lambda functions using Easy Lamb Terraform module.

## Features

- Parse a directory containing AWS Lambda functions
- Build Terraform code for deploying AWS Lambda functions
- Generate Terraform code for deploying AWS Lambda functions

## Installation

```bash
go install github.com/easy-lamb/easy-lamb-go@latest
```

## Usage

Create a `easy-lamb.json` file with your configuration at the root of your project

```json
{
  "lambdaDir": "functions",
  "terraformDir": "devops/terraform",
  "terraformFilename": "functions.tfvars",
  "buildOutput": "bin",
  "defaultParams": {
    "memory": "128",
    "timeout": "30",
    "handler": "index.handler",
    "runtime": "provided.al2023",
    "authorizer": "b2c-authorizer"
  },
  "dotenvLocation": ".env"
}
```

Run the following command

```bash
# Create a functions.tfvars file
easy-lamb-go parse
```

```bash
# Build your go functions and create functions.tfvars file
easy-lamb-go build
```
