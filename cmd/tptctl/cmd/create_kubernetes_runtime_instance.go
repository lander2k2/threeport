/*
Copyright © 2023 Threeport admin@threeport.io
*/
package cmd

import (
	"fmt"
	"net/http"
	"os"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"

	cli "github.com/threeport/threeport/pkg/cli/v0"
	config "github.com/threeport/threeport/pkg/config/v0"
)

var createKubernetesRuntimeInstanceConfigPath string

// CreateKubernetesRuntimeInstanceCmd represents the kubernetes-runtime-instance command
var CreateKubernetesRuntimeInstanceCmd = &cobra.Command{
	Use:          "kubernetes-runtime-instance",
	Example:      "tptctl create kubernetes-runtime-instance --config /path/to/config.yaml",
	Short:        "Create a new kubernetes runtime instance",
	Long:         `Create a new kubernetes runtime instance.`,
	SilenceUsage: true,
	Run: func(cmd *cobra.Command, args []string) {
		// get threeport config and extract threeport API endpoint
		threeportConfig, requestedControlPlane, err := config.GetThreeportConfig(cliArgs.ControlPlaneName)
		if err != nil {
			cli.Error("failed to get threeport config", err)
			os.Exit(1)
		}

		var apiClient *http.Client
		var apiEndpoint string

		apiClient, apiEndpoint = checkContext(cmd)
		if apiClient == nil && apiEndpoint != "" {
			apiEndpoint, err = threeportConfig.GetThreeportAPIEndpoint(requestedControlPlane)
			if err != nil {
				cli.Error("failed to get threeport API endpoint from config", err)
				os.Exit(1)
			}

			apiClient, err = threeportConfig.GetHTTPClient(requestedControlPlane)
			if err != nil {
				cli.Error("failed to create threeport API client", err)
				os.Exit(1)
			}
		}

		// load kubernetes runtime instance config
		configContent, err := os.ReadFile(createKubernetesRuntimeInstanceConfigPath)
		if err != nil {
			cli.Error("failed to read config file", err)
			os.Exit(1)
		}
		var kubernetesRuntimeInstanceConfig config.KubernetesRuntimeInstanceConfig
		if err := yaml.Unmarshal(configContent, &kubernetesRuntimeInstanceConfig); err != nil {
			cli.Error("failed to unmarshal config file yaml content", err)
			os.Exit(1)
		}

		// create kubernetes runtime instance
		kubernetesRuntimeInstance := kubernetesRuntimeInstanceConfig.KubernetesRuntimeInstance
		kri, err := kubernetesRuntimeInstance.Create(apiClient, apiEndpoint)
		if err != nil {
			cli.Error("failed to create kubernetes runtime instance", err)
			os.Exit(1)
		}

		cli.Complete(fmt.Sprintf("kubernetes runtime instance %s created\n", *kri.Name))
	},
}

func init() {
	CreateCmd.AddCommand(CreateKubernetesRuntimeInstanceCmd)

	CreateKubernetesRuntimeInstanceCmd.Flags().StringVarP(
		&createKubernetesRuntimeInstanceConfigPath,
		"config", "c", "", "Path to file with kubernetes runtime instance config.",
	)
	CreateKubernetesRuntimeInstanceCmd.MarkFlagRequired("config")
	CreateKubernetesRuntimeInstanceCmd.Flags().StringVarP(
		&cliArgs.ControlPlaneName,
		"control-plane-name", "i", "", "Optional. Name of control plane. Will default to current control plane if not provided.",
	)
}
