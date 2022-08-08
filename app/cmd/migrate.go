package cmd

import (
	"github.com/spf13/cobra"
	"go-simple/database/migrations"
	"go-simple/pkg/migrate"
)

var Migrate = &cobra.Command{
	Use: "migrate",
	Short: "Run database migration",
}

var MigrateUp = &cobra.Command{
	Use: "up",
	Short: "Run unmigrated migrations",
	Run: runUp,
}

var MigrateRollback = &cobra.Command{
	Use: "down",
	Aliases: []string{"rollback"},
	Short: "Reverse the up command",
	Run: runDown,
}

var MigrateReset = &cobra.Command{
	Use: "reset",
	Short: "Rollback all database migrations",
	Run: runReset,
}

var MigrateRefresh = &cobra.Command{
	Use: "refresh",
	Short: "Reset and re-run all migrations",
	Run: runRefresh,
}

var MigrateFresh = &cobra.Command{
	Use: "fresh",
	Short: "Drop all tables and re-run all migrations",
	Run: runFresh,
}


func init() {
	Migrate.AddCommand(
			MigrateUp,
			MigrateRollback,
			MigrateReset,
			MigrateRefresh,
			MigrateFresh,
		)
}

func migrator() *migrate.Migrator {
	// Init All Migration
	migrations.Initialize()
	return migrate.NewMigrator()
}

func runUp(cmd *cobra.Command, args []string)  {
	migrator().Up()
}

func runDown(cmd *cobra.Command, args []string) {
	migrator().Rollback()
}

func runReset(cmd *cobra.Command, args []string) {
	migrator().Reset()
}

func runRefresh(cmd *cobra.Command, args []string) {
	migrator().Refresh()
}

func runFresh(cmd *cobra.Command, args []string) {
	migrator().Fresh()
}