// generated by 'threeport-sdk gen' - do not edit

package util

import (
	"fmt"
	nats "github.com/nats-io/nats.go"
	aws_notif "github.com/threeport/threeport/internal/aws/notif"
	controlplane_notif "github.com/threeport/threeport/internal/control-plane/notif"
	gateway_notif "github.com/threeport/threeport/internal/gateway/notif"
	helmworkload_notif "github.com/threeport/threeport/internal/helm-workload/notif"
	kubernetesruntime_notif "github.com/threeport/threeport/internal/kubernetes-runtime/notif"
	observability_notif "github.com/threeport/threeport/internal/observability/notif"
	secret_notif "github.com/threeport/threeport/internal/secret/notif"
	terraform_notif "github.com/threeport/threeport/internal/terraform/notif"
	workload_notif "github.com/threeport/threeport/internal/workload/notif"
)

// Initialize the NATS Jet stream context with controller streams
func InitJetStream(nc *nats.Conn) (*nats.JetStreamContext, error) {
	js, err := nc.JetStream(nats.PublishAsyncMaxPending(256))
	if err != nil {
		return nil, fmt.Errorf("failed to create jetstream context: %w", err)
	}

	// add controller streams
	_, err = js.AddStream(&nats.StreamConfig{
		Name:     secret_notif.SecretStreamName,
		Subjects: secret_notif.GetSecretSubjects(),
	})
	if err != nil {
		return nil, fmt.Errorf("could not add stream %s: %w", secret_notif.SecretStreamName, err)
	}

	_, err = js.AddStream(&nats.StreamConfig{
		Name:     aws_notif.AwsStreamName,
		Subjects: aws_notif.GetAwsSubjects(),
	})
	if err != nil {
		return nil, fmt.Errorf("could not add stream %s: %w", aws_notif.AwsStreamName, err)
	}

	_, err = js.AddStream(&nats.StreamConfig{
		Name:     controlplane_notif.ControlPlaneStreamName,
		Subjects: controlplane_notif.GetControlPlaneSubjects(),
	})
	if err != nil {
		return nil, fmt.Errorf("could not add stream %s: %w", controlplane_notif.ControlPlaneStreamName, err)
	}

	_, err = js.AddStream(&nats.StreamConfig{
		Name:     gateway_notif.GatewayStreamName,
		Subjects: gateway_notif.GetGatewaySubjects(),
	})
	if err != nil {
		return nil, fmt.Errorf("could not add stream %s: %w", gateway_notif.GatewayStreamName, err)
	}

	_, err = js.AddStream(&nats.StreamConfig{
		Name:     helmworkload_notif.HelmWorkloadStreamName,
		Subjects: helmworkload_notif.GetHelmWorkloadSubjects(),
	})
	if err != nil {
		return nil, fmt.Errorf("could not add stream %s: %w", helmworkload_notif.HelmWorkloadStreamName, err)
	}

	_, err = js.AddStream(&nats.StreamConfig{
		Name:     kubernetesruntime_notif.KubernetesRuntimeStreamName,
		Subjects: kubernetesruntime_notif.GetKubernetesRuntimeSubjects(),
	})
	if err != nil {
		return nil, fmt.Errorf("could not add stream %s: %w", kubernetesruntime_notif.KubernetesRuntimeStreamName, err)
	}

	_, err = js.AddStream(&nats.StreamConfig{
		Name:     observability_notif.ObservabilityStreamName,
		Subjects: observability_notif.GetObservabilitySubjects(),
	})
	if err != nil {
		return nil, fmt.Errorf("could not add stream %s: %w", observability_notif.ObservabilityStreamName, err)
	}

	_, err = js.AddStream(&nats.StreamConfig{
		Name:     terraform_notif.TerraformStreamName,
		Subjects: terraform_notif.GetTerraformSubjects(),
	})
	if err != nil {
		return nil, fmt.Errorf("could not add stream %s: %w", terraform_notif.TerraformStreamName, err)
	}

	_, err = js.AddStream(&nats.StreamConfig{
		Name:     workload_notif.WorkloadStreamName,
		Subjects: workload_notif.GetWorkloadSubjects(),
	})
	if err != nil {
		return nil, fmt.Errorf("could not add stream %s: %w", workload_notif.WorkloadStreamName, err)
	}

	return &js, nil
}
