package s3

import (
	"context"
	"errors"
	"log"

	awshttp "github.com/aws/aws-sdk-go-v2/aws/transport/http"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

func ListBuckets() []types.Bucket {
	result, err := client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})
	if err != nil {
		log.Println(err)
	}

	return result.Buckets
}

func GetBucketLocation(bucket *string) s3.GetBucketLocationOutput {
	result, err := client.GetBucketLocation(context.TODO(), &s3.GetBucketLocationInput{Bucket: bucket})

	if err != nil {
		log.Printf("%+v\n", err)
	}

	return *result
}

func GetBucketACL(bucket *string, location string) s3.GetBucketAclOutput {
	client := client
	if location != "" {
		client = s3.NewFromConfig(*cfg, func(o *s3.Options) {
			o.Region = location
		})
	}
	result, err := client.GetBucketAcl(context.TODO(), &s3.GetBucketAclInput{Bucket: bucket})

	if err != nil {
		log.Printf("%+v\n", err)
	}

	return *result
}

func GetBucketVersioning(bucket *string, location string) s3.GetBucketVersioningOutput {
	client := client
	if location != "" {
		client = s3.NewFromConfig(*cfg, func(o *s3.Options) {
			o.Region = location
		})
	}
	result, err := client.GetBucketVersioning(context.TODO(), &s3.GetBucketVersioningInput{Bucket: bucket})

	if err != nil {
		log.Printf("%+v\n", err)
	}

	return *result
}

func GetBucketEncryption(bucket *string, location string) s3.GetBucketEncryptionOutput {
	client := client
	if location != "" {
		client = s3.NewFromConfig(*cfg, func(o *s3.Options) {
			o.Region = location
		})
	}
	result, err := client.GetBucketEncryption(context.TODO(), &s3.GetBucketEncryptionInput{Bucket: bucket})

	if err != nil {
		var re *awshttp.ResponseError
		if errors.As(err, &re) {
			if re.HTTPStatusCode() == 404 {
				return s3.GetBucketEncryptionOutput{}
			}
		}

		log.Printf("%+v\n", err)
	}

	return *result
}
