package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"go-gin-api/boot"
)

var cmdServer = &cobra.Command{
	Use:   "server",
	Short: "start http service",
	Run: func(cmd *cobra.Command, args []string) {
		mode := gin.ReleaseMode
		if len(args) > 0 {
			mode = args[0]
		}
		gin.SetMode(mode)
		boot.Route.Init()
	},
}
