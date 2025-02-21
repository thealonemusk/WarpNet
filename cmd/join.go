package cmd

import (
	"context"

	"github.com/thealonemusk/WarpNet/pkg/node"
	"github.com/urfave/cli/v2"
)

func Start() *cli.Command {
	return &cli.Command{
		Name:  "start",
		Usage: "Start the network without activating any interface",
		Description: `Connect over the p2p network without establishing a VPN.
Useful for setting up relays or hop nodes to improve the network connectivity.`,
		UsageText: "edgevpn start",
		Flags:     CommonFlags,
		Action: func(c *cli.Context) error {
			o, _, ll := cliToOpts(c)
			e, err := node.New(o...)
			if err != nil {
				return err
			}

			displayStart(ll)
			go handleStopSignals()

			// Start the node to the network, using our ledger
			if err := e.Start(context.Background()); err != nil {
				return err
			}

			ll.Info("Joining p2p network")
			<-context.Background().Done()
			return nil
		},
	}
}
