package cmd

import (
	"github.com/spf13/cobra"
	"go-gin-api/pkg/migrate"
)

func migrateRefreshCommand() *cobra.Command {
	rollback := &cobra.Command{
		Use:   "refresh",
		Short: "first rollback all migrations that have been run, then run all migrations",
		Run: func(cmd *cobra.Command, args []string) {
			migrate.Run(migrate.Reset)
			migrate.Run(migrate.Up)
		},
	}

	return rollback
}
