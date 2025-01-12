package time

import "time"

// GetTimeGMT7 retrieves the current time in the GMT+7.
func (t *Time) GetTimeGMT7() time.Time {
	location, _ := time.LoadLocation("Asia/Jakarta")
	return time.Now().In(location)
}
