package main

import (
	"awsec/iam"
	"awsec/s3"
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

var (
	cfg     aws.Config
	region  = flag.String("region", "", "AWS Region")
	profile = flag.String("profile", "", "AWS Profile")
)

func initConfig() aws.Config {
	var err error

	if *profile != "" && *region != "" {
		cfg, err = config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile(*profile), config.WithRegion(*region))
	} else if *region != "" {
		cfg, err = config.LoadDefaultConfig(context.TODO(), config.WithRegion(*region))
	} else if *profile != "" {
		cfg, err = config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile(*profile))
	}

	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	return cfg
}

func main() {
	flag.Parse()

	fmt.Printf(`
╔═══════╗
║ AWSEC ║
╚═══════╝`)

	initConfig()

	s3.Init(&cfg)
	iam.Init(&cfg)

	s3.Report()
	iam.Report()
}
