package utils

import (
	"time"

	"github.com/weijun-sh/approve/callapi"
	"github.com/weijun-sh/approve/params"
	"github.com/urfave/cli/v2"
)

var capi *callapi.APICaller

// InitApp init app (remember close client in the caller)
func InitApp(ctx *cli.Context) *callapi.APICaller {
	return initApp(ctx)
}

func initApp(ctx *cli.Context) *callapi.APICaller {
	SetLogger(ctx)

	//InitSyncArguments(ctx)

	configFile := GetConfigFilePath(ctx)
	params.LoadConfig(configFile)
	params.InitConfig()

	serverURL := params.GetConfig().Blockchain.GatewayURL
	capi = DialServer(serverURL)
	return capi
}

// DialServer connect to serverURL
func DialServer(serverURL string) *callapi.APICaller {
	capi := callapi.NewDefaultAPICaller()
	for {
		err := capi.DialServer(serverURL)
		if err == nil {
			break
		}
		time.Sleep(3 * time.Second)
	}
	return capi
}

