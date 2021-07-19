package types

import "time"

type AccessKeyReportFormat struct {
	ID        *string
	CreatedAt *time.Time
	Active    bool
}

type UserReportFormat struct {
	Name             *string
	PasswordLastUsed *time.Time
	Keys             []AccessKeyReportFormat
}
