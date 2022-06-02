package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"go-gin-api/boot"
)

type serverOptions struct {
	Mode string
}

func serverCommand() *cobra.Command {
	opt := serverOptions{}
	server := &cobra.Command{
		Use:   "server",
		Short: "start http service",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			boot.Boot.Init()
		},
		Run: func(cmd *cobra.Command, args []string) {
			runServer(opt)
		},
	}

	flags := server.Flags()
	flags.StringVarP(&opt.Mode, "mode", "m", gin.ReleaseMode, "运行模式")

	return server
}

func runServer(opt serverOptions) {
	gin.SetMode(opt.Mode)
	boot.Route.Init()
}
