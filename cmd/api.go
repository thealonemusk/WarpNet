package cmd

import (
	"context"
	"time"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/metrics"
	"github.com/thealonemusk/WarpNet/api"
	"github.com/thealonemusk/WarpNet/pkg/node"
	"github.com/thealonemusk/WarpNet/pkg/services"
	"github.com/urfave/cli/v2"
)

func API() *cli.Command {
	return &cli.Command{
		Name:  "api",
		Usage: "Starts an http server to display network informations",
		Description: `Start listening locally, providing an API for the network.
A simple UI interface is available to display network data.`,
		UsageText: "WarpNet api",
		Flags: append(CommonFlags,
			&cli.BoolFlag{
				Name:    "enable-healthchecks",
				EnvVars: []string{"ENABLE_HEALTHCHECKS"},
			},
			&cli.BoolFlag{
				Name: "debug",
			},
			&cli.StringFlag{
				Name:  "listen",
				Value: "127.0.0.1:8080",
				Usage: "Listening address. To listen to a socket, prefix with unix://, e.g. unix:///socket.path",
			},
		),
		Action: func(c *cli.Context) error {
			o, _, ll := cliToOpts(c)

			bwc := metrics.NewBandwidthCounter()
			o = append(o, node.WithLibp2pAdditionalOptions(libp2p.BandwidthReporter(bwc)))
			if c.Bool("enable-healthchecks") {
				o = append(o,
					services.Alive(
						time.Duration(c.Int("aliveness-healthcheck-interval"))*time.Second,
						time.Duration(c.Int("aliveness-healthcheck-scrub-interval"))*time.Second,
						time.Duration(c.Int("aliveness-healthcheck-max-interval"))*time.Second)...)
			}

			e, err := node.New(o...)
			if err != nil {
				return err
			}

			displayStart(ll)

			ctx := context.Background()
			go handleStopSignals()

			// Start the node to the network, using our ledger
			if err := e.Start(ctx); err != nil {
				return err
			}

			return api.API(ctx, c.String("listen"), 5*time.Second, 20*time.Second, e, bwc, c.Bool("debug"))
		},
	}
}
