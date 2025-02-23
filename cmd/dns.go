package cmd

import (
	"context"
	"time"

	"github.com/thealonemusk/WarpNet/pkg/node"
	"github.com/thealonemusk/WarpNet/pkg/services"
	"github.com/urfave/cli/v2"
)

func DNS() *cli.Command {
	return &cli.Command{
		Name:        "dns",
		Usage:       "Starts a local dns server",
		Description: `Start a local dns server which uses the blockchain to resolve addresses`,
		UsageText:   "WarpNet dns",
		Flags: append(CommonFlags,
			&cli.StringFlag{
				Name:    "listen",
				Usage:   "DNS listening address. Empty to disable dns server",
				EnvVars: []string{"DNSADDRESS"},
				Value:   "",
			},
			&cli.BoolFlag{
				Name:    "dns-forwarder",
				Usage:   "Enables dns forwarding",
				EnvVars: []string{"DNSFORWARD"},
				Value:   true,
			},
			&cli.IntFlag{
				Name:    "dns-cache-size",
				Usage:   "DNS LRU cache size",
				EnvVars: []string{"DNSCACHESIZE"},
				Value:   200,
			},
			&cli.StringSliceFlag{
				Name:    "dns-forward-server",
				Usage:   "List of DNS forward server, e.g. 8.8.8.8:53, 192.168.1.1:53 ...",
				EnvVars: []string{"DNSFORWARDSERVER"},
				Value:   cli.NewStringSlice("8.8.8.8:53", "1.1.1.1:53"),
			},
		),
		Action: func(c *cli.Context) error {
			o, _, ll := cliToOpts(c)

			dns := c.String("listen")
			// Adds DNS Server
			o = append(o,
				services.DNS(ll, dns,
					c.Bool("dns-forwarder"),
					c.StringSlice("dns-forward-server"),
					c.Int("dns-cache-size"),
				)...)

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

			for {
				time.Sleep(1 * time.Second)
			}
		},
	}
}
