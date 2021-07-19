package iam

import (
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
		for _, key := range report.keys {
			keysAge = append(keysAge, justamin.Duration(*key.createdAt))
		}

		var lastConn string
		if report.passwordLastUsed != nil {
			lastConn = justamin.Duration(*report.passwordLastUsed)
		}

		fmt.Fprintf(w, format, *report.name, lastConn, strings.Join(keysAge, ", "))
	}

	defer w.Flush()
}

func Check() (userReports []userReportFormat) {
	users := ListUsers()

	for _, user := range users {

		var report userReportFormat

		report.name = user.UserName
		report.passwordLastUsed = user.PasswordLastUsed

		keys := ListAccessKeys(user.UserName)

		for _, key := range keys {

			var keyReport accessKeyReportFormat

			keyReport.id = key.AccessKeyId
			keyReport.createdAt = key.CreateDate

			if key.Status == "Active" {
				keyReport.active = true
			}

			report.keys = append(report.keys, keyReport)

		}

		userReports = append(userReports, report)

	}

	return userReports
}
