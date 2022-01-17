package worker

import (
	"fmt"
	"time"

	"github.com/weijun-sh/approve/transactions"
	"github.com/weijun-sh/approve/params"
)

func initDecimal(token string) {
	transactions.InitTokenDecimal(token)
}

func Start() {
	contractAddr_token, tokenExchangeAddr_token := params.GetContracts()

	initDecimal(contractAddr_token)

	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	//fmt.Printf("ps. update every 10 minutes\n")
	fmt.Printf("=============================================================================\n")
	fmt.Printf("                             ETH APPROVE and BALANCE\n")
	fmt.Printf("=============================================================================\n")
        //l, b := Allowance("addresses.txt", contractAddr_token, tokenExchangeAddr_token)
	for _, spender := range tokenExchangeAddr_token {
		fmt.Printf(" spender | %v  allowance    balance\n", spender)
		fmt.Printf("---------+-------------------------------------------------------------------\n")
		allowance(contractAddr_token, spender)
	}
	//fmt.Printf(" %3v                                                          %10v\n", l, b)
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
}

func allowance(contractAddr, exchangeAddr string) (int, string){
	return transactions.Allowance(contractAddr, exchangeAddr)
}
