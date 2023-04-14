/*
Copyright © 2023 Threeport admin@threeport.io
*/
package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"

	"github.com/threeport/threeport/internal/cli"
	config "github.com/threeport/threeport/pkg/config/v0"
)

var (
	deleteWorkloadConfigPath string
	deleteWorkloadName       string
)

// DeleteWorkloadCmd represents the workload command
var DeleteWorkloadCmd = &cobra.Command{
	Use:     "workload",
	Example: "tptctl delete workload --config /path/to/config.yaml",
	Short:   "Delete a new workload",
	Long: `Delete a new workload. This command deletes an existing workload definition
and workload instance based on the workload config or name.`,
	SilenceUsage: true,
	Run: func(cmd *cobra.Command, args []string) {
		// get threeport config and extract threeport API endpoint
		threeportConfig := &config.ThreeportConfig{}
		if err := viper.Unmarshal(threeportConfig); err != nil {
			cli.Error("Failed to get threeport config", err)
			os.Exit(1)
		}
		apiEndpoint, err := threeportConfig.GetThreeportAPIEndpoint()
		if err != nil {
			cli.Error("failed to get threeport API endpoint from config", err)
			os.Exit(1)
		}

		// flag validation
		if err := validateDeleteWorkloadFlags(
			deleteWorkloadConfigPath,
			deleteWorkloadName,
		); err != nil {
			cli.Error("flag validation failed", err)
			os.Exit(1)
		}

		var workloadConfig config.WorkloadConfig
		if deleteWorkloadConfigPath != "" {
			// load workload definition config
			configContent, err := ioutil.ReadFile(deleteWorkloadConfigPath)
			if err != nil {
				cli.Error("failed to read config file", err)
				os.Exit(1)
			}
			if err := yaml.Unmarshal(configContent, &workloadConfig); err != nil {
				cli.Error("failed to unmarshal config file yaml content", err)
				os.Exit(1)
			}
		} else {
			workloadConfig = config.WorkloadConfig{
				Workload: config.WorkloadValues{
					Name: deleteWorkloadName,
				},
			}
		}

		// delete workload
		workload := workloadConfig.Workload
		wd, wi, err := workload.Delete(apiEndpoint)
		if err != nil {
			cli.Error("failed to delete workload", err)
			os.Exit(1)
		}

		cli.Info(fmt.Sprintf("workload definition %s deleted", *wd.Name))
		cli.Info(fmt.Sprintf("workload instance %s deleted", *wi.Name))
		cli.Complete(fmt.Sprintf("workload %s deleted", workloadConfig.Workload.Name))
	},
}

func init() {
	deleteCmd.AddCommand(DeleteWorkloadCmd)

	DeleteWorkloadCmd.Flags().StringVarP(
		&deleteWorkloadConfigPath,
		"config", "c", "", "Path to file with workload config.",
	)
	DeleteWorkloadCmd.Flags().StringVarP(
		&deleteWorkloadName,
		"name", "n", "", "Name of workload.",
	)
}

// validateDeleteWorkloadFlags validates flag inputs as needed.
func validateDeleteWorkloadFlags(workloadConfigPath, workloadName string) error {
	if workloadConfigPath == "" && workloadName == "" {
		return errors.New("must provide either workload name or path to config file")
	}

	if workloadConfigPath != "" && workloadName != "" {
		return errors.New("workload name and path to config file provided - provide only one")
	}

	return nil
}
