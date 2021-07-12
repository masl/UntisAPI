package Untis

import "time"

func ToUntisDate(time time.Time) int {
	year, month, day := time.Date()
	value := year*10000 + int(month)*100 + day
	return value
}

func ToGoDate(value int) time.Time {
	year := value / 10000
	month := (value / 100) - year*100
	day := value - month*100 - year*10000
	time := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.FixedZone("UTC+2", 2*60*60))
	return time
}

func ToUnitsTime(time time.Time) int {
	return 0
}

func ToGoTime(value int) time.Time {
	houre := value / 100
	minute := value - (houre * 100)
	time := time.Date(0, time.Month(0), 0, houre, minute, 0, 0, time.FixedZone("UTC+2", 2*60*60))
	return time
}
