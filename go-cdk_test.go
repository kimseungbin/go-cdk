package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/assertions"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/jsii-runtime-go"
	"go-cdk/hitcounter"
	"testing"
)

// import (
// 	"testing"

// 	"github.com/aws/aws-cdk-go/awscdk/v2"
// 	"github.com/aws/aws-cdk-go/awscdk/v2/assertions"
// 	"github.com/aws/jsii-runtime-go"
// )

// example tests. To run these tests, uncomment this file along with the
// example resource in go-cdk_test.go
func TestHitCounterConstruct(t *testing.T) {
	defer jsii.Close()

	// Given

	stack := awscdk.NewStack(nil, nil, nil)

	// When
	testFn := awslambda.NewFunction(stack, jsii.String("TestFunc"), &awslambda.FunctionProps{
		Code:    awslambda.Code_FromAsset(jsii.String("lambda"), nil),
		Runtime: awslambda.Runtime_NODEJS_16_X(),
		Handler: jsii.String("hello.handler"),
	})
	hitcounter.NewHitCounter(stack, "MyTestConstruct", &hitcounter.HitCounterProps{
		Downstream: testFn,
	})

	// Then
	template := assertions.Template_FromStack(stack, nil)
	template.ResourceCountIs(jsii.String("AWS::DynamoDB::Table"), jsii.Number(1))
}

// func TestGoCdkStack(t *testing.T) {
// 	// GIVEN
// 	app := awscdk.NewApp(nil)

// 	// WHEN
// 	stack := NewGoCdkStack(app, "MyStack", nil)

// 	// THEN
// 	template := assertions.Template_FromStack(stack)

// 	template.HasResourceProperties(jsii.String("AWS::SQS::Queue"), map[string]interface{}{
// 		"VisibilityTimeout": 300,
// 	})
// }
