package cli

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/wajox/krakend-proxy/internal/app"
)

var ServeCmd = &cobra.Command{
	Use:     "serve",
	Aliases: []string{"s"},
	Short:   "Start server",
	Run: func(cmd *cobra.Command, args []string) {
		log.Info().Msg("Starting")

		sigchan := make(chan os.Signal, 1)
		signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		application, err := app.InitializeApplication()
		if err != nil {
			log.Fatal().Err(err).Msg("can not initialize application")
		}

		application.Start(ctx, false)
		defer func() {
			log.Info().Err(application.Stop(false)).Msg("Finishing")
		}()

		log.Info().Msg("Starting finished")

		for {
			select {
			case <-application.ReloadCh:
				log.Info().Msg("Restarting")
				cancel()
				ctx, cancel = context.WithCancel(context.Background())
				defer cancel()

				newApp, err := application.Reload(ctx, false)
				if err != nil {
					log.Fatal().Err(err).Msg("can not reload application")
				}
				application = newApp

				log.Info().Msg("Restarted")
			case <-sigchan:
				log.Info().Msg("Terminating")

				return
			}
		}
	},
}
