package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/long-in/gdcli"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "gdcli"
	app.Usage = "A cli application for Gehirn DNS"
	app.Version = "0.0.1"

	app.Commands = []cli.Command{
		gdcli.ConfigCommand(),
		gdcli.ZoneCommand(),
		gdcli.RecordCommand(),
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	if err := app.Run(os.Args); err != nil {
		fmt.Printf("app.Run(). %s\n", err)
		os.Exit(2)
	}
}
