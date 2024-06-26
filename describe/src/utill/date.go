package util

import "time"

func GetDate(str string) (date time.Time, err error) {
	layout := "2006-01-02"
	date, err = time.Parse(layout, str)
	return
}

func GetAge(date time.Time) (age time.Duration) {
	now := time.Now()
	age = now.Sub(date)
	return age
}

func DurationToFloat64(date time.Duration) (age float64) {
	return float64(date)
}
