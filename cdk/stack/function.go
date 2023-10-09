package stack

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	awsgw "github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	awslambda "github.com/aws/aws-cdk-go/awscdklambdagoalpha/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type FunctionStackProps struct {
	awscdk.StackProps
	UserTableName *string
}

func NewFunctionStack(scope constructs.Construct, id string, props *FunctionStackProps) *awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}

	stack := awscdk.NewStack(scope, &id, &sprops)

	basePath := newBaseResource(stack, &id)
	newGetUserApi(stack, &GetUserApiProps{
		BasePath: basePath,
		ApiId:    &id,
	})

	return &stack
}

func newBaseResource(scope constructs.Construct, id *string) *awsgw.Resource {
	restApi := awsgw.NewRestApi(scope, id, &awsgw.RestApiProps{
		RestApiName: jsii.String("User Service"),
	})

	resource := restApi.Root().AddResource(jsii.String("users"), nil)
	return &resource
}

type GetUserApiProps struct {
	BasePath *awsgw.Resource
	ApiId    *string
}

func newGetUserApi(scope constructs.Construct, props *GetUserApiProps) {
	fnName := "getUser"
	fnId := fmt.Sprintf("%s-%s", *props.ApiId, fnName)

	lambda := awslambda.NewGoFunction(scope, jsii.String(fnId), &awslambda.GoFunctionProps{
		Entry:        jsii.String("cmd/handler/get_user"),
		FunctionName: jsii.String(fnId),

		// passing build flag to reduce bundle size
		Bundling: &awslambda.BundlingOptions{
			GoBuildFlags: jsii.Strings(`-ldflags "-s -w"`),
		},
		Environment: &map[string]*string{
			"HELLO": jsii.String("WORLD"),
		},
	})

	// add lambda integration
	apiInt := awsgw.NewLambdaIntegration(lambda, &awsgw.LambdaIntegrationOptions{
		RequestTemplates: &map[string]*string{
			"application/json": jsii.String("{ \"statusCode\": \"200\" }"),
		},
	})

	// add method
	(*props.BasePath).
		AddResource(jsii.String("{userID}"), nil).
		AddMethod(
			jsii.String("GET"),
			apiInt,
			&awsgw.MethodOptions{
				RequestParameters: &map[string]*bool{
					"method.request.path.userID": jsii.Bool(true),
				},
			})

}
