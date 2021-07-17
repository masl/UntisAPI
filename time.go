package UntisAPI

import (
	"strconv"
	"time"
)

/*
ToUntisDate takes a time.Time and returns the Date formatted in the Untis way (yyyymmdd).
*/
func ToUntisDate(time time.Time) int {
	year, month, day := time.Date()
	value := year*10000 + int(month)*100 + day
	return value
}

/*
ToGoDate takes a time in the Untis way (yyyymmdd) and returns the a time.Time.
*/
func ToGoDate(value int) time.Time {
	year := value / 10000
	month := (value / 100) - year*100
	day := value - month*100 - year*10000
	date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.FixedZone("UTC+2", 2*60*60)) // TODO switch to daylight saving time
	return date
}

/*
ToUnitsTime takes a time.Time and returns the time formatted in the Untis way (hhmm).
*/
func ToUnitsTime(time time.Time) int {
	value, _ := strconv.Atoi(time.Format("1504"))
	return value
}

/*
ToGoTime takes a time in the Untis way (hhmm) and returns the a time.Time.
*/
func ToGoTime(value int) time.Time {
	hour := value / 100
	minute := value - (hour * 100)
	timeVar := time.Date(0, time.Month(0), 0, hour, minute, 0, 0, time.FixedZone("UTC+2", 2*60*60))
	return timeVar
}
