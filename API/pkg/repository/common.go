package repository

import (
	"time"
)

func DateToSQLString(date time.Time) string {
	return date.Format("2006-01-02 15:04:05 -07:00")
}
