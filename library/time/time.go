package timeFun

import "time"

const FormatTimeString = "2006 01-02 15:04:05"

func TimeIntToString(num int64) string {
	return time.Unix(num, 0).Format(FormatTimeString)
}
