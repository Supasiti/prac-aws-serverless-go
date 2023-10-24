package stack

import (
	"fmt"
	"net/http"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	awsgw "github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
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

	fmt.Println()
	fmt.Printf("Creating ... Function Stack: %s\n", id)

	stack := awscdk.NewStack(scope, &id, &sprops)

	// create api gateway
	basePath := newBaseResource(stack, &id)

	// create dynamodb policy
	userTablePolicy := NewDynamodbTableIamPolicy(&IamPolicyProps{
		Actions: []*string{
			jsii.String("dynamodb:GetItem"),
			jsii.String("dynamodb:PutItem"),
		},
		PolicyName:    fmt.Sprintf("%s-dynamodb-policy", id),
		ResourceNames: []string{*props.UserTableName},
		Stack:         stack,
	})

	// create get user lambda
	newGetUserAPI(stack, &GetUserAPIProps{
		BasePath:        basePath,
		APIID:           &id,
		UserTableName:   props.UserTableName,
		UserTablePolicy: userTablePolicy,
	})

	// create create user lambda
	newCreateUserAPI(stack, &CreateUserAPIProps{
		BasePath:        basePath,
		APIID:           &id,
		UserTableName:   props.UserTableName,
		UserTablePolicy: userTablePolicy,
	})

	return &stack
}

func newBaseResource(scope constructs.Construct, id *string) *awsgw.Resource {
	restAPI := awsgw.NewRestApi(scope, id, &awsgw.RestApiProps{
		RestApiName: jsii.String("User Service"),
	})

	resource := restAPI.Root().AddResource(jsii.String("users"), nil)
	return &resource
}

type GetUserAPIProps struct {
	BasePath        *awsgw.Resource
	APIID           *string
	UserTableName   *string
	UserTablePolicy *awsiam.Policy
}

func newGetUserAPI(scope constructs.Construct, props *GetUserAPIProps) {
	fnName := "getUser"
	fnID := fmt.Sprintf("%s-%s", *props.APIID, fnName)

	fmt.Printf("Creating ... Lambda function: %s\n", fnID)

	lambda := awslambda.NewGoFunction(scope, jsii.String(fnID), &awslambda.GoFunctionProps{
		Entry:        jsii.String("cmd/handler/get_user"),
		FunctionName: jsii.String(fnID),

		// passing build flag to reduce bundle size
		Bundling: &awslambda.BundlingOptions{
			GoBuildFlags: jsii.Strings(`-ldflags "-s -w"`),
		},
		Environment: &map[string]*string{
			"USER_TABLE_NAME": props.UserTableName,
		},
	})

	// attach dynamodb policy
	lambda.Role().AttachInlinePolicy(*props.UserTablePolicy)

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
			jsii.String(http.MethodGet),
			apiInt,
			&awsgw.MethodOptions{
				RequestParameters: &map[string]*bool{
					"method.request.path.userID": jsii.Bool(true),
				},
			})
}

type CreateUserAPIProps struct {
	BasePath        *awsgw.Resource
	APIID           *string
	UserTableName   *string
	UserTablePolicy *awsiam.Policy
}

func newCreateUserAPI(scope constructs.Construct, props *CreateUserAPIProps) {
	fnName := "createUser"
	fnID := fmt.Sprintf("%s-%s", *props.APIID, fnName)

	fmt.Printf("Creating ... Lambda function: %s\n", fnID)

	lambda := awslambda.NewGoFunction(scope, jsii.String(fnID), &awslambda.GoFunctionProps{
		Entry:        jsii.String("cmd/handler/create_user"),
		FunctionName: jsii.String(fnID),

		// passing build flag to reduce bundle size
		Bundling: &awslambda.BundlingOptions{
			GoBuildFlags: jsii.Strings(`-ldflags "-s -w"`),
		},
		Environment: &map[string]*string{
			"USER_TABLE_NAME": props.UserTableName,
		},
	})

	// attach dynamodb policy
	lambda.Role().AttachInlinePolicy(*props.UserTablePolicy)

	// add lambda integration
	apiInt := awsgw.NewLambdaIntegration(lambda, &awsgw.LambdaIntegrationOptions{
		RequestTemplates: &map[string]*string{
			"application/json": jsii.String("{ \"statusCode\": \"200\" }"),
		},
	})

	// add method
	(*props.BasePath).AddMethod(jsii.String(http.MethodPost), apiInt, &awsgw.MethodOptions{})
}
