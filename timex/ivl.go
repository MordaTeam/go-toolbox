// Extension of time types.
package timex

import (
	"encoding/json"
	"fmt"
	"time"
)

var _ json.Unmarshaler = new(Interval)

// Interval represents a time interval and implements json.Unmarshaler.
// It is a wrapper around time.Duration.
type Interval time.Duration

func (i *Interval) Duration() time.Duration {
	return time.Duration(*i)
}

//nolint:wrapcheck // return default json.Marshal error
func (i *Interval) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Duration(*i).String())
}

func (i *Interval) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return fmt.Errorf("invalid interval: %w", err)
	}

	d, err := time.ParseDuration(s)
	if err != nil {
		return fmt.Errorf("invalid interval: %w", err)
	}

	*i = Interval(d)

	return nil
}
