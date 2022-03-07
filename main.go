package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-gin-api/app/cmd"
	"go-gin-api/boot"
	"os"
)

func main() {
	c := &cobra.Command{
		//命令执行前初始化操作
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			boot.Boot.Init()
		},
	}

	cmd.AddCommand(c)

	err := c.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
