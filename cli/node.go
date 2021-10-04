package cli

import (
	"github.com/kinvolk/lokomotive-update-controller/pkg/updater"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var containerCmd = &cobra.Command{
	Use:   "node",
	Short: "Manage node updates",
	Run:   runController,
}

var (
	docker  bool
	envPath string
)

func init() {
	RootCmd.AddCommand(containerCmd)
	containerCmd.Flags().BoolVarP(&docker, "docker", "d", false, "enable docker mode.")
	containerCmd.Flags().StringVar(&envPath, "envpath", "", "env file path for systemd service.")
	containerCmd.MarkFlagRequired("envpath")
}

func runController(cmd *cobra.Command, args []string) {

	cfg := updater.Config{
		Kubeconfig:     kubeconfig,
		ApplicationID:  appId,
		Interval:       interval,
		Dev:            dev,
		NebraskaServer: nebraskaServer,
		Channel:        channel,
		Docker:         docker,
		EnvPath:        envPath,
	}

	if verbose {
		log.SetLevel(log.DebugLevel)
	}

	if err := updater.ReconcileContainer(&cfg); err != nil {
		log.Fatalf("reconciling: %v", err)
	}
}
