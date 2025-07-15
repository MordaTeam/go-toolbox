package prometheus_helpers

import (
	"errors"
	"fmt"

	"github.com/MordaTeam/go-toolbox/options"
	"github.com/prometheus/client_golang/prometheus"
)

// Creates new summary opts with given name and, maybe, auxOpts.
func NewSummaryOpts(name string, auxOpts ...options.Option[prometheus.SummaryOpts]) (*prometheus.SummaryOpts, error) {
	sumOpts := prometheus.SummaryOpts{
		Name: name,
		Objectives: map[float64]float64{
			0.5:  0.01,
			0.75: 0.01,
			0.95: 0.001,
			0.99: 0.001,
		},
	}

	if err := options.ApplyOptions(&sumOpts, auxOpts...); err != nil {
		return nil, fmt.Errorf("applying opts: %w", err)
	}

	return &sumOpts, nil
}

// Sets custom namespace for summary.
// Default is empty.
func SummaryOptsWithNamespace(ns string) options.Option[prometheus.SummaryOpts] {
	return func(target *prometheus.SummaryOpts) error {
		if ns == "" {
			return errors.New("got empty namespace")
		}
		target.Namespace = ns
		return nil
	}
}

// Sets custom subsystem for summary.
// Default is empty.
func SummaryOptsWithSubsystem(ss string) options.Option[prometheus.SummaryOpts] {
	return func(target *prometheus.SummaryOpts) error {
		if ss == "" {
			return errors.New("got empty subsystem")
		}
		target.Subsystem = ss
		return nil
	}
}

// Sets custom help for summary.
// Default is empty.
func SummaryOptsWithHelp(h string) options.Option[prometheus.SummaryOpts] {
	return func(target *prometheus.SummaryOpts) error {
		if h == "" {
			return errors.New("got empty help")
		}
		target.Help = h
		return nil
	}
}

// Sets custom objectives for summary.
//
//	Default is map[float64]float64{
//		0.5:  0.01,
//		0.75: 0.01,
//		0.95: 0.001,
//		0.99: 0.001,
//	}
func SummaryOptsWithObjectives(o map[float64]float64) options.Option[prometheus.SummaryOpts] {
	return func(target *prometheus.SummaryOpts) error {
		if len(o) == 0 {
			return errors.New("got nil or empty objectives")
		}
		target.Objectives = o
		return nil
	}
}
