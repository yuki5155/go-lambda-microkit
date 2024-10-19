.PHONY: build

build:
	sam build

# Makefile for SAM deployment

# Cognito stack name - change this to your Cognito stack name
COGNITO_STACK_NAME := cognito

# Ensure that HOSTED_ZONE_ID and CUSTOM_DOMAIN are set
ifndef HOSTED_ZONE_ID
$(error HOSTED_ZONE_ID is not set)
endif

ifndef CUSTOM_DOMAIN
$(error CUSTOM_DOMAIN is not set)
endif

ifndef ALLOW_ORIGIN
$(error ALLOW_ORIGIN is not set)
endif

# Default target
.PHONY: deploy

deploy:
	sam deploy \
		--capabilities CAPABILITY_IAM \
		--parameter-overrides \
			CustomDomainName=$(CUSTOM_DOMAIN) \
			HostedZoneId=$(HOSTED_ZONE_ID) \
			CognitoStackName=$(COGNITO_STACK_NAME) \
			AllowedOrigins=$(ALLOW_ORIGIN) \
		--disable-rollback false