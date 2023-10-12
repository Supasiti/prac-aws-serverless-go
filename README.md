# prac-aws-serverless-go


This a simple CRUD application as my practice to deploy serverless application to AWS using CDK in Go.

## Set up 

Assumptions
- You have Go installed on your machine (version 1.20 is used here) 
- You have AWS SSO set up. To see how to set up see [here](https://docs.aws.amazon.com/singlesignon/latest/userguide/getting-started.html)
- You have set up AWS cli tools. See [here](https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html)

To install all dependencies:
```bash
make tidy
```

In this project, it is organised to be run against units. Any integration tests will be done on AWS environment. See deployment for further instruction on how to deploy to AWS.

## Deployment

This project uses AWS_PROFILE to handle deployment of Application. Assume that SSO has already been set up, run

```bash
export AWS_PROFILE=<Your Profile Name>
aws sso login
```

To see if it works
```bash
aws sts get-caller-identity
```
You should get something like:
```
{
    "UserId": "String....",
    "Account": <12 digit number>,
    "Arn": "arn:aws:sts::<account id>:assumed-role/..."
}
```

There are two stacks, Resource and Function stacks. Most of the times, you would not need to redeploy the resource stack. All stacks have similar patterns:
```
[Project]-[Service]-[Env]-api
[Project]-[Service]-[Env]-resource
```

To deploy all stacks
```
cdk deploy 
```

To deploy specific stack
```
cdk deploy pas-user-thara-api
```

To delete all the stacks
```
cdk destroy --all
```
or specific stack
```
cdk destroy pas-user-thara-api
```

## Useful commands

To seed users dynamodb in AWS, run
```bash
aws dynamodb batch-write-item --request-items file://dynamodb/user_data.json
```
This will seed users to table pas-user-thara-user. Modify the table name in `./dynamodb/user_data.json` to point to your table.

