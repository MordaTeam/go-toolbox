package timex_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/MordaTeam/go-toolbox/timex"
)

func TestInterval_Unmarshal(t *testing.T) {
	testCases := []struct {
		Input    string
		Expected timex.Interval
		IsErr    bool
	}{
		// Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
		{Input: `"1ns"`, Expected: timex.Interval(time.Nanosecond)},
		{Input: `"1us"`, Expected: timex.Interval(time.Microsecond)},
		{Input: `"1ms"`, Expected: timex.Interval(time.Millisecond)},
		{Input: `"1s"`, Expected: timex.Interval(time.Second)},
		{Input: `"1m"`, Expected: timex.Interval(time.Minute)},
		{Input: `"1h"`, Expected: timex.Interval(time.Hour)},
		// Invalid time units are not allowed.
		{Input: `"1x"`, IsErr: true},
		{Input: `"1"`, IsErr: true},
		{Input: `""`, IsErr: true},
		{Input: `"x"`, IsErr: true},
		{Input: `"1d"`, IsErr: true},
		{Input: `"1w"`, IsErr: true},
		{Input: `"1y"`, IsErr: true},
		{Input: `"1f"`, IsErr: true},
		{Input: `1`, IsErr: true},
		{Input: `1.1`, IsErr: true},
		{Input: `true`, IsErr: true},
		{Input: `false`, IsErr: true},
		{Input: `{}`, IsErr: true},
		{Input: `[]`, IsErr: true},
		{Input: `null`, IsErr: true},
	}

	for _, tc := range testCases {
		t.Run(tc.Input, func(t *testing.T) {
			var ivl timex.Interval
			err := json.Unmarshal([]byte(tc.Input), &ivl)

			if tc.IsErr {
				if err == nil {
					t.Errorf("expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if ivl != tc.Expected {
				t.Errorf("expected %v, got %v", tc.Expected, ivl)
			}
		})
	}
}
