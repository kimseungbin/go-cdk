package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/assertions"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/jsii-runtime-go"
	"go-cdk/hitcounter"
	"testing"
)

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
