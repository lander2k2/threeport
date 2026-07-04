package observability

import (
	"errors"
	"fmt"

	"github.com/go-logr/logr"
	v0 "github.com/threeport/threeport/pkg/api/v0"
	client_lib "github.com/threeport/threeport/pkg/client/lib/v0"
	client "github.com/threeport/threeport/pkg/client/v0"
	controller "github.com/threeport/threeport/pkg/controller/v0"
	util "github.com/threeport/threeport/pkg/util/v0"
)

// ObservabilityStackDefinitionConfig contains the configuration for an observability dashboard
// reconcile function.
type ObservabilityStackDefinitionConfig struct {
	r                            *controller.Reconciler
	observabilityStackDefinition *v0.ObservabilityStackDefinition
	log                          *logr.Logger
}

// getObservabilityStackDefinitionOperations returns the operations for an observability
// stack definition
func (c *ObservabilityStackDefinitionConfig) getObservabilityStackDefinitionOperations() *util.Operations {
	operations := util.Operations{}

	// append observability dashboard definition operations
	operations.AppendOperation(util.Operation{
		Name:   "observability dashboard",
		Create: c.createObservabilityDashboardDefinition,
		Delete: c.deleteObservabilityDashboardDefinition,
	})

	// append logging definition operations
	operations.AppendOperation(util.Operation{
		Name:   "logging",
		Create: c.createLoggingDefinition,
		Delete: c.deleteLoggingDefinition,
	})

	// append metrics definition operations
	operations.AppendOperation(util.Operation{
		Name:   "metrics",
		Create: c.createMetricsDefinition,
		Delete: c.deleteMetricsDefinition,
	})

	// append metrics storage definition operations
	operations.AppendOperation(util.Operation{
		Name:   "metrics storage",
		Create: c.createMetricsStorageDefinition,
		Delete: c.deleteMetricsStorageDefinition,
	})

	// append tracing definition operations
	operations.AppendOperation(util.Operation{
		Name:   "tracing",
		Create: c.createTracingDefinition,
		Delete: c.deleteTracingDefinition,
	})

	// append instrumentation agent definition operations
	operations.AppendOperation(util.Operation{
		Name:   "instrumentation agent",
		Create: c.createInstrumentationAgentDefinition,
		Delete: c.deleteInstrumentationAgentDefinition,
	})

	// append instrumentation gateway definition operations
	operations.AppendOperation(util.Operation{
		Name:   "instrumentation gateway",
		Create: c.createInstrumentationGatewayDefinition,
		Delete: c.deleteInstrumentationGatewayDefinition,
	})

	// append instrumentation browser relay definition operations
	operations.AppendOperation(util.Operation{
		Name:   "instrumentation browser relay",
		Create: c.createInstrumentationBrowserRelayDefinition,
		Delete: c.deleteInstrumentationBrowserRelayDefinition,
	})

	return &operations
}

// createObservabilityDashboardDefinition creates an observability dashboard definition.
func (c *ObservabilityStackDefinitionConfig) createObservabilityDashboardDefinition() error {
	// create observability dashboard definition
	observabilityDashboardDefinition := &v0.ObservabilityDashboardDefinition{
		Definition: v0.Definition{
			Name: util.Ptr(ObservabilityDashboardName(*c.observabilityStackDefinition.Name)),
		},
	}

	// set grafana helm chart version
	observabilityDashboardDefinition.GrafanaHelmChartVersion = c.observabilityStackDefinition.GrafanaHelmChartVersion

	// set grafana helm chart values
	observabilityDashboardDefinition.GrafanaHelmValuesDocument = c.observabilityStackDefinition.GrafanaHelmValuesDocument

	// create observability dashboard definition
	createdObservabilityDashboardDefinition, err := client.CreateObservabilityDashboardDefinition(
		c.r.APIClient,
		c.r.APIServer,
		observabilityDashboardDefinition,
	)
	if err != nil {
		return fmt.Errorf("failed to create observability dashboard definition: %w", err)
	}

	// update observability stack definition with observability dashboard definition id
	c.observabilityStackDefinition.ObservabilityDashboardDefinitionID = createdObservabilityDashboardDefinition.ID

	return nil
}

// deleteObservabilityDashboardDefinition deletes an observability dashboard definition.
func (c *ObservabilityStackDefinitionConfig) deleteObservabilityDashboardDefinition() error {
	// delete observability dashboard definition
	if _, err := client.DeleteObservabilityDashboardDefinition(
		c.r.APIClient,
		c.r.APIServer,
		*c.observabilityStackDefinition.ObservabilityDashboardDefinitionID,
	); err != nil && !errors.Is(err, client_lib.ErrObjectNotFound) {
		return fmt.Errorf("failed to delete observability dashboard definition: %w", err)
	}

	return nil
}

// createLoggingDefinition creates a logging definition.
func (c *ObservabilityStackDefinitionConfig) createLoggingDefinition() error {
	// create logging definition
	loggingDefinition := &v0.LoggingDefinition{
		Definition: v0.Definition{
			Name: util.Ptr(LoggingName(*c.observabilityStackDefinition.Name)),
		},
	}

	// set loki helm chart version
	loggingDefinition.LokiHelmChartVersion = c.observabilityStackDefinition.LokiHelmChartVersion

	// set loki helm chart values
	loggingDefinition.LokiHelmValuesDocument = c.observabilityStackDefinition.LokiHelmValuesDocument

	// create logging definition
	createdLoggingDefinition, err := client.CreateLoggingDefinition(
		c.r.APIClient,
		c.r.APIServer,
		loggingDefinition,
	)
	if err != nil {
		return fmt.Errorf("failed to create logging definition: %w", err)
	}

	// update observability stack definition with logging definition id
	c.observabilityStackDefinition.LoggingDefinitionID = createdLoggingDefinition.ID

	return nil
}

// deleteLoggingDefinition deletes a logging definition.
func (c *ObservabilityStackDefinitionConfig) deleteLoggingDefinition() error {
	// delete logging definition
	if _, err := client.DeleteLoggingDefinition(
		c.r.APIClient,
		c.r.APIServer,
		*c.observabilityStackDefinition.LoggingDefinitionID,
	); err != nil && !errors.Is(err, client_lib.ErrObjectNotFound) {
		return fmt.Errorf("failed to delete logging definition: %w", err)
	}

	return nil
}

// createMetricsDefinition creates a metrics definition.
func (c *ObservabilityStackDefinitionConfig) createMetricsDefinition() error {
	// create metrics definition
	metricsDefinition := &v0.MetricsDefinition{
		Definition: v0.Definition{
			Name: util.Ptr(MetricsName(*c.observabilityStackDefinition.Name)),
		},
	}

	// set kube-prometheus-stack helm chart version
	metricsDefinition.KubePrometheusStackHelmChartVersion = c.observabilityStackDefinition.KubePrometheusStackHelmChartVersion

	// set kube-prometheus-stack helm chart values
	metricsDefinition.KubePrometheusStackHelmValuesDocument = c.observabilityStackDefinition.KubePrometheusStackHelmValuesDocument

	// create metrics definition
	createdMetricsDefinition, err := client.CreateMetricsDefinition(
		c.r.APIClient,
		c.r.APIServer,
		metricsDefinition,
	)
	if err != nil {
		return fmt.Errorf("failed to create metrics definition: %w", err)
	}

	// update observability stack definition with metrics definition id
	c.observabilityStackDefinition.MetricsDefinitionID = createdMetricsDefinition.ID

	return nil
}

// deleteMetricsDefinition deletes a metrics definition.
func (c *ObservabilityStackDefinitionConfig) deleteMetricsDefinition() error {
	// delete metrics definition
	if _, err := client.DeleteMetricsDefinition(
		c.r.APIClient,
		c.r.APIServer,
		*c.observabilityStackDefinition.MetricsDefinitionID,
	); err != nil && !errors.Is(err, client_lib.ErrObjectNotFound) {
		return fmt.Errorf("failed to delete metrics definition: %w", err)
	}

	return nil
}

// createMetricsStorageDefinition creates a metrics storage definition.
func (c *ObservabilityStackDefinitionConfig) createMetricsStorageDefinition() error {
	// create metrics storage definition
	metricsStorageDefinition := &v0.MetricsStorageDefinition{
		Definition: v0.Definition{
			Name: util.Ptr(MetricsStorageName(*c.observabilityStackDefinition.Name)),
		},
	}

	// set mimir helm chart version
	metricsStorageDefinition.MimirHelmChartVersion = c.observabilityStackDefinition.MimirHelmChartVersion

	// set mimir helm chart values
	metricsStorageDefinition.MimirHelmValuesDocument = c.observabilityStackDefinition.MimirHelmValuesDocument

	// create metrics storage definition
	createdMetricsStorageDefinition, err := client.CreateMetricsStorageDefinition(
		c.r.APIClient,
		c.r.APIServer,
		metricsStorageDefinition,
	)
	if err != nil {
		return fmt.Errorf("failed to create metrics storage definition: %w", err)
	}

	// update observability stack definition with metrics storage definition id
	c.observabilityStackDefinition.MetricsStorageDefinitionID = createdMetricsStorageDefinition.ID

	return nil
}

// deleteMetricsStorageDefinition deletes a metrics storage definition.
func (c *ObservabilityStackDefinitionConfig) deleteMetricsStorageDefinition() error {
	// delete metrics storage definition
	if _, err := client.DeleteMetricsStorageDefinition(
		c.r.APIClient,
		c.r.APIServer,
		*c.observabilityStackDefinition.MetricsStorageDefinitionID,
	); err != nil && !errors.Is(err, client_lib.ErrObjectNotFound) {
		return fmt.Errorf("failed to delete metrics storage definition: %w", err)
	}

	return nil
}

// createTracingDefinition creates a tracing definition.
func (c *ObservabilityStackDefinitionConfig) createTracingDefinition() error {
	// create tracing definition
	tracingDefinition := &v0.TracingDefinition{
		Definition: v0.Definition{
			Name: util.Ptr(TracingName(*c.observabilityStackDefinition.Name)),
		},
	}

	// set tempo helm chart version
	tracingDefinition.TempoHelmChartVersion = c.observabilityStackDefinition.TempoHelmChartVersion

	// set tempo helm chart values
	tracingDefinition.TempoHelmValuesDocument = c.observabilityStackDefinition.TempoHelmValuesDocument

	// create tracing definition
	createdTracingDefinition, err := client.CreateTracingDefinition(
		c.r.APIClient,
		c.r.APIServer,
		tracingDefinition,
	)
	if err != nil {
		return fmt.Errorf("failed to create tracing definition: %w", err)
	}

	// update observability stack definition with tracing definition id
	c.observabilityStackDefinition.TracingDefinitionID = createdTracingDefinition.ID

	return nil
}

// deleteTracingDefinition deletes a tracing definition.
func (c *ObservabilityStackDefinitionConfig) deleteTracingDefinition() error {
	// delete tracing definition
	if _, err := client.DeleteTracingDefinition(
		c.r.APIClient,
		c.r.APIServer,
		*c.observabilityStackDefinition.TracingDefinitionID,
	); err != nil && !errors.Is(err, client_lib.ErrObjectNotFound) {
		return fmt.Errorf("failed to delete tracing definition: %w", err)
	}

	return nil
}

// createInstrumentationAgentDefinition creates an instrumentation agent definition.
func (c *ObservabilityStackDefinitionConfig) createInstrumentationAgentDefinition() error {
	// create instrumentation agent definition
	instrumentationAgentDefinition := &v0.InstrumentationAgentDefinition{
		Definition: v0.Definition{
			Name: util.Ptr(InstrumentationAgentName(*c.observabilityStackDefinition.Name)),
		},
	}

	// set opentelemetry-collector helm chart version
	instrumentationAgentDefinition.OtelCollectorHelmChartVersion = c.observabilityStackDefinition.OtelCollectorHelmChartVersion

	// set otel agent helm chart values
	instrumentationAgentDefinition.OtelAgentHelmValuesDocument = c.observabilityStackDefinition.OtelAgentHelmValuesDocument

	// create instrumentation agent definition
	createdInstrumentationAgentDefinition, err := client.CreateInstrumentationAgentDefinition(
		c.r.APIClient,
		c.r.APIServer,
		instrumentationAgentDefinition,
	)
	if err != nil {
		return fmt.Errorf("failed to create instrumentation agent definition: %w", err)
	}

	// update observability stack definition with instrumentation agent definition id
	c.observabilityStackDefinition.InstrumentationAgentDefinitionID = createdInstrumentationAgentDefinition.ID

	return nil
}

// deleteInstrumentationAgentDefinition deletes an instrumentation agent definition.
func (c *ObservabilityStackDefinitionConfig) deleteInstrumentationAgentDefinition() error {
	// delete instrumentation agent definition
	if _, err := client.DeleteInstrumentationAgentDefinition(
		c.r.APIClient,
		c.r.APIServer,
		*c.observabilityStackDefinition.InstrumentationAgentDefinitionID,
	); err != nil && !errors.Is(err, client_lib.ErrObjectNotFound) {
		return fmt.Errorf("failed to delete instrumentation agent definition: %w", err)
	}

	return nil
}

// createInstrumentationGatewayDefinition creates an instrumentation gateway definition.
func (c *ObservabilityStackDefinitionConfig) createInstrumentationGatewayDefinition() error {
	// create instrumentation gateway definition
	instrumentationGatewayDefinition := &v0.InstrumentationGatewayDefinition{
		Definition: v0.Definition{
			Name: util.Ptr(InstrumentationGatewayName(*c.observabilityStackDefinition.Name)),
		},
	}

	// set opentelemetry-collector helm chart version
	instrumentationGatewayDefinition.OtelCollectorHelmChartVersion = c.observabilityStackDefinition.OtelCollectorHelmChartVersion

	// set otel gateway helm chart values
	instrumentationGatewayDefinition.OtelGatewayHelmValuesDocument = c.observabilityStackDefinition.OtelGatewayHelmValuesDocument

	// create instrumentation gateway definition
	createdInstrumentationGatewayDefinition, err := client.CreateInstrumentationGatewayDefinition(
		c.r.APIClient,
		c.r.APIServer,
		instrumentationGatewayDefinition,
	)
	if err != nil {
		return fmt.Errorf("failed to create instrumentation gateway definition: %w", err)
	}

	// update observability stack definition with instrumentation gateway definition id
	c.observabilityStackDefinition.InstrumentationGatewayDefinitionID = createdInstrumentationGatewayDefinition.ID

	return nil
}

// deleteInstrumentationGatewayDefinition deletes an instrumentation gateway definition.
func (c *ObservabilityStackDefinitionConfig) deleteInstrumentationGatewayDefinition() error {
	// delete instrumentation gateway definition
	if _, err := client.DeleteInstrumentationGatewayDefinition(
		c.r.APIClient,
		c.r.APIServer,
		*c.observabilityStackDefinition.InstrumentationGatewayDefinitionID,
	); err != nil && !errors.Is(err, client_lib.ErrObjectNotFound) {
		return fmt.Errorf("failed to delete instrumentation gateway definition: %w", err)
	}

	return nil
}

// createInstrumentationBrowserRelayDefinition creates an instrumentation browser relay definition.
func (c *ObservabilityStackDefinitionConfig) createInstrumentationBrowserRelayDefinition() error {
	// create instrumentation browser relay definition
	instrumentationBrowserRelayDefinition := &v0.InstrumentationBrowserRelayDefinition{
		Definition: v0.Definition{
			Name: util.Ptr(InstrumentationBrowserRelayName(*c.observabilityStackDefinition.Name)),
		},
	}

	// set opentelemetry-collector helm chart version
	instrumentationBrowserRelayDefinition.OtelCollectorHelmChartVersion = c.observabilityStackDefinition.OtelCollectorHelmChartVersion

	// set otel browser relay helm chart values
	instrumentationBrowserRelayDefinition.OtelBrowserRelayHelmValuesDocument = c.observabilityStackDefinition.OtelBrowserRelayHelmValuesDocument

	// create instrumentation browser relay definition
	createdInstrumentationBrowserRelayDefinition, err := client.CreateInstrumentationBrowserRelayDefinition(
		c.r.APIClient,
		c.r.APIServer,
		instrumentationBrowserRelayDefinition,
	)
	if err != nil {
		return fmt.Errorf("failed to create instrumentation browser relay definition: %w", err)
	}

	// update observability stack definition with instrumentation browser relay definition id
	c.observabilityStackDefinition.InstrumentationBrowserRelayDefinitionID = createdInstrumentationBrowserRelayDefinition.ID

	return nil
}

// deleteInstrumentationBrowserRelayDefinition deletes an instrumentation browser relay definition.
func (c *ObservabilityStackDefinitionConfig) deleteInstrumentationBrowserRelayDefinition() error {
	// delete instrumentation browser relay definition
	if _, err := client.DeleteInstrumentationBrowserRelayDefinition(
		c.r.APIClient,
		c.r.APIServer,
		*c.observabilityStackDefinition.InstrumentationBrowserRelayDefinitionID,
	); err != nil && !errors.Is(err, client_lib.ErrObjectNotFound) {
		return fmt.Errorf("failed to delete instrumentation browser relay definition: %w", err)
	}

	return nil
}
