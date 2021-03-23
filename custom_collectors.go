package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

type CustomCollector func(metric PrometheusMetric, stats *JVBStatistics) *prometheus.Metric

func ConfSizesByParticipants(metric PrometheusMetric, stats *JVBStatistics) *prometheus.Metric {

	conSizes, sum := conferenceSizesHistogram(stats.ConferenceSizes)

	promMetric, err := prometheus.NewConstHistogram(metric.PrometheusDescription, sum, float64(sum), conSizes)
	if err != nil {
		return nil
	}

	return &promMetric
}

func ConfSizesByAudioSenders(metric PrometheusMetric, stats *JVBStatistics) *prometheus.Metric {

	conSizes, sum := conferenceSizesHistogram(stats.ConferencesByAudioSenders)

	promMetric, err := prometheus.NewConstHistogram(metric.PrometheusDescription, sum, float64(sum), conSizes)
	if err != nil {
		return nil
	}

	return &promMetric
}

func ConfSizesByVideoSenders(metric PrometheusMetric, stats *JVBStatistics) *prometheus.Metric {

	conSizes, sum := conferenceSizesHistogram(stats.ConferencesByVideoSenders)

	promMetric, err := prometheus.NewConstHistogram(metric.PrometheusDescription, sum, float64(sum), conSizes)
	if err != nil {
		return nil
	}

	return &promMetric
}
