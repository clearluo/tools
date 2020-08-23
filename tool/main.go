package main

import (
	"os"
	"tool/common/basic/runBefore"
	"tool/dbd"
	"tool/finance"

	"github.com/urfave/cli"
)

func init() {
	runBefore.InitRun()
}

func main() {
	app := cli.NewApp()
	app.Name = "tool"
	app.Author = "clearluo"
	app.Version = "1.0.0"
	app.Usage = "./tool dbd|finance"
	app.Commands = []cli.Command{
		cli.Command{
			Name:   "dbd",
			Action: doDbd,
		},
		cli.Command{
			Name:   "finance",
			Action: doFinance,
		},
	}
	app.Run(os.Args)
}
func doDbd(ctx *cli.Context) error {
	dbd.Dbd()
	return nil
}
func doFinance(ctx *cli.Context) error {
	finance.Finance()
	return nil
}
