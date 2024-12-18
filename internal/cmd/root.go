package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	fPostgresURL string
	fRedisURL    string
)

var RootCmd = &cobra.Command{
	Use: "openatlas",
}

func init() {
	RootCmd.PersistentFlags().StringVar(&fPostgresURL, "postgres-url", os.Getenv("POSTGRES_URL"), "")
	RootCmd.PersistentFlags().StringVar(&fRedisURL, "redis-url", os.Getenv("REDIS_URL"), "")
}
