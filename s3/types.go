package s3

import "github.com/aws/aws-sdk-go-v2/service/s3/types"

type bucketReportFormat struct {
	name       *string
	region     *types.BucketLocationConstraint
	owner      *types.Owner
	full       []string
	write_acp  []string
	read_acp   []string
	write      []string
	read       []string
	versioning string
	encryption []string
}
