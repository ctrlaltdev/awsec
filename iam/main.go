package iam

import (
	"awsec/iam/types"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"

	"github.com/ctrlaltdev/justamin"
)

var (
	cfg    *aws.Config
	client *iam.Client
)

func Init(config *aws.Config) {
	cfg = config
	client = iam.NewFromConfig(*cfg)
}

func Report() {
	fmt.Printf(`

┏━━━━━┓
┃ IAM ┃
┗━━━━━┛`)

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 8, 8, 2, ' ', 0)

	format := "\n %s\t%s\t%s\t"

	fmt.Fprintf(w, format, "USER", "LAST LOGIN", "ACCESS KEYS AGE")
	fmt.Fprintf(w, format, "------", "-----", "-----")

	reports := Check()

	for _, report := range reports {
		var keysAge []string
		for _, key := range report.Keys {
			keysAge = append(keysAge, justamin.Duration(*key.CreatedAt))
		}

		var lastConn string
		if report.PasswordLastUsed != nil {
			lastConn = justamin.Duration(*report.PasswordLastUsed)
		}

		fmt.Fprintf(w, format, *report.Name, lastConn, strings.Join(keysAge, ", "))
	}

	defer w.Flush()
}

func Check() (userReports []types.UserReportFormat) {
	users := ListUsers()

	for _, user := range users {

		var report types.UserReportFormat

		report.Name = user.UserName
		report.PasswordLastUsed = user.PasswordLastUsed

		keys := ListAccessKeys(user.UserName)

		for _, key := range keys {

			var keyReport types.AccessKeyReportFormat

			keyReport.ID = key.AccessKeyId
			keyReport.CreatedAt = key.CreateDate

			if key.Status == "Active" {
				keyReport.Active = true
			}

			report.Keys = append(report.Keys, keyReport)

		}

		userReports = append(userReports, report)

	}

	return userReports
}
