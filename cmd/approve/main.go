package main

import (
	"fmt"
	"os"

	"github.com/weijun-sh/approve/cmd/utils"
	"github.com/weijun-sh/approve/worker"
	"github.com/urfave/cli/v2"
)

var (
	clientIdentifier = "autotrade"
	// Git SHA1 commit hash of the release (set via linker flags)
	gitCommit = ""
	// The app that holds all commands and flags.
	app = utils.NewApp(clientIdentifier, gitCommit, "the autotrade command line interface")
)

func initApp() {
	app.Action = approve
	app.HideVersion = true // we have a command to print the version
	app.Copyright = "Copyright 2017-2020 The Anyswap Authors"
	app.Commands = []*cli.Command{
		//byLiquidityCommand,
		//byVolumeCommand,
		//calcRewardsCommand,
		//sendRewardsCommand,
		//importRewardsCommand,
		//insertAccountCommand,
		utils.LicenseCommand,
		utils.VersionCommand,
	}
	app.Flags = []cli.Flag{
		utils.ConfigFileFlag,
		utils.ActionFlag,
		utils.KeyStoreFileFlag,
		utils.PasswordFileFlag,
		utils.PairFlag,
		utils.AmountFlag,
		//utils.SyncFromFlag,
		//utils.SyncToFlag,
		//utils.OverwriteFlag,
		//utils.OnlySyncAccountFlag,
		//utils.VerbosityFlag,
		//utils.LogFileFlag,
		//utils.LogRotationFlag,
		//utils.LogMaxAgeFlag,
		//utils.JSONFormatFlag,
		//utils.ColorFormatFlag,
	}
}

func main() {
	initApp()
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func approve(ctx *cli.Context) error {
	if ctx.NArg() > 0 {
		return fmt.Errorf("invalid command: %q", ctx.Args().Get(0))
	}

	capi := utils.InitApp(ctx)
	defer capi.CloseClient()

	worker.Start()
	return nil
}
