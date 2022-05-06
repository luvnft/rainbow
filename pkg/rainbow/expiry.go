package rainbow

import (
	"time"
)

// IsExpiryAvailable takes the expiries from "func Expiries(" below and compare it with the option expiry
func IsExpiryAvailable(expiries []time.Time, expiry time.Time) bool {

	for _, e := range expiries {
		if expiry.Equal(e) {
			return true
		}
	}

	return false
}

// return today
// tomorrow
// next n fridays
// hour is the hour at which the option expires:
// Deribit is 8:00 UTC (hour=8)
// Delta is 12:00 UTC (hour=12)
// there is going to be double but it doesn't matter in the end
// if the expiry of an option is present, it will be fetched anyway and won't circle back to that expiry
// i.e it won't fetch twice the same option
func Expiries(hour int) []time.Time {
	expiries := make([]time.Time, 0, 5)

	today, tomorrow := TodayTomorrow(hour)
	expiries = append(expiries, today, tomorrow)
	//we start looking for Fridays on the day after tomorrow ;-)
	expiries = append(expiries, NextNFridays(today, 3)...)
	expiries = append(expiries, LastFridayOfEachQuarter(today, hour)...)

	return expiries

}
func TodayTomorrow(hour int) (time.Time, time.Time) {
	t := time.Now()

	today := time.Date(t.Year(), t.Month(), t.Day(), hour, 0, 0, 0, time.UTC)
	tomorrow := today.Add(24 * time.Hour)
	return today, tomorrow
}

// NextNFridays returns the next n Fridays
// t included if it's a Friday
func NextNFridays(t time.Time, fridays int) []time.Time {
	expiries := make([]time.Time, fridays)
	// adding "7" if not it counts the friday to the current week, not the next one
	d := t.Add(time.Duration(7-t.Weekday()+time.Friday) * 24 * time.Hour)
	for i := 0; i < fridays; i++ {
		expiries[i] = d.Add(time.Duration(i) * 7 * 24 * time.Hour)
	}
	return expiries
}

// function inspired by
// https://www.rosettacode.org/wiki/Last_Friday_of_each_month#Go
func LastFridayOfEachQuarter(t time.Time, hour int) []time.Time {
	y := t.Year()
	expiries := make([]time.Time, 0, 4)

	i := 0
	for m := t.Month(); i < 12; m++ {
		if m%3 == 0 {
			d := time.Date(y, m+1, 1, hour, 0, 0, 0, time.UTC).Add(-24 * time.Hour)
			d = d.Add(-time.Duration((d.Weekday()+7-time.Friday)%7) * 24 * time.Hour)
			expiries = append(expiries, d)

		}
		i++
	}
	return expiries
}
