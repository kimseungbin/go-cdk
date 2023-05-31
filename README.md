# Welcome to your CDK Go project!

This is a blank project for CDK development with Go.

The `cdk.json` file tells the CDK toolkit how to execute your app.

## Useful commands

 * `cdk deploy`      deploy this stack to your default AWS account/region
 * `cdk diff`        compare deployed stack with current state
 * `cdk synth`       emits the synthesized CloudFormation template
 * `go test`         run unit tests

# Go-CDK

This is a demo repository showing basic skills on AWS CDK v2 in Go.

It creates Lambdas in 2 languages
- Go
  - GoHelloHandler
- JS
  - HelloHandler
  - HitCounterHandler

The **GoHelloHandler** utilizes the [Amazon Lambda Golang Library
](https://pkg.go.dev/github.com/aws/aws-cdk-go/awscdklambdagoalpha/v2#section-readme) to compile Go code before deploying. This approach allows for Lambda deployment similar to JS or Python.

The **HelloHandler** is a simple Lambda example written in JS that returns the requested path.

The **HitCounterHandler** acts as an upstream handler for the HelloHandler. When requested, it increments the hit count for a given path in DynamoDB. It is part of HelloHitCounter Construct.

It also creates 1 DynamoDB table within HelloHitCounter Construct. The table has `path` as partition key.


# Notes

Mainly following the tutorial from: https://cdkworkshop.com/
[AWS SDK for Go V2](https://aws.github.io/aws-sdk-go-v2/docs/getting-started/)
[Testing Constructs](https://docs.aws.amazon.com/cdk/v2/guide/testing.html)