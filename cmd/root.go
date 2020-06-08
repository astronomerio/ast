package cmd

import (
	"io"

	"github.com/astronomer/astro-cli/houston"
	"github.com/spf13/cobra"
)

var (
	workspaceId   string
	workspaceRole string
	role          string
)

// NewRootCmd adds all of the primary commands for the cli
func NewRootCmd(client *houston.Client, out io.Writer) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "astro",
		Short: "Astronomer - CLI",
		Long:  "astro is a command line interface for working with the Astronomer Platform.",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			PersistentPreRunCheck(client, cmd, out)
		},
	}

	rootCmd.AddCommand(
		newAuthRootCmd(client, out),
		newWorkspaceCmd(client, out),
		newVersionCmd(out),
		newUpgradeCheckCmd(out),
		newUserCmd(client, out),
		newClusterRootCmd(client, out),
		newDevRootCmd(client, out),
		newCompletionCmd(client, out),
		newConfigRootCmd(client, out),
		newDeploymentRootCmd(client, out),
		newDeployCmd(client, out),
		newSaRootCmd(client, out),
		// TODO: remove newAirflowRootCmd, after 1.0 we have only devRootCmd
		newAirflowRootCmd(client, out),
		newLogsDeprecatedCmd(client, out),
	)
	return rootCmd
}
