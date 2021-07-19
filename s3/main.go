package s3

import (
	"awsec/s3/types"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var (
	cfg    *aws.Config
	client *s3.Client
)

func Init(config *aws.Config) {
	cfg = config
	client = s3.NewFromConfig(*cfg)
}

func Report() {
	fmt.Printf(`

┏━━━━━┓
┃ S3  ┃
┗━━━━━┛`)

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 8, 8, 2, ' ', 0)

	format := "\n %s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t"

	fmt.Fprintf(w, format, "BUCKET", "FULL CONTROL", "WRITE ACP", "READ ACP", "WRITE", "READ", "VERSIONING", "ENCRYPTION")
	fmt.Fprintf(w, format, "------", "-----", "-----", "-----", "-----", "-----", "-----", "-----")

	reports := Check()

	for _, report := range reports {
		fmt.Fprintf(w, format, *report.Name, strings.Join(report.Full, ", "), strings.Join(report.Write_acp, ", "), strings.Join(report.Read_acp, ", "), strings.Join(report.Write, ", "), strings.Join(report.Read, ", "), report.Versioning, strings.Join(report.Encryption, ", "))
	}

	defer w.Flush()
}

func Check() (s3Reports []types.BucketReportFormat) {
	buckets := ListBuckets()

	for _, bucket := range buckets {

		report := types.BucketReportFormat{}
		report.Name = bucket.Name
		loc := GetBucketLocation(bucket.Name)
		report.Region = &loc.LocationConstraint

		vers := GetBucketVersioning(report.Name, string(*report.Region))
		if vers.Status == "Enabled" {
			report.Versioning = "enabled"
		} // else {
		// 	report.versioning = "disabled"
		// }

		enc := GetBucketEncryption(report.Name, string(*report.Region))
		if enc.ServerSideEncryptionConfiguration != nil {
			for _, rule := range enc.ServerSideEncryptionConfiguration.Rules {
				if rule.ApplyServerSideEncryptionByDefault != nil {
					report.Encryption = append(report.Encryption, string(rule.ApplyServerSideEncryptionByDefault.SSEAlgorithm))
				}
			}
		}

		acl := GetBucketACL(report.Name, string(*report.Region))
		if acl.Owner.DisplayName != nil {
			report.Owner = acl.Owner
		}

		for _, grant := range acl.Grants {
			grantee := []string{}
			// if grant.Grantee.Type == "CanonicalUser" {
			// 	grantee = bucketGranteeAssign(grantee, grant.Grantee)
			// 	bucketGrantAssign(&report, strings.Join(grantee, " "), grant.Permission)
			// }
			if grant.Grantee.Type == "Group" {
				grantee = bucketGranteeAssign(grantee, grant.Grantee)
				bucketGrantAssign(&report, strings.Join(grantee, " "), grant.Permission)
			}
		}

		// var owner string

		// if report.owner != nil {
		// 	owner = *report.owner.DisplayName
		// } else {
		// 	owner = "NONE"
		// }

		s3Reports = append(s3Reports, report)
	}

	return s3Reports
}
