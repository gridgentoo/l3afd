// Copyright Contributors to the L3AF Project.
// SPDX-License-Identifier: Apache-2.0

package stats

/*
	TODO:
		- Remove dependent modules
			"github.com/prometheus/client_golang/prometheus"
			"github.com/prometheus/client_golang/prometheus/promauto"
			"github.com/prometheus/client_golang/prometheus/promhttp"
*/
import (
	"context"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/prometheus"
	api "go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/sdk/metric"

	"github.com/rs/zerolog/log"
)

var (
	NFStartCount  *CounterValue
	NFStopCount   *CounterValue
	NFUpdateCount *CounterValue
	NFRunning     *GaugeValue
	NFStartTime   *GaugeValue
	NFMonitorMap  *GaugeValue
)

func SetupMetrics(hostname, daemonName, metricsAddr string) {
	ctx := context.Background()

	exporter, err := prometheus.New()
	if err != nil {
		log.Fatal().Err(err).Msgf("Failed to create Prometheus exporter")
	}
	provider := metric.NewMeterProvider(metric.WithReader(exporter))
	meter := provider.Meter("l3af-project/l3afd")

	baseAttribs := []attribute.KeyValue { attribute.Key("host").String(hostname) }

	metricName := daemonName + "_NFStartCount"
	startCount, err := meter.Int64Counter(metricName, api.WithDescription("The count of network functions started"))
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create NFStartCount")
	}
	NFStartCount = NewCounterValue(ctx, &startCount, metricName, baseAttribs)

	metricName = daemonName + "_NFStopCount"
	stopCount, err := meter.Int64Counter(metricName, api.WithDescription("The count of network functions stopped"))
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create NFStartCount")
	}
	NFStopCount = NewCounterValue(ctx, &stopCount, metricName, baseAttribs)

	metricName = daemonName + "_NFUpdateCount"
	updateCount, err := meter.Int64Counter(metricName, api.WithDescription("The count of network functions updated"))
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create NFUpdateCount")
	}
	NFUpdateCount = NewCounterValue(ctx, &updateCount, metricName, baseAttribs)

	gaugeValues := []*GaugeValue{}

	metricName = daemonName + "_NFRunning"
	runningGugage, err := meter.Float64ObservableGauge(metricName, api.WithDescription("This value indicates network functions is running or not"))
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create NFRunning")
	}
	NFRunning = NewGaugeValue(&runningGugage, metricName, baseAttribs)
	gaugeValues = append(gaugeValues, NFRunning)

	metricName = daemonName + "_NFStartTime"
	startTimeGuage, err := meter.Float64ObservableGauge(metricName, api.WithDescription("This value indicates start time of the network function since unix epoch in seconds"))
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create NFStartTime")
	}
	NFStartTime = NewGaugeValue(&startTimeGuage, metricName, baseAttribs)
	gaugeValues = append(gaugeValues, NFStartTime)

	metricName = daemonName + "_NFMonitorMap"
	monitorMapGuage, err := meter.Float64ObservableGauge(metricName, api.WithDescription("This value indicates network function monitor counters"))
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create NFMonitorMap")
	}
	NFMonitorMap = NewGaugeValue(&monitorMapGuage, metricName, baseAttribs)
	gaugeValues = append(gaugeValues, NFMonitorMap)

	// Gauge value update is done through callback
	for _, gaugeVal := range gaugeValues {
		gaugeCopy := gaugeVal.Gauge
		_, err = meter.RegisterCallback(func(_ context.Context, o api.Observer) error {
			o.ObserveFloat64(*gaugeCopy, gaugeVal.GetValue(), gaugeVal.GetMeasurementOptions())
			return nil
		}, *gaugeCopy)

		if err != nil {
			log.Warn().Err(err).Msgf("Failed to update metric value for %s", gaugeVal.MetricName)
		}
	}

	// Adding web endpoint
	go func() {
		// Expose the registered metrics via HTTP.
		http.Handle("/metrics", promhttp.Handler())
		if err := http.ListenAndServe(metricsAddr, nil); err != nil {
			log.Fatal().Err(err).Msg("Failed to launch prometheus metrics endpoint")
		}
	}()
}

func Incr(counterVec *CounterValue, ebpfProgram, direction, ifaceName string) {

	if counterVec == nil {
		log.Warn().Msg("Metrics: counter vector is nil and needs to be initialized before Incr")
		return
	}

	updatedAttributes := map[string]string {
		"ebpf_program": ebpfProgram,
		"direction": direction,
		"ifaceName": ifaceName,
	}

	counterVec.SetAttributes(updatedAttributes)
	counterVec.Increment()
}

func Set(value float64, gaugeVec *GaugeValue, ebpfProgram, direction, ifaceName string) {

	if gaugeVec == nil {
		log.Warn().Msg("Metrics: gauge vector is nil and needs to be initialized before Set")
		return
	}

	updatedAttributes := map[string]string {
		"ebpf_program":   ebpfProgram,
		"direction":      direction,
		"interface_name": ifaceName,
	}

	gaugeVec.SetAttributes(updatedAttributes)
	gaugeVec.SetValue(value)
}

func SetValue(value float64, gaugeVec *GaugeValue, ebpfProgram, mapName, ifaceName string) {

	if gaugeVec == nil {
		log.Warn().Msg("Metrics: gauge vector is nil and needs to be initialized before SetValue")
		return
	}

	updatedAttributes := map[string]string {
		"ebpf_program":   ebpfProgram,
		"map_name":       mapName,
		"interface_name": ifaceName,
	}

	gaugeVec.SetAttributes(updatedAttributes)
	gaugeVec.SetValue(value)
}

func SetWithVersion(value float64, gaugeVec *GaugeValue, ebpfProgram, version, direction, ifaceName string) {

	if gaugeVec == nil {
		log.Warn().Msg("Metrics: gauge vector is nil and needs to be initialized before Set")
		return
	}

	updatedAttributes := map[string]string {
		"ebpf_program":   ebpfProgram,
		"version":        version,
		"direction":      direction,
		"interface_name": ifaceName,
	}

	gaugeVec.SetAttributes(updatedAttributes)
	gaugeVec.SetValue(value)
}
