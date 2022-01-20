package utils

import "time"

func FormatTime(timeStamp int64) string {
	timestr := time.Unix(timeStamp, 0).Format("2006-01-02 15:04:05")
	return timestr
}
