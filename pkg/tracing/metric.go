package tracing

import (
	"context"
	"log"
	"time"

	"github.com/rushteam/beauty/pkg/core"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdoutmetric"

	"go.opentelemetry.io/otel/metric"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
)

var meter metric.Meter

type MetricOption func(c *metricComponent)

func WithMetricReader(reader sdkmetric.Reader) MetricOption {
	return func(o *metricComponent) {
		o.options = append(o.options, sdkmetric.WithReader(reader))
	}
}

func WithMetricOption(opt sdkmetric.Option) MetricOption {
	return func(o *metricComponent) {
		o.options = append(o.options, opt)
	}
}

func WithMetricProvider(provider metric.MeterProvider) MetricOption {
	return func(o *metricComponent) {
		o.provider = provider
	}
}

func WithMetricStdoutReader() MetricOption {
	exporter, err := stdoutmetric.New(
		stdoutmetric.WithPrettyPrint(),
	)
	if err != nil {
		log.Fatal(err)
	}
	return WithMetricReader(sdkmetric.NewPeriodicReader(
		exporter,
		sdkmetric.WithInterval(5*time.Second), // default 1m, for test 5s
	))
}

type metricComponent struct {
	provider metric.MeterProvider
	options  []sdkmetric.Option
}

func (c *metricComponent) Name() string {
	return "metric"
}

func (c *metricComponent) Init() context.CancelFunc {
	if c.provider == nil {
		meterProvider := sdkmetric.NewMeterProvider(
			// metric.WithResource(res),
			c.options...,
		)
		otel.SetMeterProvider(meterProvider)
		return func() {
			_ = meterProvider.Shutdown(context.Background())
		}
	}
	otel.SetMeterProvider(c.provider)
	return func() {}
}

func NewMetric(opts ...MetricOption) core.Component {
	c := &metricComponent{}
	if len(opts) == 0 {
		opts = append(opts, WithMetricStdoutReader())
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

func Meter() metric.Meter {
	if meter == nil {
		meter = otel.GetMeterProvider().Meter("beauty")
	}
	return meter
}