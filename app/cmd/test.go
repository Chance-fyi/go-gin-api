package cmd

import (
	"github.com/spf13/cobra"
)

var cmdTest = &cobra.Command{
	Use:   "test",
	Short: "run temporary test code",
	Run: func(cmd *cobra.Command, args []string) {
		//测试代码调试

	},
}
