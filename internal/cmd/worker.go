package cmd

import (
	"fmt"

	"github.com/KevinSnyderCodes/OpenAtlas/internal/db"
	"github.com/KevinSnyderCodes/OpenAtlas/internal/tasks"
	"github.com/avast/retry-go"
	"github.com/hibiken/asynq"
	"github.com/jackc/pgx/v5"
	"github.com/spf13/cobra"
)

var WorkerCmd = &cobra.Command{
	Use: "worker",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()

		var conn *pgx.Conn
		retry.Do(func() error {
			var err error

			conn, err = pgx.Connect(ctx, fPostgresURL)
			if err != nil {
				return fmt.Errorf("error connecting to postgres: %w", err)
			}

			return nil
		})

		queries := db.New(conn)

		asynqHandler := tasks.NewHandler(queries)

		asyncServer := asynq.NewServer(asynq.RedisClientOpt{Addr: fRedisURL}, asynq.Config{})
		asyncMux := asynq.NewServeMux()
		asyncMux.HandleFunc(tasks.TypeRunProcess, asynqHandler.HandleRunProcessTask)

		go func() {
			if err := asyncServer.Run(asyncMux); err != nil {
				panic(fmt.Errorf("error running async server: %w", err))
			}
		}()

		ch := make(chan struct{})
		<-ch

		return nil
	},
}

func init() {
	RootCmd.AddCommand(WorkerCmd)
}
