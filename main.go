package main

import (
	"fmt"
	"os"

	"github.com/thealonemusk/WarpNet/cmd"
	internal "github.com/thealonemusk/WarpNet/internal"
	"github.com/urfave/cli/v2"
)

func main() {

	app := &cli.App{
		Name:        "WarpNet",
		Version:     internal.Version,
		Authors:     []*cli.Author{{Name: "Ashutosh Jha"}},
		Usage:       "WarpNet --config /etc/WarpNet/config.yaml",
		Description: "WarpNet uses libp2p to build an immutable trusted blockchain addressable p2p network",
		Copyright:   cmd.Copyright,
		Flags:       cmd.MainFlags(),
		Commands: []*cli.Command{
			cmd.Start(),
			cmd.API(),
			cmd.ServiceAdd(),
			cmd.ServiceConnect(),
			cmd.FileReceive(),
			cmd.Proxy(),
			cmd.FileSend(),
			cmd.DNS(),
			cmd.Peergate(),
		},

		Action: cmd.Main(),
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
