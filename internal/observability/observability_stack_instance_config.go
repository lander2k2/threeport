package observability

import (
	"errors"
	"fmt"

	"github.com/go-logr/logr"
	helmworkload "github.com/threeport/threeport/internal/helm-workload"
	v0 "github.com/threeport/threeport/pkg/api/v0"
	client_lib "github.com/threeport/threeport/pkg/client/lib/v0"
	client "github.com/threeport/threeport/pkg/client/v0"
	controller "github.com/threeport/threeport/pkg/controller/v0"
	util "github.com/threeport/threeport/pkg/util/v0"
)

// ObservabilityStackInstanceConfig contains the configuration for
// an observability stack instance reconcile function.
type ObservabilityStackInstanceConfig struct {
	r                                     *controller.Reconciler
	observabilityStackInstance            *v0.ObservabilityStackInstance
	observabilityStackDefinition          *v0.ObservabilityStackDefinition
	log                                   *logr.Logger
	grafanaHelmValuesDocument             string
	kubePrometheusStackHelmValuesDocument string
	mimirHelmValuesDocument               string
	lokiHelmValuesDocument                string
	tempoHelmValuesDocument               string
	otelAgentHelmValuesDocument           string
	otelGatewayHelmValuesDocument         string
	otelBrowserRelayHelmValuesDocument    string
}

// getObservabilityStackInstanceOperations returns the operations
// for an observabiblity stack instance
func (c *ObservabilityStackInstanceConfig) getObservabilityStackInstanceOperations() *util.Operations {
	operations := util.Operations{}

	// append observability dashboard operations
	operations.AppendOperation(util.Operation{
		Name:   "observability dashboard",
		Create: c.createObservabilityDashboardInstance,
		Delete: c.deleteObservabilityDashboardInstance,
	})

	if *c.observabilityStackInstance.LoggingEnabled {
		// append logging operations
		operations.AppendOperation(util.Operation{
			Name:   "logging",
			Create: c.createLoggingInstance,
			Delete: c.deleteLoggingInstance,
		})
	}

	if *c.observabilityStackInstance.MetricsEnabled {
		// append metrics operations
		operations.AppendOperation(util.Operation{
			Name:   "metrics",
			Create: c.createMetricsInstance,
			Delete: c.deleteMetricsInstance,
		})
	}

	if *c.observabilityStackInstance.MetricsStorageEnabled {
		// append metrics storage operations
		operations.AppendOperation(util.Operation{
			Name:   "metrics storage",
			Create: c.createMetricsStorageInstance,
			Delete: c.deleteMetricsStorageInstance,
		})
	}

	if *c.observabilityStackInstance.TracingEnabled {
		// append tracing operations
		operations.AppendOperation(util.Operation{
			Name:   "tracing",
			Create: c.createTracingInstance,
			Delete: c.deleteTracingInstance,
		})
	}

	if *c.observabilityStackInstance.InstrumentationAgentEnabled {
		// append instrumentation agent operations
		operations.AppendOperation(util.Operation{
			Name:   "instrumentation agent",
			Create: c.createInstrumentationAgentInstance,
			Delete: c.deleteInstrumentationAgentInstance,
		})
	}

	if *c.observabilityStackInstance.InstrumentationGatewayEnabled {
		// append instrumentation gateway operations
		operations.AppendOperation(util.Operation{
			Name:   "instrumentation gateway",
			Create: c.createInstrumentationGatewayInstance,
			Delete: c.deleteInstrumentationGatewayInstance,
		})
	}

	if *c.observabilityStackInstance.InstrumentationBrowserRelayEnabled {
		// append instrumentation browser relay operations
		operations.AppendOperation(util.Operation{
			Name:   "instrumentation browser relay",
			Create: c.createInstrumentationBrowserRelayInstance,
			Delete: c.deleteInstrumentationBrowserRelayInstance,
		})
	}

	return &operations
}

// createObservabilityDashboardInstance creates an observability dashboard instance
func (c *ObservabilityStackInstanceConfig) createObservabilityDashboardInstance() error {
	// create observability dashboard instance
	observabilityDashboardInstance, err := client.CreateObservabilityDashboardInstance(
		c.r.APIClient,
		c.r.APIServer,
		&v0.ObservabilityDashboardInstance{
			Instance: v0.Instance{
				Name: util.Ptr(ObservabilityDashboardName(*c.observabilityStackInstance.Name)),
			},
			KubernetesRuntimeInstanceID:        c.observabilityStackInstance.KubernetesRuntimeInstanceID,
			ObservabilityDashboardDefinitionID: c.observabilityStackDefinition.ObservabilityDashboardDefinitionID,
			GrafanaHelmValuesDocument:          &c.grafanaHelmValuesDocument,
		})
	if err != nil {
		return fmt.Errorf("failed to create observability dashboard instance: %w", err)
	}

	// update observability dashboard instance id
	c.observabilityStackInstance.ObservabilityDashboardInstanceID = observabilityDashboardInstance.ID

	return nil
}

// deleteObservabilityDashboardInstance deletes an observability dashboard instance
func (c *ObservabilityStackInstanceConfig) deleteObservabilityDashboardInstance() error {
	// delete observability dashboard instance
	if _, err := client.DeleteObservabilityDashboardInstance(
		c.r.APIClient,
		c.r.APIServer,
		*c.observabilityStackInstance.ObservabilityDashboardInstanceID,
	); err != nil && !errors.Is(err, client_lib.ErrObjectNotFound) {
		return fmt.Errorf("failed to delete observability dashboard instance: %w", err)
	}

	return nil
}

// createMetricsInstance creates a metrics instance
func (c *ObservabilityStackInstanceConfig) createMetricsInstance() error {
	// create metrics instance
	metricsInstance, err := client.CreateMetricsInstance(
		c.r.APIClient,
		c.r.APIServer,
		&v0.MetricsInstance{
			Instance: v0.Instance{
				Name: util.Ptr(MetricsName(*c.observabilityStackInstance.Name)),
			},
			KubernetesRuntimeInstanceID:           c.observabilityStackInstance.KubernetesRuntimeInstanceID,
			MetricsDefinitionID:                   c.observabilityStackDefinition.MetricsDefinitionID,
			KubePrometheusStackHelmValuesDocument: &c.kubePrometheusStackHelmValuesDocument,
		})
	if err != nil {
		return fmt.Errorf("failed to create metrics instance: %w", err)
	}

	// update metrics instance id
	c.observabilityStackInstance.MetricsInstanceID = metricsInstance.ID

	return nil
}

// deleteMetricsInstance deletes a metrics instance
func (c *ObservabilityStackInstanceConfig) deleteMetricsInstance() error {
	// delete metrics instance
	if _, err := client.DeleteMetricsInstance(
		c.r.APIClient,
		c.r.APIServer,
		*c.observabilityStackInstance.MetricsInstanceID,
	); err != nil && !errors.Is(err, client_lib.ErrObjectNotFound) {
		return fmt.Errorf("failed to delete metrics instance: %w", err)
	}

	return nil
}

// createLoggingInstance creates a logging instance
func (c *ObservabilityStackInstanceConfig) createLoggingInstance() error {
	// create logging instance
	loggingInstance, err := client.CreateLoggingInstance(
		c.r.APIClient,
		c.r.APIServer,
		&v0.LoggingInstance{
			Instance: v0.Instance{
				Name: util.Ptr(LoggingName(*c.observabilityStackInstance.Name)),
			},
			KubernetesRuntimeInstanceID: c.observabilityStackInstance.KubernetesRuntimeInstanceID,
			LoggingDefinitionID:         c.observabilityStackDefinition.LoggingDefinitionID,
			LokiHelmValuesDocument:      &c.lokiHelmValuesDocument,
		})
	if err != nil {
		return fmt.Errorf("failed to create logging instance: %w", err)
	}

	// update logging instance id
	c.observabilityStackInstance.LoggingInstanceID = loggingInstance.ID

	return nil
}

// deleteLoggingInstance deletes a logging instance
func (c *ObservabilityStackInstanceConfig) deleteLoggingInstance() error {
	// delete logging instance
	if _, err := client.DeleteLoggingInstance(
		c.r.APIClient,
		c.r.APIServer,
		*c.observabilityStackInstance.LoggingInstanceID,
	); err != nil && !errors.Is(err, client_lib.ErrObjectNotFound) {
		return fmt.Errorf("failed to delete logging instance: %w", err)
	}

	return nil
}

// createMetricsStorageInstance creates a metrics storage instance
func (c *ObservabilityStackInstanceConfig) createMetricsStorageInstance() error {
	// create metrics storage instance
	metricsStorageInstance, err := client.CreateMetricsStorageInstance(
		c.r.APIClient,
		c.r.APIServer,
		&v0.MetricsStorageInstance{
			Instance: v0.Instance{
				Name: util.Ptr(MetricsStorageName(*c.observabilityStackInstance.Name)),
			},
			KubernetesRuntimeInstanceID: c.observabilityStackInstance.KubernetesRuntimeInstanceID,
			MetricsStorageDefinitionID:  c.observabilityStackDefinition.MetricsStorageDefinitionID,
			MimirHelmValuesDocument:     &c.mimirHelmValuesDocument,
		})
	if err != nil {
		return fmt.Errorf("failed to create metrics storage instance: %w", err)
	}

	// update metrics storage instance id
	c.observabilityStackInstance.MetricsStorageInstanceID = metricsStorageInstance.ID

	return nil
}

// deleteMetricsStorageInstance deletes a metrics storage instance
func (c *ObservabilityStackInstanceConfig) deleteMetricsStorageInstance() error {
	// delete metrics storage instance
	if _, err := client.DeleteMetricsStorageInstance(
		c.r.APIClient,
		c.r.APIServer,
		*c.observabilityStackInstance.MetricsStorageInstanceID,
	); err != nil && !errors.Is(err, client_lib.ErrObjectNotFound) {
		return fmt.Errorf("failed to delete metrics storage instance: %w", err)
	}

	return nil
}

// createTracingInstance creates a tracing instance
func (c *ObservabilityStackInstanceConfig) createTracingInstance() error {
	// create tracing instance
	tracingInstance, err := client.CreateTracingInstance(
		c.r.APIClient,
		c.r.APIServer,
		&v0.TracingInstance{
			Instance: v0.Instance{
				Name: util.Ptr(TracingName(*c.observabilityStackInstance.Name)),
			},
			KubernetesRuntimeInstanceID: c.observabilityStackInstance.KubernetesRuntimeInstanceID,
			TracingDefinitionID:         c.observabilityStackDefinition.TracingDefinitionID,
			TempoHelmValuesDocument:     &c.tempoHelmValuesDocument,
		})
	if err != nil {
		return fmt.Errorf("failed to create tracing instance: %w", err)
	}

	// update tracing instance id
	c.observabilityStackInstance.TracingInstanceID = tracingInstance.ID

	return nil
}

// deleteTracingInstance deletes a tracing instance
func (c *ObservabilityStackInstanceConfig) deleteTracingInstance() error {
	// delete tracing instance
	if _, err := client.DeleteTracingInstance(
		c.r.APIClient,
		c.r.APIServer,
		*c.observabilityStackInstance.TracingInstanceID,
	); err != nil && !errors.Is(err, client_lib.ErrObjectNotFound) {
		return fmt.Errorf("failed to delete tracing instance: %w", err)
	}

	return nil
}

// createInstrumentationAgentInstance creates an instrumentation agent instance
func (c *ObservabilityStackInstanceConfig) createInstrumentationAgentInstance() error {
	// create instrumentation agent instance
	instrumentationAgentInstance, err := client.CreateInstrumentationAgentInstance(
		c.r.APIClient,
		c.r.APIServer,
		&v0.InstrumentationAgentInstance{
			Instance: v0.Instance{
				Name: util.Ptr(InstrumentationAgentName(*c.observabilityStackInstance.Name)),
			},
			KubernetesRuntimeInstanceID:      c.observabilityStackInstance.KubernetesRuntimeInstanceID,
			InstrumentationAgentDefinitionID: c.observabilityStackDefinition.InstrumentationAgentDefinitionID,
			OtelAgentHelmValuesDocument:      &c.otelAgentHelmValuesDocument,
		})
	if err != nil {
		return fmt.Errorf("failed to create instrumentation agent instance: %w", err)
	}

	// update instrumentation agent instance id
	c.observabilityStackInstance.InstrumentationAgentInstanceID = instrumentationAgentInstance.ID

	return nil
}

// deleteInstrumentationAgentInstance deletes an instrumentation agent instance
func (c *ObservabilityStackInstanceConfig) deleteInstrumentationAgentInstance() error {
	// delete instrumentation agent instance
	if _, err := client.DeleteInstrumentationAgentInstance(
		c.r.APIClient,
		c.r.APIServer,
		*c.observabilityStackInstance.InstrumentationAgentInstanceID,
	); err != nil && !errors.Is(err, client_lib.ErrObjectNotFound) {
		return fmt.Errorf("failed to delete instrumentation agent instance: %w", err)
	}

	return nil
}

// createInstrumentationGatewayInstance creates an instrumentation gateway instance
func (c *ObservabilityStackInstanceConfig) createInstrumentationGatewayInstance() error {
	// create instrumentation gateway instance
	instrumentationGatewayInstance, err := client.CreateInstrumentationGatewayInstance(
		c.r.APIClient,
		c.r.APIServer,
		&v0.InstrumentationGatewayInstance{
			Instance: v0.Instance{
				Name: util.Ptr(InstrumentationGatewayName(*c.observabilityStackInstance.Name)),
			},
			KubernetesRuntimeInstanceID:        c.observabilityStackInstance.KubernetesRuntimeInstanceID,
			InstrumentationGatewayDefinitionID: c.observabilityStackDefinition.InstrumentationGatewayDefinitionID,
			OtelGatewayHelmValuesDocument:      &c.otelGatewayHelmValuesDocument,
		})
	if err != nil {
		return fmt.Errorf("failed to create instrumentation gateway instance: %w", err)
	}

	// update instrumentation gateway instance id
	c.observabilityStackInstance.InstrumentationGatewayInstanceID = instrumentationGatewayInstance.ID

	return nil
}

// deleteInstrumentationGatewayInstance deletes an instrumentation gateway instance
func (c *ObservabilityStackInstanceConfig) deleteInstrumentationGatewayInstance() error {
	// delete instrumentation gateway instance
	if _, err := client.DeleteInstrumentationGatewayInstance(
		c.r.APIClient,
		c.r.APIServer,
		*c.observabilityStackInstance.InstrumentationGatewayInstanceID,
	); err != nil && !errors.Is(err, client_lib.ErrObjectNotFound) {
		return fmt.Errorf("failed to delete instrumentation gateway instance: %w", err)
	}

	return nil
}

// createInstrumentationBrowserRelayInstance creates an instrumentation browser relay instance
func (c *ObservabilityStackInstanceConfig) createInstrumentationBrowserRelayInstance() error {
	// create instrumentation browser relay instance
	instrumentationBrowserRelayInstance, err := client.CreateInstrumentationBrowserRelayInstance(
		c.r.APIClient,
		c.r.APIServer,
		&v0.InstrumentationBrowserRelayInstance{
			Instance: v0.Instance{
				Name: util.Ptr(InstrumentationBrowserRelayName(*c.observabilityStackInstance.Name)),
			},
			KubernetesRuntimeInstanceID:             c.observabilityStackInstance.KubernetesRuntimeInstanceID,
			InstrumentationBrowserRelayDefinitionID: c.observabilityStackDefinition.InstrumentationBrowserRelayDefinitionID,
			OtelBrowserRelayHelmValuesDocument:      &c.otelBrowserRelayHelmValuesDocument,
		})
	if err != nil {
		return fmt.Errorf("failed to create instrumentation browser relay instance: %w", err)
	}

	// update instrumentation browser relay instance id
	c.observabilityStackInstance.InstrumentationBrowserRelayInstanceID = instrumentationBrowserRelayInstance.ID

	return nil
}

// deleteInstrumentationBrowserRelayInstance deletes an instrumentation browser relay instance
func (c *ObservabilityStackInstanceConfig) deleteInstrumentationBrowserRelayInstance() error {
	// delete instrumentation browser relay instance
	if _, err := client.DeleteInstrumentationBrowserRelayInstance(
		c.r.APIClient,
		c.r.APIServer,
		*c.observabilityStackInstance.InstrumentationBrowserRelayInstanceID,
	); err != nil && !errors.Is(err, client_lib.ErrObjectNotFound) {
		return fmt.Errorf("failed to delete instrumentation browser relay instance: %w", err)
	}

	return nil
}

// setMergedObservabilityStackInstanceValues sets the merged values for
// an observability stack instance
func (c *ObservabilityStackInstanceConfig) setMergedObservabilityStackInstanceValues() error {
	var err error

	// merge grafana values
	c.grafanaHelmValuesDocument, err = helmworkload.MergeHelmValuesPtrs(
		c.observabilityStackInstance.GrafanaHelmValuesDocument,
		c.observabilityStackDefinition.GrafanaHelmValuesDocument,
	)
	if err != nil {
		return fmt.Errorf("failed to merge grafana helm values: %w", err)
	}

	// Only configure grafana service monitor if metrics are enabled, as
	// this depends on the ServiceMonitor CRD being installed in the cluster
	// and the kube-prometheus-stack being installed to scrape its metrics.
	if *c.observabilityStackInstance.MetricsEnabled {
		// merge grafana prometheus service monitor
		c.grafanaHelmValuesDocument, err = helmworkload.MergeHelmValuesString(
			c.grafanaHelmValuesDocument,
			grafanaPrometheusServiceMonitor,
		)
		if err != nil {
			return fmt.Errorf("failed to merge grafana prometheus service monitor: %w", err)
		}
	}

	// merge kube-prometheus-stack values
	c.kubePrometheusStackHelmValuesDocument, err = helmworkload.MergeHelmValuesPtrs(
		c.observabilityStackInstance.KubePrometheusStackHelmValuesDocument,
		c.observabilityStackDefinition.KubePrometheusStackHelmValuesDocument,
	)
	if err != nil {
		return fmt.Errorf("failed to merge kube-prometheus-stack helm values: %w", err)
	}

	// merge mimir values
	c.mimirHelmValuesDocument, err = helmworkload.MergeHelmValuesPtrs(
		c.observabilityStackInstance.MimirHelmValuesDocument,
		c.observabilityStackDefinition.MimirHelmValuesDocument,
	)
	if err != nil {
		return fmt.Errorf("failed to merge mimir helm values: %w", err)
	}

	// merge loki values
	c.lokiHelmValuesDocument, err = helmworkload.MergeHelmValuesPtrs(
		c.observabilityStackInstance.LokiHelmValuesDocument,
		c.observabilityStackDefinition.LokiHelmValuesDocument,
	)
	if err != nil {
		return fmt.Errorf("failed to merge loki helm values: %w", err)
	}

	// merge tempo values
	c.tempoHelmValuesDocument, err = helmworkload.MergeHelmValuesPtrs(
		c.observabilityStackInstance.TempoHelmValuesDocument,
		c.observabilityStackDefinition.TempoHelmValuesDocument,
	)
	if err != nil {
		return fmt.Errorf("failed to merge tempo helm values: %w", err)
	}

	// merge otel agent values
	c.otelAgentHelmValuesDocument, err = helmworkload.MergeHelmValuesPtrs(
		c.observabilityStackInstance.OtelAgentHelmValuesDocument,
		c.observabilityStackDefinition.OtelAgentHelmValuesDocument,
	)
	if err != nil {
		return fmt.Errorf("failed to merge otel agent helm values: %w", err)
	}

	// merge otel gateway values
	c.otelGatewayHelmValuesDocument, err = helmworkload.MergeHelmValuesPtrs(
		c.observabilityStackInstance.OtelGatewayHelmValuesDocument,
		c.observabilityStackDefinition.OtelGatewayHelmValuesDocument,
	)
	if err != nil {
		return fmt.Errorf("failed to merge otel gateway helm values: %w", err)
	}

	// merge otel browser relay values
	c.otelBrowserRelayHelmValuesDocument, err = helmworkload.MergeHelmValuesPtrs(
		c.observabilityStackInstance.OtelBrowserRelayHelmValuesDocument,
		c.observabilityStackDefinition.OtelBrowserRelayHelmValuesDocument,
	)
	if err != nil {
		return fmt.Errorf("failed to merge otel browser relay helm values: %w", err)
	}

	return nil
}
