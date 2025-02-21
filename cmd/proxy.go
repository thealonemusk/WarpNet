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

func Proxy() *cli.Command {
	return &cli.Command{
		Name:        "proxy",
		Usage:       "Starts a local http proxy server to egress nodes",
		Description: `Start a proxy locally, providing an ingress point for the network.`,
		UsageText:   "WarpNet proxy",
		Flags: append(CommonFlags,
			&cli.StringFlag{
				Name:    "listen",
				Value:   ":8080",
				Usage:   "Listening address",
				EnvVars: []string{"PROXYLISTEN"},
			},
			&cli.BoolFlag{
				Name: "debug",
			},
			&cli.IntFlag{
				Name:    "interval",
				Usage:   "proxy announce time interval",
				EnvVars: []string{"PROXYINTERVAL"},
				Value:   120,
			},
			&cli.IntFlag{
				Name:    "dead-interval",
				Usage:   "interval (in seconds) wether detect egress nodes offline",
				EnvVars: []string{"PROXYDEADINTERVAL"},
				Value:   600,
			},
		),
		Action: func(c *cli.Context) error {
			o, _, ll := cliToOpts(c)

			o = append(o, services.Proxy(
				time.Duration(c.Int("interval"))*time.Second,
				time.Duration(c.Int("dead-interval"))*time.Second,
				c.String("listen"))...)

			bwc := metrics.NewBandwidthCounter()
			o = append(o, node.WithLibp2pAdditionalOptions(libp2p.BandwidthReporter(bwc)))

			e, err := node.New(o...)
			if err != nil {
				return err
			}

			displayStart(ll)

			go handleStopSignals()

			ctx := context.Background()
			// Start the node to the network, using our ledger
			if err := e.Start(ctx); err != nil {
				return err
			}

			return api.API(ctx, c.String("listen"), 5*time.Second, 20*time.Second, e, bwc, c.Bool("debug"))
		},
	}
}
