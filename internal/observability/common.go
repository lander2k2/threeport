package observability

import (
	"fmt"
)

const (
	GrafanaHelmRepo       = "https://grafana-community.github.io/helm-charts"
	PrometheusHelmRepo    = "https://prometheus-community.github.io/helm-charts"
	MimirHelmRepo         = "https://grafana.github.io/helm-charts"
	LokiHelmRepo          = "https://grafana-community.github.io/helm-charts"
	TempoHelmRepo         = "https://grafana-community.github.io/helm-charts"
	OtelCollectorHelmRepo = "https://open-telemetry.github.io/opentelemetry-helm-charts"
)

// ObservabilityDashboardName returns the name of an observability dashboard object
func ObservabilityDashboardName(name string) string {
	return fmt.Sprintf("%s-observability-dashboard", name)
}

// MetricsName returns the name of a metrics object
func MetricsName(name string) string {
	return fmt.Sprintf("%s-metrics", name)
}

// MetricsStorageName returns the name of a metrics storage object
func MetricsStorageName(name string) string {
	return fmt.Sprintf("%s-metrics-storage", name)
}

// LoggingName returns the name of a logging chart
func LoggingName(name string) string {
	return fmt.Sprintf("%s-logging", name)
}

// TracingName returns the name of a tracing object
func TracingName(name string) string {
	return fmt.Sprintf("%s-tracing", name)
}

// InstrumentationAgentName returns the name of an instrumentation agent object
func InstrumentationAgentName(name string) string {
	return fmt.Sprintf("%s-instrumentation-agent", name)
}

// InstrumentationGatewayName returns the name of an instrumentation gateway object
func InstrumentationGatewayName(name string) string {
	return fmt.Sprintf("%s-instrumentation-gateway", name)
}

// InstrumentationBrowserRelayName returns the name of an instrumentation browser
// relay object
func InstrumentationBrowserRelayName(name string) string {
	return fmt.Sprintf("%s-instrumentation-browser-relay", name)
}

// GrafanaChartName returns the name of the grafana chart
func GrafanaChartName(name string) string {
	return fmt.Sprintf("%s-grafana", name)
}

// KubePrometheusStackChartName returns the name of the kube-prometheus-stack
// chart
func KubePrometheusStackChartName(name string) string {
	return fmt.Sprintf("%s-kube-prometheus-stack", name)
}

// MimirHelmChartName returns the name of the mimir chart
func MimirHelmChartName(name string) string {
	return fmt.Sprintf("%s-mimir", name)
}

// LokiHelmChartName returns the name of the loki chart
func LokiHelmChartName(name string) string {
	return fmt.Sprintf("%s-loki", name)
}

// TempoHelmChartName returns the name of the tempo chart
func TempoHelmChartName(name string) string {
	return fmt.Sprintf("%s-tempo", name)
}

// OtelAgentHelmChartName returns the name of the OTel Collector Agent chart
func OtelAgentHelmChartName(name string) string {
	return fmt.Sprintf("%s-otel-agent", name)
}

// OtelGatewayHelmChartName returns the name of the OTel Collector Gateway chart
func OtelGatewayHelmChartName(name string) string {
	return fmt.Sprintf("%s-otel-gateway", name)
}

// OtelBrowserRelayHelmChartName returns the name of the OTel Collector Browser
// Relay chart
func OtelBrowserRelayHelmChartName(name string) string {
	return fmt.Sprintf("%s-otel-browser-relay", name)
}
