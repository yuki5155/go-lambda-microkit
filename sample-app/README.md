# Deployment Process

This project uses a Makefile to simplify the deployment process using AWS Serverless Application Model (SAM). The main deployment target is `deploy`, which can be run using the command `make deploy`.

## Prerequisites

Before running the deployment, ensure you have the following:

1. AWS SAM CLI installed and configured
2. Necessary environment variables set:
   - `HOSTED_ZONE_ID`: The ID of your AWS Route 53 hosted zone
   - `CUSTOM_DOMAIN`: The custom domain name for your application

## Deployment Target Explanation

The `deploy` target in the Makefile does the following:

1. Runs the `sam deploy` command with specific options:
   - `--capabilities CAPABILITY_IAM`: Allows SAM to create IAM roles during deployment
   - `--parameter-overrides`: Passes several parameters to the SAM template:
     - `CustomDomainName`: Set to the value of the `CUSTOM_DOMAIN` environment variable
     - `HostedZoneId`: Set to the value of the `HOSTED_ZONE_ID` environment variable
     - `CognitoStackName`: Set to "cognito" (defined at the top of the Makefile)
   - `--disable-rollback false`: Allows the stack to roll back in case of a deployment failure

## Usage

To deploy your application:

1. Set the required environment variables:
   ```
   export HOSTED_ZONE_ID=your_hosted_zone_id
   export CUSTOM_DOMAIN=your_custom_domain
   ```
2. Run the deployment command:
   ```
   make deploy
   ```

Note: The Makefile includes error checking to ensure that `HOSTED_ZONE_ID` and `CUSTOM_DOMAIN` are set before running the deployment. If either is missing, the deployment will fail with an error message.

## Customization

You can customize the deployment by modifying the following variables in the Makefile:
- `COGNITO_STACK_NAME`: Change this if your Cognito stack has a different name
- Add or modify parameter overrides in the `sam deploy` command as needed for your specific application

Remember to update your SAM template to use these parameters appropriately.