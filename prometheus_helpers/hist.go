package prometheus_helpers

import (
	"errors"
	"fmt"
	"time"

	"github.com/MordaTeam/go-toolbox/options"
	"github.com/prometheus/client_golang/prometheus"
)

// Creates new histogram opts with given name and, maybe, auxOpts.
func NewHistOpts(name string, auxOpts ...options.Option[prometheus.HistogramOpts]) (*prometheus.HistogramOpts, error) {
	histOpts := prometheus.HistogramOpts{
		Name: name,
		Buckets: []float64{
			float64(time.Microsecond * 100),
			float64(time.Millisecond),
			float64(time.Millisecond * 5),
			float64(time.Millisecond * 10),
			float64(time.Millisecond * 20),
			float64(time.Millisecond * 50),
			float64(time.Millisecond * 100),
		},
	}

	if err := options.ApplyOptions(&histOpts, auxOpts...); err != nil {
		return nil, fmt.Errorf("applying opts: %w", err)
	}

	return &histOpts, nil
}

// Sets custom namespace for histogram.
// Default is empty.
func HistOptsWithNamespace(ns string) options.Option[prometheus.HistogramOpts] {
	return func(target *prometheus.HistogramOpts) error {
		target.Namespace = ns
		return nil
	}
}

// Sets custom subsystem for histogram.
// Default is empty.
func HistOptsWithSubsystem(ss string) options.Option[prometheus.HistogramOpts] {
	return func(target *prometheus.HistogramOpts) error {
		target.Subsystem = ss
		return nil
	}
}

// Sets custom help for histogram.
// Default is empty.
func HistOptsWithHelp(h string) options.Option[prometheus.HistogramOpts] {
	return func(target *prometheus.HistogramOpts) error {
		target.Help = h
		return nil
	}
}

// Sets custom buckets for histogram.
// Default is [100µs, 1ms, 5ms, 10ms, 20ms, 50ms, 100ms].
func HistOptsWithBuckets(b []float64) options.Option[prometheus.HistogramOpts] {
	return func(target *prometheus.HistogramOpts) error {
		if len(b) == 0 {
			return errors.New("got nil or empty buckets")
		}

		target.Buckets = b
		return nil
	}
}
