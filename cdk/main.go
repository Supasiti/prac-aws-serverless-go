package main

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/jsii-runtime-go"

	"github.com/supasiti/prac-aws-serverless-go/cdk/stack"
)

var (
	project    = "pas"
	service    = "user"
	rootId     = fmt.Sprintf("%s-%s-thara", project, service)
	tableName  = fmt.Sprintf("%s-user", rootId)
	apiId      = fmt.Sprintf("%s-api", rootId)
	resourceId = fmt.Sprintf("%s-resource", rootId)
)

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	// create resource
	resourceDescription := fmt.Sprintf("%s resource stack", rootId)

	stack.NewResourceStack(app, resourceId, &stack.ResourceStackProps{
		StackProps: awscdk.StackProps{
			Description: jsii.String(resourceDescription),
		},
		UserTableName: jsii.String(tableName),
	})

	// create api stack
	apiDescription := fmt.Sprintf("%s function stack", rootId)

	stack.NewFunctionStack(app, apiId, &stack.FunctionStackProps{
		StackProps: awscdk.StackProps{
			Description: jsii.String(apiDescription),
		},
		UserTableName: jsii.String(tableName),
	})

	app.Synth(nil)
}
