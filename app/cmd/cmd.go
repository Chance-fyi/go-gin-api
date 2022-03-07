package cmd

import (
	"github.com/spf13/cobra"
)

func AddCommand(cmd *cobra.Command) {
	cmd.AddCommand(
		cmdServer,
		cmdTest,
	)
}
