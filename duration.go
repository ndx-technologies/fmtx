package fmtx

import (
	"strconv"
	"time"
)

// Duration renders a duration for a table cell, keeping the unit count small:
// 400ms, 2.7s, 1m3s, 1h2m, 1d1h. Zero and negative durations render as 0s, and
// anything under 100ms rounds to 0s.
func Duration(d time.Duration) string {
	if d <= 0 {
		return "0s"
	}
	switch {
	case d >= 24*time.Hour:
		d = d.Round(time.Hour)
		return strconv.Itoa(int(d.Hours())/24) + "d" + strconv.Itoa(int(d.Hours())%24) + "h"
	case d >= time.Hour:
		d = d.Round(time.Minute)
		return strconv.Itoa(int(d.Hours())) + "h" + strconv.Itoa(int(d.Minutes())%60) + "m"
	case d >= time.Minute:
		d = d.Round(time.Second)
		return strconv.Itoa(int(d.Minutes())) + "m" + strconv.Itoa(int(d.Seconds())%60) + "s"
	default:
		return d.Round(100 * time.Millisecond).String()
	}
}

// TimeOffset is the time from start as 00:00.000, for event timelines. Timestamps
// before start clamp to zero. Minutes are not wrapped, so an offset past the hour
// grows to 60:00.000 and beyond instead of becoming 01:00:00.000.
func TimeOffset(ts, start time.Time) string {
	d := max(ts.Sub(start), 0)
	m := int(d.Minutes())
	s := int(d.Seconds()) % 60
	ms := int(d.Milliseconds()) % 1000
	return strconv.Itoa(m/10) + strconv.Itoa(m%10) + ":" +
		strconv.Itoa(s/10) + strconv.Itoa(s%10) + "." +
		strconv.Itoa(ms/100) + strconv.Itoa((ms/10)%10) + strconv.Itoa(ms%10)
}
