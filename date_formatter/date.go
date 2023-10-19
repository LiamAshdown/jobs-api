package date_formatter

import "time"

func FormatDate(date string) string {
	time, err := time.Parse("2006-01-02 15:04:05", date)

	if err != nil {
		panic(err)
	}

	return time.Format("2 Jan 06")
}
