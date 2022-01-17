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
	contractAddr_token :=  params.Contracts[pairUpper].Token
	tokenExchangeAddr_token := params.Contracts[pairUpper].Exchange

	initDecimal(contractAddr_token)

	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	fmt.Printf("ps. update every 10 minutes\n")
	fmt.Printf("=======================================================================\n")
	fmt.Printf("                        ETH APPROVE and BALANCE\n")
	fmt.Printf("=======================================================================\n")
	fmt.Printf("     | owner                                       allowance    balance\n")
	fmt.Printf("-----+-----------------------------------------------------------------\n")
        //l, b := Allowance("addresses.txt", contractAddr_token, tokenExchangeAddr_token)
        allowance(contractAddr_token, tokenExchangeAddr_token)
	//fmt.Printf(" %3v                                                          %10v\n", l, b)
}

func allowance(contractAddr, exchangeAddr string) (int, string){
	return transactions.Allowance(contractAddr, exchangeAddr)
}
