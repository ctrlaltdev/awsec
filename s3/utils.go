package s3

import (
	t "awsec/s3/types"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

func bucketGrantAssign(data *t.BucketReportFormat, grantee string, perm types.Permission) {
	if perm == "FULL_CONTROL" {
		data.Full = append(data.Full, grantee)
	} else if perm == "WRITE" {
		data.Write = append(data.Write, grantee)
	} else if perm == "WRITE_ACP" {
		data.Write_acp = append(data.Write_acp, grantee)
	} else if perm == "READ" {
		data.Read = append(data.Read, grantee)
	} else if perm == "READ_ACP" {
		data.Read_acp = append(data.Read_acp, grantee)
	}
}

func bucketGranteeAssign(data []string, grantee *types.Grantee) []string {
	if grantee.DisplayName != nil {
		data = append(data, *grantee.DisplayName)
	}
	if grantee.EmailAddress != nil {
		data = append(data, *grantee.EmailAddress)
	}
	// if grantee.ID != nil {
	// 	data = append(data, *grantee.ID)
	// }
	if grantee.URI != nil {
		data = append(data, strings.Replace(*grantee.URI, "http://acs.amazonaws.com/groups/", "", -1))
	}

	return data
}
