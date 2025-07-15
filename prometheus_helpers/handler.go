package prometheus_helpers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/MordaTeam/go-toolbox/options"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Handler struct {
	http.Handler
	registry *prometheus.Registry
}

// Creates new Handler with given opts.
func NewHandler(opts ...options.Option[Handler]) (*Handler, error) {
	h := Handler{
		registry: prometheus.NewRegistry(),
	}

	if err := options.ApplyOptions(&h, opts...); err != nil {
		return nil, fmt.Errorf("applying opts: %w", err)
	}

	h.Handler = promhttp.HandlerFor(
		h.registry,
		promhttp.HandlerOpts{
			ErrorHandling: promhttp.ContinueOnError,
		},
	)

	return &h, nil
}

// Sets custom registry to Handler.
// Creates new by default.
func HandlerWithRegistry(r *prometheus.Registry) options.Option[Handler] {
	return func(target *Handler) error {
		if r == nil {
			return errors.New("got nil registry")
		}

		target.registry = r
		return nil
	}
}

// Adds collectors to handler registry, so they will be returned by further ServeHTTP() calls.
func (h *Handler) Register(cols ...prometheus.Collector) error {
	for idx, col := range cols {
		if err := h.registry.Register(col); err != nil {
			return fmt.Errorf("registering collector %d: %w", idx, err)
		}
	}
	return nil
}
