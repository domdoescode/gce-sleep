package cmd

import (
	"testing"
	"time"
)

func TestShouldBeRunning(t *testing.T) {
	tests := []struct {
		Name            string
		Now             time.Time
		Start           time.Time
		Stop            time.Time
		ShouldBeRunning bool
	}{
		{
			Name:            "Should be running (within hours, same day)",
			Now:             time.Date(2017, time.January, 14, 12, 0, 0, 0, time.UTC),
			Start:           time.Date(2017, time.January, 14, 11, 0, 0, 0, time.UTC),
			Stop:            time.Date(2017, time.January, 14, 13, 0, 0, 0, time.UTC),
			ShouldBeRunning: true,
		},
		{
			Name:            "Shouldn't be running (outside of hours, same day)",
			Now:             time.Date(2017, time.January, 14, 14, 0, 0, 0, time.UTC),
			Start:           time.Date(2017, time.January, 14, 11, 0, 0, 0, time.UTC),
			Stop:            time.Date(2017, time.January, 14, 13, 0, 0, 0, time.UTC),
			ShouldBeRunning: false,
		},
		{
			Name:            "Should be running (within hours, different day)",
			Now:             time.Date(2017, time.January, 15, 12, 0, 0, 0, time.UTC),
			Start:           time.Date(2017, time.January, 14, 11, 0, 0, 0, time.UTC),
			Stop:            time.Date(2017, time.January, 14, 13, 0, 0, 0, time.UTC),
			ShouldBeRunning: true,
		},
		{
			Name:            "Shouldn't be running (outside of hours, different day)",
			Now:             time.Date(2017, time.January, 15, 0, 0, 0, 0, time.UTC),
			Start:           time.Date(2017, time.January, 14, 11, 0, 0, 0, time.UTC),
			Stop:            time.Date(2017, time.January, 14, 13, 0, 0, 0, time.UTC),
			ShouldBeRunning: false,
		},
		{
			Name:            "Should be running (same start time)",
			Now:             time.Date(2017, time.January, 14, 11, 0, 0, 0, time.UTC),
			Start:           time.Date(2017, time.January, 14, 11, 0, 0, 0, time.UTC),
			Stop:            time.Date(2017, time.January, 14, 13, 0, 0, 0, time.UTC),
			ShouldBeRunning: true,
		},
		{
			Name:            "Should be running (same stop time)",
			Now:             time.Date(2017, time.January, 14, 11, 0, 0, 0, time.UTC),
			Start:           time.Date(2017, time.January, 14, 11, 0, 0, 0, time.UTC),
			Stop:            time.Date(2017, time.January, 14, 13, 0, 0, 0, time.UTC),
			ShouldBeRunning: true,
		},
		{
			Name:            "Shouldn't be running (minute before start time)",
			Now:             time.Date(2017, time.January, 14, 10, 59, 0, 0, time.UTC),
			Start:           time.Date(2017, time.January, 14, 11, 0, 0, 0, time.UTC),
			Stop:            time.Date(2017, time.January, 14, 13, 0, 0, 0, time.UTC),
			ShouldBeRunning: false,
		},
		{
			Name:            "Shouldn't be running (minute after stop time)",
			Now:             time.Date(2017, time.January, 14, 13, 1, 0, 0, time.UTC),
			Start:           time.Date(2017, time.January, 14, 11, 0, 0, 0, time.UTC),
			Stop:            time.Date(2017, time.January, 14, 13, 0, 0, 0, time.UTC),
			ShouldBeRunning: false,
		},
		{
			Name:            "Shouldn't be running (same start hour, before minutes)",
			Now:             time.Date(2017, time.January, 14, 11, 0, 0, 0, time.UTC),
			Start:           time.Date(2017, time.January, 14, 11, 30, 0, 0, time.UTC),
			Stop:            time.Date(2017, time.January, 14, 13, 0, 0, 0, time.UTC),
			ShouldBeRunning: false,
		},
		{
			Name:            "Should be running (same start hour, after minutes)",
			Now:             time.Date(2017, time.January, 14, 11, 31, 0, 0, time.UTC),
			Start:           time.Date(2017, time.January, 14, 11, 30, 0, 0, time.UTC),
			Stop:            time.Date(2017, time.January, 14, 13, 0, 0, 0, time.UTC),
			ShouldBeRunning: true,
		},
		{
			Name:            "Should be running (same stop hour, before minutes)",
			Now:             time.Date(2017, time.January, 14, 13, 0, 0, 0, time.UTC),
			Start:           time.Date(2017, time.January, 14, 11, 0, 0, 0, time.UTC),
			Stop:            time.Date(2017, time.January, 14, 13, 30, 0, 0, time.UTC),
			ShouldBeRunning: true,
		},
		{
			Name:            "Shouldn't be running (same stop hour, after minutes)",
			Now:             time.Date(2017, time.January, 14, 13, 31, 0, 0, time.UTC),
			Start:           time.Date(2017, time.January, 14, 11, 0, 0, 0, time.UTC),
			Stop:            time.Date(2017, time.January, 14, 13, 30, 0, 0, time.UTC),
			ShouldBeRunning: false,
		},
	}

	for _, test := range tests {
		if shouldBeRunning(test.Now, test.Start, test.Stop) != test.ShouldBeRunning {
			t.Error(test.Name)
		}
	}
}
