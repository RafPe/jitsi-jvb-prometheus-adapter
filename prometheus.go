package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/prometheus/client_golang/prometheus"
)

type PrometheusMetric struct {
	FieldName             string               // Field from struct to lookup for values
	CounterName           string               // json representation of FieldName from struct
	CounterType           prometheus.ValueType // Type of prometheus counter
	CounterDesc           string               // Description of the counter
	PrometheusDescription *prometheus.Desc     // Prometheus description of metric
	CustomCollectorFunc   CustomCollector      // [OPTIONAL] Custom collect function for stats which require additonal logic
}
type Exporter struct {
	jvbEndpoint       string
	prometheusMetrics []PrometheusMetric
}

func NewExporter(jvbEndpoint string) *Exporter {
	exporter := &Exporter{
		jvbEndpoint: jvbEndpoint,
	}

	//dynamically create prometheus descriptions
	exporter.buildPrometheusMetric()

	return exporter
}

func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {

	for _, metric := range e.prometheusMetrics {
		ch <- metric.PrometheusDescription
	}
}

func (e *Exporter) Collect(ch chan<- prometheus.Metric) {

	jvbStats, err := e.getJVBStatistics()
	if err != nil {
		fmt.Println(err)
		return
	}

	if jvbStats == nil {
		fmt.Println("error nil !?")
		fmt.Println(err)
		return
	}

	for _, metric := range e.prometheusMetrics {
		if metric.CustomCollectorFunc != nil {
			resMetric := metric.CustomCollectorFunc(metric, jvbStats)
			if resMetric != nil {
				ch <- *resMetric
			}
			continue
		}

		metricValue := e.getMetricValue(jvbStats, metric.FieldName)

		ch <- prometheus.MustNewConstMetric(metric.PrometheusDescription, metric.CounterType, metricValue)
	}

}

//getJVBStatistics Retrieves statistics from JVB APIs endpoint
func (e *Exporter) getJVBStatistics() (*JVBStatistics, error) {
	// Create a Resty Client
	client := resty.New()

	resp, err := client.R().
		SetResult(JVBStatistics{}).
		EnableTrace().
		Get(os.Getenv(envJVBEndpoint))
	// Explore response object
	fmt.Println("Response Info:")
	fmt.Println("  Error      :", err)
	fmt.Println("  Status Code:", resp.StatusCode())
	fmt.Println("  Status     :", resp.Status())
	fmt.Println("  Proto      :", resp.Proto())
	fmt.Println("  Time       :", resp.Time())
	fmt.Println("  Received At:", resp.ReceivedAt())
	fmt.Println("  Body       :\n", resp)
	fmt.Println()

	// Explore trace info
	fmt.Println("Request Trace Info:")
	ti := resp.Request.TraceInfo()
	fmt.Println("  DNSLookup     :", ti.DNSLookup)
	fmt.Println("  ConnTime      :", ti.ConnTime)
	fmt.Println("  TCPConnTime   :", ti.TCPConnTime)
	fmt.Println("  TLSHandshake  :", ti.TLSHandshake)
	fmt.Println("  ServerTime    :", ti.ServerTime)
	fmt.Println("  ResponseTime  :", ti.ResponseTime)
	fmt.Println("  TotalTime     :", ti.TotalTime)
	fmt.Println("  IsConnReused  :", ti.IsConnReused)
	fmt.Println("  IsConnWasIdle :", ti.IsConnWasIdle)
	fmt.Println("  ConnIdleTime  :", ti.ConnIdleTime)
	fmt.Println("  RequestAttempt:", ti.RequestAttempt)
	fmt.Println("  RemoteAddr    :", ti.RemoteAddr.String())

	if err != nil {
		return nil, err
	}

	//TODO: Handle application errors
	// if resp.IsError() {
	// 	e := resp.Error().(*NetworkListErrorv2)
	// 	if e.Status != 0 {
	// 		return nil, e
	// 	}
	// }

	stats := resp.Result().(*JVBStatistics)

	return stats, nil
}

//getMetricValue Retrieves metric value from stats struct by fieldName
func (e *Exporter) getMetricValue(stats *JVBStatistics, fieldName string) float64 {
	r := reflect.ValueOf(stats)
	fieldVal := reflect.Indirect(r).FieldByName(fieldName)
	promValue := 7.21

	fieldKind := fieldVal.Kind()
	switch fieldKind {
	case reflect.Int:
		promValue = float64(fieldVal.Int())
	case reflect.Float64:
		promValue = float64(fieldVal.Float())
	}

	return promValue
}

//buildPrometheusMetric creates prometheus descritpion for given metrics
func (e *Exporter) buildPrometheusMetric() {

	fields := reflect.TypeOf(&JVBStatistics{}).Elem()
	num := fields.NumField()

	for i := 0; i < num; i++ {
		field := fields.Field(i)

		tag := parsePromTagFromField(field)
		if tag == nil {
			continue
		}
		tag.CounterName = *parseJSONTagFromField(field)
		tag.FieldName = field.Name

		tag.PrometheusDescription = prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, tag.CounterName),
			tag.CounterDesc,
			[]string{}, nil,
		)

		//TODO: edge case scenarios - use more streamlined method of associating func
		if tag.CounterName == "conference_sizes" {
			tag.CustomCollectorFunc = ConfSizesByParticipants
		}

		//TODO: edge case scenarios - use more streamlined method of associating func
		if tag.CounterName == "conferences_by_audio_senders" {
			tag.CustomCollectorFunc = ConfSizesByAudioSenders
		}

		//TODO: edge case scenarios - use more streamlined method of associating func
		if tag.CounterName == "conferences_by_video_senders" {
			tag.CustomCollectorFunc = ConfSizesByVideoSenders
		}

		e.prometheusMetrics = append(e.prometheusMetrics, *tag)

	}
}

//parseJSONTagFromField Extracts JSON field name from the struct
func parseJSONTagFromField(rsf reflect.StructField) *string {
	jsonTag := rsf.Tag.Get("json")

	if jsonTag != "" {
		JSONTagValues := strings.SplitN(jsonTag, ",", 1)

		return &JSONTagValues[0]
	}

	return nil
}

//parsePromTagFromField Extracts PrometheusTag counter type and desc from struct associated tag
func parsePromTagFromField(rsf reflect.StructField) *PrometheusMetric {

	tag := rsf.Tag.Get("prom")
	prometheusTag := &PrometheusMetric{
		CustomCollectorFunc: nil,
	}

	if tag != "" {

		//TODO: Handle case where struct would not contain properly formatted tags
		//		Since this is v0 - we will fix in prod release :)
		promTagValues := strings.SplitN(tag, ";", 2)
		prometheusTag.CounterDesc = promTagValues[1]

		//TODO: We can move this later on to more dynamic value mapping
		//		or just simplify
		switch promTagValues[0] {
		case "Gauge":
			prometheusTag.CounterType = prometheus.GaugeValue
		case "Counter":
			prometheusTag.CounterType = prometheus.CounterValue
		case "Untyped":
			prometheusTag.CounterType = prometheus.UntypedValue
		}

		return prometheusTag
	}

	return nil
}

//conferenceSizesHistogram Creates histogram for conference sizes
func conferenceSizesHistogram(conferencesData []int) (conferenceSizesHistogram map[float64]uint64, sum uint64) {
	var sizes = make(map[float64]uint64)

	//calculate sum for histogram
	sum = 0
	for _, v := range conferencesData {
		sum += uint64(v)
	}

	//for the histgram buckets we need to omit the last field b/c the +inf bucket is added automatically
	conferencesData = conferencesData[:len(conferencesData)-1]

	//the bucket values have to be cumulative
	var i int
	for i = len(conferencesData) - 1; i >= 0; i-- {
		var cumulative int
		var j int
		for j = i; j >= 0; j-- {
			cumulative += conferencesData[j]
		}
		conferencesData[i] = cumulative
	}

	for i, v := range conferencesData {
		sizes[float64(i)] = uint64(v)
	}

	return sizes, sum
}
