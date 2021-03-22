package main

import (
	"log"
	"reflect"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
)

type PrometheusTag struct {
	CounterName           string
	CounterType           prometheus.ValueType
	CounterDesc           string
	PrometheusDescription *prometheus.Desc
}

//getPrometheusDesc creates prometheus descritpion for given metrics
func (pt *PrometheusTag) getPrometheusDesc() *prometheus.Desc {

	pt.PrometheusDescription = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, subsystem, pt.CounterName),
		pt.CounterDesc,
		[]string{}, nil,
	)

	return pt.PrometheusDescription

}

type Exporter struct {
	JvbEndpoint string
}

func NewExporter(jvbEndpoint string) *Exporter {
	return &Exporter{
		JvbEndpoint: jvbEndpoint,
	}
}

func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {

	//dynamically create prometheus descriptions
	for _, tag := range promTags {
		ch <- tag.getPrometheusDesc()
	}

	//TODO: edge case with ConferenceSizes
}

func (e *Exporter) Collect(ch chan<- prometheus.Metric) {

	//dynamically create prometheus descriptions
	for _, tag := range promTags {
		log.Println(tag.CounterName)
		ch <- prometheus.MustNewConstMetric(tag.PrometheusDescription, tag.CounterType, 1)
	}

	// ch <- prometheus.MustNewConstMetric(promTags[0].PrometheusDescription, promTags[0].CounterType, 1)

}

//TODO: Create pass by pointer
//listPrometheusTags lists Prometheus data from prom tags within the struct
func listPrometheusTags() []PrometheusTag {
	fields := reflect.TypeOf(&JVBStatistics{}).Elem()
	num := fields.NumField()

	arrPrometheusTags := make([]PrometheusTag, 0, num)

	for i := 0; i < num; i++ {
		field := fields.Field(i)

		tag := parsePromTagFromField(field)
		if tag == nil {
			continue
		}
		tag.CounterName = *parseJSONTagFromField(field)
		tag.getPrometheusDesc()

		arrPrometheusTags = append(arrPrometheusTags, *tag)

	}

	return arrPrometheusTags
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

//parsePromTagFromField Extracts PrometheusTag counter type and desc from
//						struct associated tag
func parsePromTagFromField(rsf reflect.StructField) *PrometheusTag {

	//TODO: Properly handle situation where tag would return empty string
	tag := rsf.Tag.Get("prom")
	prometheusTag := &PrometheusTag{}

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
