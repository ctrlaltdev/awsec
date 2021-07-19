package types

import "github.com/aws/aws-sdk-go-v2/service/s3/types"

type BucketReportFormat struct {
	Name       *string
	Region     *types.BucketLocationConstraint
	Owner      *types.Owner
	Full       []string
	Write_acp  []string
	Read_acp   []string
	Write      []string
	Read       []string
	Versioning string
	Encryption []string
}
