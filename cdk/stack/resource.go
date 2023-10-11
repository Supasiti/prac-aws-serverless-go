package stack

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type ResourceStackProps struct {
	awscdk.StackProps
	UserTableName *string
}

func NewResourceStack(scope constructs.Construct, id string, props *ResourceStackProps) *awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}

	fmt.Println()
	fmt.Printf("Creating ... Resource Stack: %s\n", id)

	stack := awscdk.NewStack(scope, &id, &sprops)

	newUserTable(stack, id, props)

	return &stack
}

func newUserTable(scope constructs.Construct, id string, props *ResourceStackProps) *awsdynamodb.TableV2 {
	tableId := fmt.Sprintf("%s-user", id)
	table := awsdynamodb.NewTableV2(scope, jsii.String(tableId), &awsdynamodb.TablePropsV2{
		TableName:     props.UserTableName,
		Billing:       awsdynamodb.Billing_OnDemand(),
		RemovalPolicy: awscdk.RemovalPolicy_DESTROY,
		PartitionKey: &awsdynamodb.Attribute{
			Name: jsii.String("$pk"),
			Type: awsdynamodb.AttributeType_STRING,
		},
		SortKey: &awsdynamodb.Attribute{
			Name: jsii.String("$sk"),
			Type: awsdynamodb.AttributeType_STRING,
		},
	})

	return &table
}
