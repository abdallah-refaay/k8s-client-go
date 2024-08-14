package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/rest"
)

var configFlags *genericclioptions.ConfigFlags

var rootCmd = &cobra.Command{
	Use:  "kubectl limits",
	RunE: run,
}

func run(cmd *cobra.Command, args []string) error {
	restConfig, err := configFlags.ToRESTConfig()
	if err != nil {
		return err
	}

	restConfig.WarningHandler = rest.NoWarnings{}
	restConfig.QPS = 1000
	restConfig.Burst = 1000

	dyn, err := rest.RESTClientFor(restConfig)
	if err != nil {
		return fmt.Errorf("failed to create dynamic client: %w", err)
	}

}
