package utils

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"kaznet-status/database"
	"net/http"
)

var (
	IPsStatus = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "ips_status",
			Help: "Info about IP status.",
		},
		[]string{"ip", "lat", "lon", "city"},
	)
)

func init() {
	// Metrics have to be registered to be exposed:
	prometheus.MustRegister(IPsStatus)
}

func InitPrometheus() {
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}

func ChangeIpData(ipData *database.IP) {
	IPsStatus.With(prometheus.Labels{"ip": ipData.IP, "lat": fmt.Sprintf("%f", ipData.Lat), "lon": fmt.Sprintf("%f", ipData.Lon), "city": ipData.City}).Set(float64(ipData.Status))
}
