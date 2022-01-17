package params

import (
	"errors"
	"fmt"
	//"strings"

	"github.com/fsn-dev/fsn-go-sdk/efsn/common"
)

// CheckConfig check config
func CheckConfig() (err error) {
	switch {
	case config == nil:
		return errors.New("empty config")
	case config.Blockchain == nil:
		return errors.New("must config Blockchain")
	case config.Contracts == nil:
		return errors.New("must config Contracs")
	}
	//err = checkExchangeConfig()
	//if err != nil {
	//	return err
	//}
	return nil
}

func checkExchangeConfig() error {
//	pairsMap := make(map[string]struct{})
//	exchangeMap := make(map[string]struct{})
//	tokenMap := make(map[string]struct{})
//	for _, ex := range config.Contracts {
//		if err := ex.check(); err != nil {
//			return err
//		}
//		pairs := strings.ToLower(ex.Pair)
//		exchange := strings.ToLower(ex.Exchange)
//		token := strings.ToLower(ex.Token)
//		if _, exist := pairsMap[pairs]; exist {
//			return fmt.Errorf("duplicate pairs %v", ex.Pair)
//		}
//		if _, exist := exchangeMap[exchange]; exist {
//			return fmt.Errorf("duplicate exchange %v", ex.Exchange)
//		}
//		if _, exist := tokenMap[token]; exist {
//			return fmt.Errorf("duplicate token %v", ex.Token)
//		}
//		pairsMap[pairs] = struct{}{}
//		exchangeMap[exchange] = struct{}{}
//		tokenMap[token] = struct{}{}
//	}
	return nil
}

func (ex *ContractConfig) check() error {
	//if !common.IsHexAddress(ex.Exchange) {
	//	return fmt.Errorf("[check exchange] wrong exchange address '%v'", ex.Exchange)
	//}
	//if ex.Pair == "" {
	//	return fmt.Errorf("[check exchange] empty exchange pairs (exchange %v)", ex.Exchange)
	//}
	if !common.IsHexAddress(ex.Token) {
		return fmt.Errorf("[check exchange] wrong exchange token '%v' (exchange %v)", ex.Token, ex.Exchange)
	}
	return nil
}

