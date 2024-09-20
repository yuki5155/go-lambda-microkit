require github.com/aws/aws-lambda-go v1.36.1

replace gopkg.in/yaml.v2 => gopkg.in/yaml.v2 v2.2.8

require (
	github.com/golang/mock v1.6.0
	github.com/yuki5155/go-lambda-microkit v0.0.0-unpublished
)

require (
	golang.org/x/mod v0.4.2 // indirect
	golang.org/x/sys v0.0.0-20210510120138-977fb7262007 // indirect
	golang.org/x/tools v0.1.1 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
)

replace github.com/yuki5155/go-lambda-microkit => ../

module hello-world

go 1.19
