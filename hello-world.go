package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2integrations"
	"github.com/aws/aws-cdk-go/awscdklambdagoalpha/v2"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type HelloWorldStackProps struct {
	awscdk.StackProps
}

func NewHelloWorldStack(scope constructs.Construct, id string, props *HelloWorldStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	lambdaApiFunc := awscdklambdagoalpha.NewGoFunction(stack, jsii.String("lambdaapi"), &awscdklambdagoalpha.GoFunctionProps{
		Entry: jsii.String("lambda_app/cmd/api"),
		Bundling: &awscdklambdagoalpha.BundlingOptions{
			ForcedDockerBundling: jsii.Bool(true),
		},
	})

	lambdaApiIntegration := awsapigatewayv2integrations.NewHttpLambdaIntegration(
		jsii.String("LambdaAPIIntegration"),
		lambdaApiFunc,
		&awsapigatewayv2integrations.HttpLambdaIntegrationProps{})

	api := awsapigatewayv2.NewHttpApi(stack, jsii.String("hello-world-api"), &awsapigatewayv2.HttpApiProps{})
	api.AddRoutes(&awsapigatewayv2.AddRoutesOptions{
		Path:        jsii.String("/hello"),
		Integration: lambdaApiIntegration,
	})

	awscdk.NewCfnOutput(stack, jsii.String("hello-world-api-url"), &awscdk.CfnOutputProps{
		Value:      api.Url(),
		ExportName: jsii.String("hello-world-api-url"),
	})
	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	NewHelloWorldStack(app, "HelloWorldStack", &HelloWorldStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env() *awscdk.Environment {
	// If unspecified, this stack will be "environment-agnostic".
	// Account/Region-dependent features and context lookups will not work, but a
	// single synthesized template can be deployed anywhere.
	//---------------------------------------------------------------------------
	return nil

	// Uncomment if you know exactly what account and region you want to deploy
	// the stack to. This is the recommendation for production stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String("123456789012"),
	//  Region:  jsii.String("us-east-1"),
	// }

	// Uncomment to specialize this stack for the AWS Account and Region that are
	// implied by the current CLI configuration. This is recommended for dev
	// stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
	//  Region:  jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
	// }
}
