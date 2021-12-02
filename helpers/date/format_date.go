package format_date

import "time"

const (
	LayoutISO = "2006-01-02T15:04:05.999999999Z07:00"
	layoutUS  = "Monday, January 2, 2006, 03:04:01 PM"
)

func FormatDate(t string) string {
	v, _ := time.Parse(LayoutISO, t)
	return v.Format(layoutUS)
}
