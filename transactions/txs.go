package transactions

import (
	"fmt"
	"math/big"
	"strings"
	//"time"

	"github.com/weijun-sh/approve/client"
	"github.com/weijun-sh/approve/params"
	"github.com/fsn-dev/fsn-go-sdk/efsn/accounts/abi/bind"
	"github.com/fsn-dev/fsn-go-sdk/efsn/common"
)

var (
	newToken map[string]*Token = make(map[string]*Token)
	newToken_contract map[string]*Token = make(map[string]*Token)
	rebaseToken *Cyctoken
	sendTxsNow bool = true
)

func Allowance(contractAddr, exchangeAddr string) (int, string) {
        addresses := params.GetAccount()
	total := new(big.Float).SetFloat64(0)
	totalNotNil := new(big.Float).SetFloat64(0)
	var nn int = 0
	for i, kf := range addresses.Address {
		a, b, _ := callAllowance(i, kf, contractAddr, exchangeAddr)
		if a.Cmp(new(big.Float).SetFloat64(0)) > 0 {
			totalNotNil = new(big.Float).Add(totalNotNil, b)
			nn += 1
		}
		total = new(big.Float).Add(total, b)
	}
	bAString := fmt.Sprintf("%0.4f", total)
	bString := fmt.Sprintf("%0.4f", totalNotNil)
	fmt.Printf("-----------------------------------------------------------------------------\n")
	//fmt.Printf("\n summary\n")
	//fmt.Printf("================================\n")
	//fmt.Printf("           | amount    balance\n")
	//fmt.Printf("-----------+--------------------\n")
	fmt.Printf("                                                total  %9v %10v\n", len(addresses.Address), bAString)
	fmt.Printf("                                            allowance  %9v %10v\n", nn, bString)
	//fmt.Printf("================================\n")
	fmt.Printf("=============================================================================\n")
	return len(addresses.Address), bString
}

func callAllowance(i int, keyAddress, contractAddr, exchangeAddr string) (*big.Float, *big.Float, error) {
	//return allowance(i, keyAddress, contractAddr, exchangeAddr) //tx
	return allowance2(i, keyAddress, contractAddr, exchangeAddr) //tx
}

func allowance2(i int, keyAddress, contractAddr, exchangeAddr string) (*big.Float, *big.Float, error) {
	contractAddr = strings.ToLower(contractAddr)
	token_contract, err := getTokenOfContract(contractAddr)
	if err != nil {
		return nil, nil, err
	}

	from := common.HexToAddress(keyAddress)

	allownceRet, err := token_contract.Allowance(&bind.CallOpts{Pending: false}, from, common.HexToAddress(exchangeAddr))
	if err != nil {
		return nil, nil, err
	}
	ret := new(big.Float).Quo(new(big.Float).SetInt(allownceRet), params.DecimalToken[contractAddr])
	balance_token, errb := token_contract.BalanceOf(&bind.CallOpts{Pending: false}, common.HexToAddress(keyAddress))
	if errb != nil {
		return nil, nil, err
	}
	convertBalance := new(big.Float).Quo(new(big.Float).SetInt(balance_token), params.DecimalToken[contractAddr])
	aString := fmt.Sprintf("%10v", ret)
	if ret.Cmp(big.NewFloat(100000000)) >= 0 {
		aString = "        -1"
	}
	bString := fmt.Sprintf("%0.4f", convertBalance)
	fmt.Printf(" %7v | %v %v %10v\n", i+1, from.String(), aString, bString)
	return ret, convertBalance, nil
}

func getTokenOfContract(contractAddr string) (*Token, error) {
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
