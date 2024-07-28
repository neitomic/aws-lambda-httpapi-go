# AWS Lambda API CDK Go project template

This project is a getting started entry point to implement AWS API Gateway HTTP API with AWS Lambda integration.

## Use this template

```bash
# Install gonew
go install golang.org/x/tools/cmd/gonew@latest

gonew github.com/neitomic/aws-lambda-httpapi-go your.domain/api-project
```

Some additional steps:
- rename the stack name (rename hello-world.go and change the stack name inside that file)
- look at the entrypoint: `lambda_app/cmd/api/main.go` 

## Useful commands

* `cdk bootstrap`   bootstrap the aws environment
* `cdk deploy`      deploy this stack to your default AWS account/region
* `cdk diff`        compare deployed stack with current state
* `cdk synth`       emits the synthesized CloudFormation template
* `go test`         run unit tests