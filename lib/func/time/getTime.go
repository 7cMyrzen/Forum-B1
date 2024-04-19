package time

import "time"

func GetCurrentTime() string {
	newTime := time.Now()
	formattedTime := newTime.Format("2006-01-02 15:04:05")
	return formattedTime
}

func FormatTime(t string) string {
	return t[8:10] + "/" + t[5:7] + "/" + t[0:4]
}
