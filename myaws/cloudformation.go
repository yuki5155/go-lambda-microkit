package myaws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
)

type CloudFormationAPI interface {
	DescribeStacks(ctx context.Context, params *cloudformation.DescribeStacksInput, optFns ...func(*cloudformation.Options)) (*cloudformation.DescribeStacksOutput, error)
}

type CloudFormationClientInterface interface {
	GetCloudFormationOutput(ctx context.Context, stackName, outputKey string) (string, error)
}

type cloudFormationClient struct {
	API CloudFormationAPI
}

func NewCloudFormationClient(api CloudFormationAPI) CloudFormationClientInterface {
	return &cloudFormationClient{API: api}
}

func (c *cloudFormationClient) GetCloudFormationOutput(ctx context.Context, stackName, outputKey string) (string, error) {
	resp, err := c.API.DescribeStacks(ctx, &cloudformation.DescribeStacksInput{
		StackName: &stackName,
	})
	if err != nil {
		return "", err
	}

	for _, stack := range resp.Stacks {
		for _, output := range stack.Outputs {
			if *output.OutputKey == outputKey {
				return *output.OutputValue, nil
			}
		}
	}

	return "", fmt.Errorf("output key '%s' not found in stack '%s'", outputKey, stackName)
}
