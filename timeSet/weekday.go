package timeSet

import "time"

func WeekdayFunc(t time.Weekday) string {
	switch t {
	case time.Sunday:
		return "日"
	case time.Monday:
		return "月"
	case time.Tuesday:
		return "火"
	case time.Wednesday:
		return "水"
	case time.Thursday:
		return "木"
	case time.Friday:
		return "金"
	case time.Saturday:
		return "土"
	default:
		return "error"
	}
}
