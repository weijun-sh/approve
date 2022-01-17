package client

import (
	"context"
	"time"

	"github.com/weijun-sh/approve/log"
	"github.com/weijun-sh/approve/params"
	"github.com/fsn-dev/fsn-go-sdk/efsn/ethclient"
)

var (
	client     *ethclient.Client
	cliContext = context.Background()
)

func dialServer() (err error) {
	client, err = ethclient.Dial(params.ServerURL)
	if err != nil {
		log.Error("[syncer] client connection error", "server", params.ServerURL, "err", err)
		return err
	}
	log.Debug("[syncer] client connection succeed", "server", params.ServerURL)
	return nil
}

func CloseClient() {
	if client != nil {
		client.Close()
	}
}

func GetClient() *ethclient.Client {
	for {
		err := dialServer()
		if err == nil {
			break
		}
		time.Sleep(3 * time.Second)
	}
	return client
}

