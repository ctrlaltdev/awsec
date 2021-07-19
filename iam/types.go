package iam

import "time"

type accessKeyReportFormat struct {
	id        *string
	createdAt *time.Time
	active    bool
}

type userReportFormat struct {
	name             *string
	passwordLastUsed *time.Time
	keys             []accessKeyReportFormat
}
