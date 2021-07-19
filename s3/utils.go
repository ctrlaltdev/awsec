package s3

import (
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

func bucketGrantAssign(data *bucketReportFormat, grantee string, perm types.Permission) {
	if perm == "FULL_CONTROL" {
		data.full = append(data.full, grantee)
	} else if perm == "WRITE" {
		data.write = append(data.write, grantee)
	} else if perm == "WRITE_ACP" {
		data.write_acp = append(data.write_acp, grantee)
	} else if perm == "READ" {
		data.read = append(data.read, grantee)
	} else if perm == "READ_ACP" {
		data.read_acp = append(data.read_acp, grantee)
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
