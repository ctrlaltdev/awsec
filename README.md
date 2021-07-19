# Basic AWS Security Assessment

Will check some common security configurations.
So far: check S3 buckets and IAM user settings.

## TODO: Installation

```sh
brew install ctrlaltdev/tap/awsec
```
or 
```sh
brew tap ctrlaltdev/tap
brew install awsec
```

## Standalone: How To Use

The tool needs AWS credentials to work, and will look into common places automatically (~/.aws/config, ~/.aws/credentials, env vars)

You can specify a default region in your ~/.aws/config or using the -region flag
You can use an AWS profile using the -profile flag

```sh
awsec -profile prod -region us-west-2
```

## Module: How To Use

You can use only specific subpart of the tool as modules:

```sh
go get -u github.com/ctrlaltdev/awsec/s3
```
or
```sh
go get -u github.com/ctrlaltdev/awsec/iam
```

In those cases, you need to initialize the module with an aws config:
```go
package main

import (
  "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
  "github.com/ctrlaltdev/awsec/iam"
  "github.com/ctrlaltdev/awsec/s3"
)

var cfg aws.Config

func main() {
  var err error
  cfg, err = config.LoadDefaultConfig(context.TODO())

  if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

  s3.Init(&cfg)
	iam.Init(&cfg)

  s3Reports := s3.Check()
  iamReports := iam.Check()

  // ...
}
```
