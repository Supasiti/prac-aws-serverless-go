package stack

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/jsii-runtime-go"
)

type IamPolicyProps struct {
	Actions       []*string
	Effect        awsiam.Effect // this is an enum of type string
	PolicyName    string
	ResourceNames []string
	Stack         awscdk.Stack
}

func NewDynamodbTableIamPolicy(props *IamPolicyProps) *awsiam.Policy {
	return NewServiceIamPolicy(&ServiceIamPolicyProps{
		IamPolicyProps: *props,
		ResourceType:   "table/",
		Service:        "dynamodb",
	})
}

type ServiceIamPolicyProps struct {
	IamPolicyProps
	ResourceType string
	Service      string
}

func NewServiceIamPolicy(props *ServiceIamPolicyProps) *awsiam.Policy {
	fmt.Printf("Creating ... IAM policy for %s\n", props.Service)

	resourceArns := newResources(props)
	stmtProps := &awsiam.PolicyStatementProps{
		Actions:   &props.Actions,
		Effect:    awsiam.Effect_ALLOW,
		Resources: &resourceArns,
	}

	if len(props.Effect) > 0 {
		stmtProps.Effect = props.Effect
	}

	stmt := awsiam.NewPolicyStatement(stmtProps)
	policy := awsiam.NewPolicy(props.Stack, jsii.String(props.PolicyName), &awsiam.PolicyProps{
		Statements: &[]awsiam.PolicyStatement{stmt},
	})

	return &policy

}

func newResources(props *ServiceIamPolicyProps) []*string {
	accountId := props.Stack.Account()
	region := props.Stack.Region()

	result := []*string{}
	for _, resourceName := range props.ResourceNames {
		// ARN looks like
		// arn:aws:<service>:<region>:<accountId>:<resourceType><name>
		arn := fmt.Sprintf("arn:aws:%s:%s:%s:%s%s", props.Service, *region, *accountId, props.ResourceType, resourceName)

		result = append(result, jsii.String(arn))
	}
	return result

}
