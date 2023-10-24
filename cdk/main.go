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
	rootID     = fmt.Sprintf("%s-%s-thara", project, service)
	tableName  = fmt.Sprintf("%s-user", rootID)
	apiID      = fmt.Sprintf("%s-api", rootID)
	resourceID = fmt.Sprintf("%s-resource", rootID)
)

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	// create resource
	resourceDescription := fmt.Sprintf("%s resource stack", rootID)

	stack.NewResourceStack(app, resourceID, &stack.ResourceStackProps{
		StackProps: awscdk.StackProps{
			Description: jsii.String(resourceDescription),
		},
		UserTableName: jsii.String(tableName),
	})

	// create api stack
	apiDescription := fmt.Sprintf("%s function stack", rootID)

	stack.NewFunctionStack(app, apiID, &stack.FunctionStackProps{
		StackProps: awscdk.StackProps{
			Description: jsii.String(apiDescription),
		},
		UserTableName: jsii.String(tableName),
	})

	app.Synth(nil)
}
