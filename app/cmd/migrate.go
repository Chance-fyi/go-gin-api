package cmd

import (
	"github.com/spf13/cobra"
	"go-gin-api/boot"
	_ "go-gin-api/database/migrations"
	"go-gin-api/pkg/migrate"
)

func migrateCommand() *cobra.Command {
	m := &cobra.Command{
		Use:   "migrate",
		Short: "run all unexecuted migrations",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			boot.Boot.Init()
		},
		Run: func(cmd *cobra.Command, args []string) {
			migrate.Run(migrate.Up)
		},
	}

	m.AddCommand(
		migrateRollbackCommand(),
		migrateResetCommand(),
		migrateRefreshCommand(),
	)

	return m
}