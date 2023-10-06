package main

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/jsii-runtime-go"
)

var (
	project   = "pas"
	service   = "user"
	rootId    = fmt.Sprintf("%s-%s-thara", project, service)
	tableName = fmt.Sprintf("%s-user", rootId)
	fnStackId = fmt.Sprintf("%s-api", rootId)
)

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)
	fnDescription := fmt.Sprintf("%s function stack", rootId)

	NewFunctionStack(app, fnStackId, &FunctionStackProps{
		StackProps: awscdk.StackProps{
			Description: &fnDescription,
		},
		UserTableName: &tableName,
	})

	app.Synth(nil)
}
