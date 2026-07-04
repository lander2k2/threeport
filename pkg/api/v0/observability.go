package v0

// ObservabilityStackDefinition defines an observability stack.
type ObservabilityStackDefinition struct {
	Common         `swaggerignore:"true" mapstructure:",squash"`
	Definition     `mapstructure:",squash"`
	Reconciliation `mapstructure:",squash"`

	// Dashboard
	// The observability dashboard definition that belongs to this resource.
	ObservabilityDashboardDefinitionID *uint `json:",omitempty" validate:"optional" relationship:"owns"`

	// The version of the grafana helm chart to use from the helm repo, e.g. 1.2.3
	GrafanaHelmChartVersion *string `json:",omitempty" validate:"optional"`

	// Optional Helm workload definition values that can be provided to configure the
	// underlying grafana chart.
	GrafanaHelmValuesDocument *string `json:",omitempty" validate:"optional"`

	// Metrics
	// The metrics definition that belongs to this resource.
	MetricsDefinitionID *uint `json:",omitempty" validate:"optional" relationship:"owns"`

	// The version of the kube-prometheus-stack helm chart to use from the helm repo, e.g. 1.2.3
	KubePrometheusStackHelmChartVersion *string `json:",omitempty" validate:"optional"`

	// Optional Helm workload definition values that can be provided to configure the
	// underlying kube-prometheus-stack chart.
	KubePrometheusStackHelmValuesDocument *string `json:",omitempty" validate:"optional"`

	// Logging
	// The logging definition that belongs to this resource.
	LoggingDefinitionID *uint `json:",omitempty" validate:"optional" relationship:"owns"`

	// The version of the loki helm chart to use from the helm repo, e.g. 1.2.3
	LokiHelmChartVersion *string `json:",omitempty" validate:"optional"`

	// Optional Helm workload definition values that can be provided to configure the
	// underlying loki chart.
	LokiHelmValuesDocument *string `json:",omitempty" validate:"optional"`

	// Tracing
	// The tracing definition that belongs to this resource.
	TracingDefinitionID *uint `json:",omitempty" validate:"optional" relationship:"owns"`

	// The version of the tempo helm chart to use from the helm repo, e.g. 1.2.3
	TempoHelmChartVersion *string `json:",omitempty" validate:"optional"`

	// Optional Helm workload definition values that can be provided to configure the
	// underlying tempo chart.
	TempoHelmValuesDocument *string `json:",omitempty" validate:"optional"`

	// Instrumentation (OTel Collector pipeline)
	// The instrumentation agent definition that belongs to this resource.
	InstrumentationAgentDefinitionID *uint `json:",omitempty" validate:"optional" relationship:"owns"`

	// The instrumentation gateway definition that belongs to this resource.
	InstrumentationGatewayDefinitionID *uint `json:",omitempty" validate:"optional" relationship:"owns"`

	// The instrumentation browser relay definition that belongs to this resource.
	InstrumentationBrowserRelayDefinitionID *uint `json:",omitempty" validate:"optional" relationship:"owns"`

	// The version of the opentelemetry-collector helm chart to use from the helm
	// repo, e.g. 1.2.3. Shared by the OTel Agent, Gateway and Browser Relay releases.
	OtelCollectorHelmChartVersion *string `json:",omitempty" validate:"optional"`

	// Optional Helm workload definition values that can be provided to configure the
	// OTel Collector Agent (DaemonSet).
	OtelAgentHelmValuesDocument *string `json:",omitempty" validate:"optional"`

	// Optional Helm workload definition values that can be provided to configure the
	// OTel Collector Gateway (Deployment).
	OtelGatewayHelmValuesDocument *string `json:",omitempty" validate:"optional"`

	// Optional Helm workload definition values that can be provided to configure the
	// OTel Collector Browser Relay (Deployment).
	OtelBrowserRelayHelmValuesDocument *string `json:",omitempty" validate:"optional"`

	// The version of the mimir-distributed helm chart to use from the helm repo, e.g. 1.2.3
	MimirHelmChartVersion *string `json:",omitempty" validate:"optional"`

	// Optional Helm workload definition values that can be provided to configure the
	// underlying mimir chart.
	MimirHelmValuesDocument *string `json:",omitempty" validate:"optional"`

	// The associated observability stack instances that are deployed from this definition.
	ObservabilityStackInstances []*ObservabilityStackInstance `json:",omitempty" validate:"optional,association"`
}

// ObservabilityStackInstance is a deployed instance of an observability stack.
type ObservabilityStackInstance struct {
	Common         `swaggerignore:"true" mapstructure:",squash"`
	Instance       `mapstructure:",squash"`
	Reconciliation `mapstructure:",squash"`

	// The observability stack definition that belongs to this resource.
	ObservabilityStackDefinitionID *uint `json:",omitempty" validate:"required" gorm:"not null" relationship:"requires"`

	// The kubernetes runtime where the observability stack is installed.
	KubernetesRuntimeInstanceID *uint `json:",omitempty" validate:"required" gorm:"not null" relationship:"requires"`

	// If true, metrics will be enabled for the observability stack.
	MetricsEnabled *bool `json:",omitempty" validate:"optional" gorm:"default:true"`

	// If true, logging will be enabled for the observability stack.
	LoggingEnabled *bool `json:",omitempty" validate:"optional" gorm:"default:true"`

	// If true, tracing will be enabled for the observability stack.
	TracingEnabled *bool `json:",omitempty" validate:"optional" gorm:"default:true"`

	// If true, the OTel Collector Agent (DaemonSet) will be enabled for the
	// observability stack. Required for logging and tracing signals.
	InstrumentationAgentEnabled *bool `json:",omitempty" validate:"optional" gorm:"default:true"`

	// If true, the OTel Collector Gateway (Deployment) will be enabled for the
	// observability stack. Required when tracing is in use with tail_sampling.
	// This Gateway can be disabled if when tracking is in use with head_sampling.
	InstrumentationGatewayEnabled *bool `json:",omitempty" validate:"optional" gorm:"default:true"`

	// If true, the OTel Collector Browser Relay (Deployment) will be enabled for the
	// observability stack. Only needed when a single-page app or similar browser
	// client is exporting telemetry.
	InstrumentationBrowserRelayEnabled *bool `json:",omitempty" validate:"optional" gorm:"default:false"`

	// Dashboard
	// The observability dashboard instance that belongs to this resource.
	ObservabilityDashboardInstanceID *uint `json:",omitempty" validate:"optional" relationship:"owns"`

	// Optional Helm workload instance values that can be provided to configure the
	// underlying grafana chart.
	GrafanaHelmValuesDocument *string `json:",omitempty" validate:"optional"`

	// Metrics
	// The metrics instance that belongs to this resource.
	MetricsInstanceID *uint `json:",omitempty" validate:"optional" relationship:"owns"`

	// Optional Helm workload instance values that can be provided to configure the
	// underlying kube-prometheus-stack chart.
	KubePrometheusStackHelmValuesDocument *string `json:",omitempty" validate:"optional"`

	// Logging
	// The logging instance that belongs to this resource.
	LoggingInstanceID *uint `json:",omitempty" validate:"optional" relationship:"owns"`

	// Optional Helm workload instance values that can be provided to configure the
	// underlying loki chart.
	LokiHelmValuesDocument *string `json:",omitempty" validate:"optional"`

	// Tracing
	// The tracing instance that belongs to this resource.
	TracingInstanceID *uint `json:",omitempty" validate:"optional" relationship:"owns"`

	// Optional Helm workload instance values that can be provided to configure the
	// underlying tempo chart.
	TempoHelmValuesDocument *string `json:",omitempty" validate:"optional"`

	// Instrumentation (OTel Collector pipeline)
	// The instrumentation agent instance that belongs to this resource.
	InstrumentationAgentInstanceID *uint `json:",omitempty" validate:"optional" relationship:"owns"`

	// The instrumentation gateway instance that belongs to this resource.
	InstrumentationGatewayInstanceID *uint `json:",omitempty" validate:"optional" relationship:"owns"`

	// The instrumentation browser relay instance that belongs to this resource.
	InstrumentationBrowserRelayInstanceID *uint `json:",omitempty" validate:"optional" relationship:"owns"`

	// Optional Helm workload instance values that can be provided to configure the
	// OTel Collector Agent (DaemonSet).
	OtelAgentHelmValuesDocument *string `json:",omitempty" validate:"optional"`

	// Optional Helm workload instance values that can be provided to configure the
	// OTel Collector Gateway (Deployment).
	OtelGatewayHelmValuesDocument *string `json:",omitempty" validate:"optional"`

	// Optional Helm workload instance values that can be provided to configure the
	// OTel Collector Browser Relay (Deployment).
	OtelBrowserRelayHelmValuesDocument *string `json:",omitempty" validate:"optional"`

	// Optional Helm workload instance values that can be provided to configure the
	// underlying mimir chart.
	MimirHelmValuesDocument *string `json:",omitempty" validate:"optional"`
}

// ObservabilityDashboardDefinition is the definition of an observability dashboard.
type ObservabilityDashboardDefinition struct {
	Common         `swaggerignore:"true" mapstructure:",squash"`
	Definition     `mapstructure:",squash"`
	Reconciliation `mapstructure:",squash"`

	// The Grafana Helm workload definition that belongs to this resource.
	GrafanaHelmWorkloadDefinitionID *uint `json:",omitempty" validate:"optional" relationship:"owns;type:HelmWorkloadDefinition"`

	// The version of the grafana helm chart to use from the helm repo, e.g. 1.2.3
	GrafanaHelmChartVersion *string `json:",omitempty" validate:"optional" gorm:"default:'12.7.2'"`

	// Optional Helm workload definition values that can be provided to configure the
	// underlying grafana chart.
	GrafanaHelmValuesDocument *string `json:",omitempty" validate:"optional"`

	// The associated observability dashboard instances that are deployed from this definition.
	ObservabilityDashboardInstances []*ObservabilityDashboardInstance `json:",omitempty" validate:"optional,association"`
}

// ObservabilityDashboardInstances is a deployed instance of an observability dashboard.
type ObservabilityDashboardInstance struct {
	Common         `swaggerignore:"true" mapstructure:",squash"`
	Instance       `mapstructure:",squash"`
	Reconciliation `mapstructure:",squash"`

	// The observability dashboard definition that belongs to this resource.
	ObservabilityDashboardDefinitionID *uint `json:",omitempty" validate:"required" gorm:"not null" relationship:"requires"`

	// The kubernetes runtime where the observability dashboard is installed.
	KubernetesRuntimeInstanceID *uint `json:",omitempty" validate:"required" gorm:"not null" relationship:"requires"`

	// The Grafana Helm workload instance that belongs to this resource.
	GrafanaHelmWorkloadInstanceID *uint `json:",omitempty" validate:"optional" relationship:"owns;type:HelmWorkloadInstance"`

	// Optional Helm workload definition values that can be provided to configure the
	// underlying grafana chart.
	GrafanaHelmValuesDocument *string `json:",omitempty" validate:"optional"`
}

// MetricsDefinition is the definition of a metrics aggregation layer for a workload.
type MetricsDefinition struct {
	Common         `swaggerignore:"true" mapstructure:",squash"`
	Definition     `mapstructure:",squash"`
	Reconciliation `mapstructure:",squash"`

	// The kube-prometheus-stack Helm workload definition that belongs to this resource.
	KubePrometheusStackHelmWorkloadDefinitionID *uint `json:",omitempty" validate:"optional" relationship:"owns;type:HelmWorkloadDefinition"`

	// The version of the kube-prometheus-stack helm chart to use from the helm repo, e.g. 1.2.3
	KubePrometheusStackHelmChartVersion *string `json:",omitempty" validate:"optional" gorm:"default:'87.10.0'"`

	// Optional Helm workload definition values that can be provided to configure the
	// underlying kube-prometheus-stack chart.
	KubePrometheusStackHelmValuesDocument *string `json:",omitempty" validate:"optional"`

	// The mimir-distributed Helm workload definition that belongs to this resource.
	MimirHelmWorkloadDefinitionID *uint `json:",omitempty" validate:"optional" relationship:"owns;type:HelmWorkloadDefinition"`

	// The version of the mimir-distributed helm chart to use from the helm repo, e.g. 1.2.3
	MimirHelmChartVersion *string `json:",omitempty" validate:"optional" gorm:"default:'6.1.0'"`

	// Optional Helm workload definition values that can be provided to configure the
	// underlying mimir chart.
	MimirHelmValuesDocument *string `json:",omitempty" validate:"optional"`

	// The associated metrics instances that are deployed from this definition.
	MetricsInstances []*MetricsInstance `json:",omitempty" validate:"optional,association"`
}

// MetricsInstances is a deployed instance of a metrics aggregation layer for a workload.
type MetricsInstance struct {
	Common         `swaggerignore:"true" mapstructure:",squash"`
	Instance       `mapstructure:",squash"`
	Reconciliation `mapstructure:",squash"`

	// The metrics definition that belongs to this resource.
	MetricsDefinitionID *uint `json:",omitempty" validate:"required" gorm:"not null" relationship:"requires"`

	// The kubernetes runtime where the metrics is installed.
	KubernetesRuntimeInstanceID *uint `json:",omitempty" validate:"required" gorm:"not null" relationship:"requires"`

	// The kube-prometheus-stack helm workload instance that belongs to this resource.
	KubePrometheusStackHelmWorkloadInstanceID *uint `json:",omitempty" validate:"optional" relationship:"owns;type:HelmWorkloadInstance"`

	// Optional Helm workload instance values that can be provided to configure the
	// underlying kube-prometheus-stack chart.
	KubePrometheusStackHelmValuesDocument *string `json:",omitempty" validate:"optional"`

	// The mimir-distributed helm workload instance that belongs to this resource.
	MimirHelmWorkloadInstanceID *uint `json:",omitempty" validate:"optional" relationship:"owns;type:HelmWorkloadInstance"`

	// Optional Helm workload instance values that can be provided to configure the
	// underlying mimir chart.
	MimirHelmValuesDocument *string `json:",omitempty" validate:"optional"`
}

// LoggingDefinition is the definition of a logging implementation for a workload.
type LoggingDefinition struct {
	Common         `swaggerignore:"true" mapstructure:",squash"`
	Definition     `mapstructure:",squash"`
	Reconciliation `mapstructure:",squash"`

	// The loki Helm workload definition that belongs to this resource.
	LokiHelmWorkloadDefinitionID *uint `json:",omitempty" validate:"optional" relationship:"owns;type:HelmWorkloadDefinition"`

	// The version of the loki helm chart to use from the helm repo, e.g. 1.2.3
	LokiHelmChartVersion *string `json:",omitempty" validate:"optional" gorm:"default:'7.1.0'"`

	// Optional Helm workload definition values that can be provided to configure the
	// underlying loki chart.
	LokiHelmValuesDocument *string `json:",omitempty" validate:"optional"`

	// The associated logging instances that are deployed from this definition.
	LoggingInstances []*LoggingInstance `json:",omitempty" validate:"optional,association"`
}

// LoggingInstances is a deployed instance of a logging implementation for a workload.
type LoggingInstance struct {
	Common         `swaggerignore:"true" mapstructure:",squash"`
	Instance       `mapstructure:",squash"`
	Reconciliation `mapstructure:",squash"`

	// The logging definition that belongs to this resource.
	LoggingDefinitionID *uint `json:",omitempty" validate:"required" gorm:"not null" relationship:"requires"`

	// The kubernetes runtime where the logging is installed.
	KubernetesRuntimeInstanceID *uint `json:",omitempty" validate:"required" gorm:"not null" relationship:"requires"`

	// The loki helm workload instance that belongs to this resource.
	LokiHelmWorkloadInstanceID *uint `json:",omitempty" validate:"optional" relationship:"owns;type:HelmWorkloadInstance"`

	// Optional Helm workload instance values that can be provided to configure the
	// underlying loki chart.
	LokiHelmValuesDocument *string `json:",omitempty" validate:"optional"`
}

// TracingDefinition is the definition of a distributed tracing layer for a workload.
type TracingDefinition struct {
	Common         `swaggerignore:"true" mapstructure:",squash"`
	Definition     `mapstructure:",squash"`
	Reconciliation `mapstructure:",squash"`

	// The tempo Helm workload definition that belongs to this resource.
	TempoHelmWorkloadDefinitionID *uint `json:",omitempty" validate:"optional" relationship:"owns;type:HelmWorkloadDefinition"`

	// The version of the tempo helm chart to use from the helm repo, e.g. 1.2.3
	TempoHelmChartVersion *string `json:",omitempty" validate:"optional" gorm:"default:'2.2.3'"`

	// Optional Helm workload definition values that can be provided to configure the
	// underlying tempo chart.
	TempoHelmValuesDocument *string `json:",omitempty" validate:"optional"`

	// The associated tracing instances that are deployed from this definition.
	TracingInstances []*TracingInstance `json:",omitempty" validate:"optional,association"`
}

// TracingInstance is a deployed instance of a distributed tracing layer for a workload.
type TracingInstance struct {
	Common         `swaggerignore:"true" mapstructure:",squash"`
	Instance       `mapstructure:",squash"`
	Reconciliation `mapstructure:",squash"`

	// The tracing definition that belongs to this resource.
	TracingDefinitionID *uint `json:",omitempty" validate:"required" gorm:"not null" relationship:"requires"`

	// The kubernetes runtime where the tracing is installed.
	KubernetesRuntimeInstanceID *uint `json:",omitempty" validate:"required" gorm:"not null" relationship:"requires"`

	// The tempo helm workload instance that belongs to this resource.
	TempoHelmWorkloadInstanceID *uint `json:",omitempty" validate:"optional" relationship:"owns;type:HelmWorkloadInstance"`

	// Optional Helm workload instance values that can be provided to configure the
	// underlying tempo chart.
	TempoHelmValuesDocument *string `json:",omitempty" validate:"optional"`
}

// InstrumentationAgentDefinition is the definition of the OTel Collector Agent
// (DaemonSet) for a workload.
type InstrumentationAgentDefinition struct {
	Common         `swaggerignore:"true" mapstructure:",squash"`
	Definition     `mapstructure:",squash"`
	Reconciliation `mapstructure:",squash"`

	// The version of the opentelemetry-collector helm chart to use from the helm
	// repo, e.g. 1.2.3.
	OtelCollectorHelmChartVersion *string `json:",omitempty" validate:"optional" gorm:"default:'0.162.0'"`

	// The OTel Collector Agent (DaemonSet) Helm workload definition that belongs to
	// this resource.
	OtelAgentHelmWorkloadDefinitionID *uint `json:",omitempty" validate:"optional" relationship:"owns;type:HelmWorkloadDefinition"`

	// Optional Helm workload definition values that can be provided to configure the
	// OTel Collector Agent (DaemonSet).
	OtelAgentHelmValuesDocument *string `json:",omitempty" validate:"optional"`

	// The associated instrumentation agent instances that are deployed from this definition.
	InstrumentationAgentInstances []*InstrumentationAgentInstance `json:",omitempty" validate:"optional,association"`
}

// InstrumentationAgentInstance is a deployed instance of the OTel Collector Agent
// (DaemonSet) for a workload.
type InstrumentationAgentInstance struct {
	Common         `swaggerignore:"true" mapstructure:",squash"`
	Instance       `mapstructure:",squash"`
	Reconciliation `mapstructure:",squash"`

	// The instrumentation agent definition that belongs to this resource.
	InstrumentationAgentDefinitionID *uint `json:",omitempty" validate:"required" gorm:"not null" relationship:"requires"`

	// The kubernetes runtime where the instrumentation agent is installed.
	KubernetesRuntimeInstanceID *uint `json:",omitempty" validate:"required" gorm:"not null" relationship:"requires"`

	// The OTel Collector Agent (DaemonSet) helm workload instance that belongs to
	// this resource.
	OtelAgentHelmWorkloadInstanceID *uint `json:",omitempty" validate:"optional" relationship:"owns;type:HelmWorkloadInstance"`

	// Optional Helm workload instance values that can be provided to configure the
	// OTel Collector Agent (DaemonSet).
	OtelAgentHelmValuesDocument *string `json:",omitempty" validate:"optional"`
}

// InstrumentationGatewayDefinition is the definition of the OTel Collector Gateway
// (Deployment) for a workload.
type InstrumentationGatewayDefinition struct {
	Common         `swaggerignore:"true" mapstructure:",squash"`
	Definition     `mapstructure:",squash"`
	Reconciliation `mapstructure:",squash"`

	// The version of the opentelemetry-collector helm chart to use from the helm
	// repo, e.g. 1.2.3.
	OtelCollectorHelmChartVersion *string `json:",omitempty" validate:"optional" gorm:"default:'0.162.0'"`

	// The OTel Collector Gateway (Deployment) Helm workload definition that belongs
	// to this resource.
	OtelGatewayHelmWorkloadDefinitionID *uint `json:",omitempty" validate:"optional" relationship:"owns;type:HelmWorkloadDefinition"`

	// Optional Helm workload definition values that can be provided to configure the
	// OTel Collector Gateway (Deployment).
	OtelGatewayHelmValuesDocument *string `json:",omitempty" validate:"optional"`

	// The associated instrumentation gateway instances that are deployed from this definition.
	InstrumentationGatewayInstances []*InstrumentationGatewayInstance `json:",omitempty" validate:"optional,association"`
}

// InstrumentationGatewayInstance is a deployed instance of the OTel Collector
// Gateway (Deployment) for a workload.
type InstrumentationGatewayInstance struct {
	Common         `swaggerignore:"true" mapstructure:",squash"`
	Instance       `mapstructure:",squash"`
	Reconciliation `mapstructure:",squash"`

	// The instrumentation gateway definition that belongs to this resource.
	InstrumentationGatewayDefinitionID *uint `json:",omitempty" validate:"required" gorm:"not null" relationship:"requires"`

	// The kubernetes runtime where the instrumentation gateway is installed.
	KubernetesRuntimeInstanceID *uint `json:",omitempty" validate:"required" gorm:"not null" relationship:"requires"`

	// The OTel Collector Gateway (Deployment) helm workload instance that belongs to
	// this resource.
	OtelGatewayHelmWorkloadInstanceID *uint `json:",omitempty" validate:"optional" relationship:"owns;type:HelmWorkloadInstance"`

	// Optional Helm workload instance values that can be provided to configure the
	// OTel Collector Gateway (Deployment).
	OtelGatewayHelmValuesDocument *string `json:",omitempty" validate:"optional"`
}

// InstrumentationBrowserRelayDefinition is the definition of the OTel Collector
// Browser Relay (Deployment) for a workload.
type InstrumentationBrowserRelayDefinition struct {
	Common         `swaggerignore:"true" mapstructure:",squash"`
	Definition     `mapstructure:",squash"`
	Reconciliation `mapstructure:",squash"`

	// The version of the opentelemetry-collector helm chart to use from the helm
	// repo, e.g. 1.2.3.
	OtelCollectorHelmChartVersion *string `json:",omitempty" validate:"optional" gorm:"default:'0.162.0'"`

	// The OTel Collector Browser Relay (Deployment) Helm workload definition that
	// belongs to this resource.
	OtelBrowserRelayHelmWorkloadDefinitionID *uint `json:",omitempty" validate:"optional" relationship:"owns;type:HelmWorkloadDefinition"`

	// Optional Helm workload definition values that can be provided to configure the
	// OTel Collector Browser Relay (Deployment).
	OtelBrowserRelayHelmValuesDocument *string `json:",omitempty" validate:"optional"`

	// The associated instrumentation browser relay instances that are deployed from this definition.
	InstrumentationBrowserRelayInstances []*InstrumentationBrowserRelayInstance `json:",omitempty" validate:"optional,association"`
}

// InstrumentationBrowserRelayInstance is a deployed instance of the OTel Collector
// Browser Relay (Deployment) for a workload.
type InstrumentationBrowserRelayInstance struct {
	Common         `swaggerignore:"true" mapstructure:",squash"`
	Instance       `mapstructure:",squash"`
	Reconciliation `mapstructure:",squash"`

	// The instrumentation browser relay definition that belongs to this resource.
	InstrumentationBrowserRelayDefinitionID *uint `json:",omitempty" validate:"required" gorm:"not null" relationship:"requires"`

	// The kubernetes runtime where the instrumentation browser relay is installed.
	KubernetesRuntimeInstanceID *uint `json:",omitempty" validate:"required" gorm:"not null" relationship:"requires"`

	// The OTel Collector Browser Relay (Deployment) helm workload instance that
	// belongs to this resource.
	OtelBrowserRelayHelmWorkloadInstanceID *uint `json:",omitempty" validate:"optional" relationship:"owns;type:HelmWorkloadInstance"`

	// Optional Helm workload instance values that can be provided to configure the
	// OTel Collector Browser Relay (Deployment).
	OtelBrowserRelayHelmValuesDocument *string `json:",omitempty" validate:"optional"`
}
