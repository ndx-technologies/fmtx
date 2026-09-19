package fmtx

import (
	"testing"
	"time"
)

func TestDurationS(t *testing.T) {
	tests := map[time.Duration]string{
		0:                       "0s",
		-time.Second:            "0s",
		20 * time.Millisecond:   "0s", // under 100ms rounds down
		100 * time.Millisecond:  "100ms",
		400 * time.Millisecond:  "400ms",
		2700 * time.Millisecond: "2.7s",
		5 * time.Second:         "5s",
		59 * time.Second:        "59s",
		63 * time.Second:        "1m3s",
		125 * time.Second:       "2m5s",
		time.Hour:               "1h0m",
		63 * time.Minute:        "1h3m",
		25 * time.Hour:          "1d1h",
		49 * time.Hour:          "2d1h",
	}

	for d, s := range tests {
		t.Run(d.String(), func(t *testing.T) {
			if got := Duration(d); got != s {
				t.Errorf("DurationS(%s) = %q, want %q", d, got, s)
			}
		})
	}
}

func TestOffsetS(t *testing.T) {
	start := time.Date(2026, 9, 18, 17, 49, 54, 0, time.UTC)

	tests := map[time.Duration]string{
		0:                       "00:00.000",
		232 * time.Millisecond:  "00:00.232",
		2990 * time.Millisecond: "00:02.990",
		3*time.Minute + 8*time.Second + 838*time.Millisecond: "03:08.838",
		12*time.Minute + 30*time.Second:                      "12:30.000",
		63 * time.Minute:                                     "63:00.000", // minutes are not wrapped
	}

	for offset, s := range tests {
		t.Run(offset.String(), func(t *testing.T) {
			if got := TimeOffset(start.Add(offset), start); got != s {
				t.Error(offset, s, got)
			}
		})
	}

	if s := TimeOffset(start.Add(-time.Second), start); s != "00:00.000" {
		t.Error(s)
	}
}
