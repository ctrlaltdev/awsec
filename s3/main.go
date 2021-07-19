package s3

import (
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
		fmt.Fprintf(w, format, *report.name, strings.Join(report.full, ", "), strings.Join(report.write_acp, ", "), strings.Join(report.read_acp, ", "), strings.Join(report.write, ", "), strings.Join(report.read, ", "), report.versioning, strings.Join(report.encryption, ", "))
	}

	defer w.Flush()
}

func Check() (s3Reports []bucketReportFormat) {
	buckets := ListBuckets()

	for _, bucket := range buckets {

		report := bucketReportFormat{}
		report.name = bucket.Name
		loc := GetBucketLocation(bucket.Name)
		report.region = &loc.LocationConstraint

		vers := GetBucketVersioning(report.name, string(*report.region))
		if vers.Status == "Enabled" {
			report.versioning = "enabled"
		} // else {
		// 	report.versioning = "disabled"
		// }

		enc := GetBucketEncryption(report.name, string(*report.region))
		if enc.ServerSideEncryptionConfiguration != nil {
			for _, rule := range enc.ServerSideEncryptionConfiguration.Rules {
				if rule.ApplyServerSideEncryptionByDefault != nil {
					report.encryption = append(report.encryption, string(rule.ApplyServerSideEncryptionByDefault.SSEAlgorithm))
				}
			}
		}

		acl := GetBucketACL(report.name, string(*report.region))
		if acl.Owner.DisplayName != nil {
			report.owner = acl.Owner
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
