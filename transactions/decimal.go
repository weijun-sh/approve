package transactions

import (
	"fmt"
	"math"
	"math/big"
	"strings"
	"os"

	"github.com/weijun-sh/approve/client"
	"github.com/weijun-sh/approve/log"
	"github.com/weijun-sh/approve/params"
	"github.com/weijun-sh/approve/tools"
	"github.com/fsn-dev/fsn-go-sdk/efsn/common"
)

func InitTokenDecimal(token string) {
	token = strings.ToLower(token)
	decimal, err := getDecimal(token)
	if err != nil {
		log.Warn("InitTokenDecimal", "token", token, "decimal", "nil")
		os.Exit(1)
	}
	d, _ := decimal.Float64()
	//fmt.Printf("InitTokenDecimal, d: %v, decimal: %v, token: %v\n", d, decimal, token)
	params.DecimalToken[token] = big.NewFloat(math.Pow(10, d))
}

func getDecimal(contractAddr string) (*big.Float, error) {
	token, err := GetTokenOfContract(contractAddr)
	if err != nil {
		fmt.Printf("GetTokenOfContract failed, err=%v\n", err)
		return nil, err
	}
	decimal, err := token.Decimals(nil)
	if err != nil {
		fmt.Printf("Decimals failed, err=%v\n", err)
		return nil, err
	}
	return tools.BigIntToFloat(decimal), nil
}

func GetTokenOfExchange(exchangeAddr string) (*Token, error) {
	var err error
	if newToken[exchangeAddr] == nil {
		ethClient := client.GetClient()
		newToken[exchangeAddr], err = NewToken(common.HexToAddress(exchangeAddr), ethClient)
		if err != nil {
			fmt.Printf("err: %v, NewToken(address: %v, client: %v)\n", err, exchangeAddr, ethClient)
			return nil, err
		}
	}
	return newToken[exchangeAddr], nil
}

func GetTokenOfContract(contractAddr string) (*Token, error) {
	var err error
	if newToken_contract[contractAddr] == nil {
		ethClient := client.GetClient()
		newToken_contract[contractAddr], err = NewToken(common.HexToAddress(contractAddr), ethClient)
		if err != nil {
			fmt.Printf("err: %v, NewToken(address: %v, client: %v)\n", err, contractAddr, ethClient)
			return nil, err
		}
	}
	return newToken_contract[contractAddr], nil
}

