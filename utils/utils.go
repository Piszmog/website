package utils

import "time"

func FormatDate(dateStr string) string {
	if dateStr == "Present" {
		return dateStr
	}
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return ""
	}
	return date.Format("Jan 2006")
}
