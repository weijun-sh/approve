package utils

import (
	"fmt"
	"strings"
	"time"

	"github.com/weijun-sh/approve/callapi"
	"github.com/weijun-sh/approve/log"
	"github.com/weijun-sh/approve/params"
	"github.com/fsn-dev/fsn-go-sdk/efsn/common"
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
	keystore := GetKeyStoreFile(ctx)
	password := GetPasswordFile(ctx)
	action := GetAction(ctx)
	pair := GetPair(ctx)
	amount := GetAmount(ctx)
	params.UpdateConfig(keystore, password, action, pair, amount)
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

func initAllExchanges() {
	for _, factory := range params.GetFactories() {
		initExchangesInFactory(factory)
	}
}

func CheckContract() error {
	initAllExchanges()
	err := verifyConfig()
	return err
}

func initExchangesInFactory(factory common.Address) {
	log.Info("initExchangesInFactory", "factory", factory.String())
	tokenCount := capi.LoopGetFactoryTokenCount(factory)
	for i := uint64(1); i <= tokenCount; i++ {
		token := capi.LoopGetFactoryTokenWithID(factory, i)
		exchange := capi.LoopGetFactoryExchange(factory, token)
		params.AddTokenAndExchange(token, exchange)
	}
	log.Info("initExchangesInFactory success", "factory", factory.String(), "tokenCount", tokenCount, "added", len(params.AllExchanges))
}

func verifyConfig() error {
	err := checkContracts()
	if err != nil {
		return err
	}
	err = checkAction()
	if err != nil {
		return err
	}
	return nil
}

func checkAction() error {
	config := params.GetConfig()
	if strings.ToLower(config.Action) == "exchange" {
		if config.Exchange.Sycles <= 0 {
			return fmt.Errorf("exchange Sycles want > 0")
		}
		if config.Exchange.Rounds <= 0 {
			return fmt.Errorf("exchange Round want > 0")
		}
	}
	return nil
}

func checkContracts() error {
	config := params.GetConfig()
	for _, ex := range config.Contracts {
		exchange := common.HexToAddress(ex.Exchange)
		token := common.HexToAddress(ex.Token)
		wantToken := capi.GetExchangeTokenAddress(exchange)
		if token != wantToken {
			return fmt.Errorf("exchange token mismatch. exchange %v want token %v, but have %v", ex.Exchange, wantToken.String(), ex.Token)
		}
		factory := capi.GetExchangeFactoryAddress(exchange)
		if !params.IsConfigedFactory(factory) {
			return fmt.Errorf("exchange %v 's factory %v is not configed", ex.Exchange, factory.String())
		}
		log.Info("verify exchange token success", "exchange", ex.Exchange, "token", ex.Token, "factory", factory.String())
	}
	return nil
}

