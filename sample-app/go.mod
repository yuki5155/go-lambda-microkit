module github.com/yuki5155/go-lambda-microkit

go 1.21.3

require (
	github.com/golang/mock v1.6.0
	github.com/yuki5155/go-lambda-microkit/myaws v0.0.0-unpublished
)

require (
	github.com/aws/aws-sdk-go-v2 v1.31.0 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.3.18 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.6.18 // indirect
	github.com/aws/aws-sdk-go-v2/service/cloudformation v1.54.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider v1.45.3 // indirect
	github.com/aws/smithy-go v1.21.0 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
)

replace github.com/yuki5155/go-lambda-microkit/myaws => ../myaws
