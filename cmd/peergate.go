package cmd

import (
	"fmt"
	//ashutosh
	"github.com/thealonemusk/WarpNet/pkg/trustzone/authprovider"
	"github.com/urfave/cli/v2"
)

func Peergate() *cli.Command {
	return &cli.Command{
		Name:        "peergater",
		Usage:       "peergater ecdsa-genkey",
		Description: `Peergater auth utilities`,
		Subcommands: cli.Commands{
			{
				Name: "ecdsa-genkey",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name: "privkey",
					},
					&cli.BoolFlag{
						Name: "pubkey",
					},
				},
				Action: func(c *cli.Context) error {
					priv, pub, err := ecdsa.GenerateKeys()
					if !c.Bool("privkey") && !c.Bool("pubkey") {
						fmt.Printf("Private key: %s\n", string(priv))
						fmt.Printf("Public key: %s\n", string(pub))
					} else if c.Bool("privkey") {
						fmt.Printf(string(priv))
					} else if c.Bool("pubkey") {
						fmt.Printf(string(pub))
					}
					return err
				},
			},
		},
	}
}
