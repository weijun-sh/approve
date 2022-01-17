package params

import (
	"errors"
	"fmt"
	"math/big"
	"os"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/weijun-sh/approve/log"
	"github.com/weijun-sh/approve/tools"
	"github.com/fsn-dev/fsn-go-sdk/efsn/common"
)

const (
	defaultBlockTime uint64 = 13
)

var (
	config = &Config{}

	linkToken = []string{"FSN", "ETH", "BNB", "HT", "FTM"}
	DecimalToken map[string]*big.Float = make(map[string]*big.Float)
	ServerURL string
	CoinTypeCoin string
	CoinTypeToken string
	TokenType = make(map[string]string)
	Pair string
	factoryAddresses []common.Address
	Contracts = make(map[string]*ContractConfig)
	AverageBlockTime uint64
	RewardToken string
        StrategySycle string
	Slipage float64
	DeadLine uint64
	Amount float64
	Fee float64

	GasLimit uint64
	GasPrice uint64

	AllExchanges = make(map[common.Address]struct{}) //temp var
	AllTokens    = make(map[common.Address]struct{}) //temp var
)

func InitConfig() {
        config := GetConfig()
        ServerURL = config.Blockchain.GatewayURL
	CoinTypeCoin, CoinTypeToken = parsePairs(config.Pair)
	Pair = fmt.Sprintf("%v/%v", CoinTypeToken, CoinTypeCoin)
	AverageBlockTime = config.Blocknumber.AverageBlockTime
	RewardToken = config.Mine.RewardToken
	StrategySycle = config.Mine.StrategySycle
	Slipage = config.Slipage.Slipage
	DeadLine = config.Blockchain.DeadLine
	Fee = config.Fee.Fee
	if err := CheckConfig(); err != nil {
		log.Fatalf("Check config failed. %v", err)
	}
	parseContracts()
	CheckPairContract()
}

func GetExchange(pair string) (string, []string, error) {
	if Contracts[pair] != nil {
		return Contracts[pair].Token, Contracts[pair].Exchange, nil
	}
	return "", nil, errors.New("GetExchange nil")
}

func parseContracts() {
	for _, ex := range config.Contracts {
		c, t := parsePairs(ex.Pair)
		pair := strings.ToUpper(fmt.Sprintf("%v/%v", t, c))
		Contracts[pair] = &ContractConfig{}
		Contracts[pair].Pair = ex.Pair
		Contracts[pair].Token = ex.Token
		Contracts[pair].Exchange = ex.Exchange
		//TokenType[strings.ToLower(ex.Token)] = t
		//TokenType[strings.ToLower(ex.Exchange)] = t
	}
}

func CheckPairContract() {
	pair := fmt.Sprintf("%v/%v", CoinTypeToken, CoinTypeCoin)
	if Contracts[pair] == nil {
		log.Info("checkPairContract", "pair", pair, "not exist", "")
		os.Exit(1)
	}
}

func parsePairs(pair string) (string, string) {
	pairLower := strings.ToUpper(pair)
        pairSlice := strings.Split(pairLower, "/")
        coinType := pairSlice[0]
        tokenType := pairSlice[1]
	if len(coinType) > 0 && len(tokenType) > 0 {
		for _, token := range linkToken {
			if coinType == token {
				return coinType, tokenType
			}
			if tokenType == token {
				return tokenType, coinType
			}
		}
	}
        os.Exit(1)
	return "", ""
}

func GetDecimal(token string) (*big.Float, error) {
	token = strings.ToLower(token)
	if DecimalToken[token] == nil {
		log.Info("GetDecimal", "token", token, "decimail", "nil")
		return nil, errors.New("decimail nil")
	}
	return DecimalToken[token], nil
}

func GetTokenType(token string) (string, error) {
	token = strings.ToLower(token)
	if TokenType[token] == "" {
		log.Info("GetTokenType", "token", token, "TokenType", "")
		return "", errors.New("TokenType nil")
	}
	return TokenType[token], nil
}

// GetAverageBlockTime average block time
func GetAverageBlockTime() uint64 {
	avg := config.Blocknumber.AverageBlockTime
	if avg == 0 {
		avg = defaultBlockTime
	}
	return avg
}

// IsScanAllExchange is scan all exchange
func IsScanAllExchange() bool {
	return config.Sync.ScanAllExchange
}

// IsRecordTokenAccount is record token account
func IsRecordTokenAccount() bool {
	return config.Sync.RecordTokenAccount
}

// IsConfigedExchange return true if exchange is configed
//func IsConfigedExchange(exchange string) bool {
//	return GetExchangePairs(exchange) != ""
//}

// IsConfigedToken return true if token is configed
//func IsConfigedToken(token string) bool {
//	return GetConfigedExchange(token) != ""
//}

// GetConfigedExchange get configed exchange
func GetConfigedExchange(token string) []string {
	for _, ex := range config.Contracts {
		if strings.EqualFold(ex.Token, token) {
			return ex.Exchange
		}
	}
	return nil
}

// GetExchangePairs get pairs from config
func GetExchangePairs(exchange string) string {
//	for _, ex := range config.Contracts {
//		if strings.EqualFold(ex.Exchange, exchange) {
//			return ex.Pair
//		}
//	}
	return ""
}

// GetExchangeToken get exchane token from config
//func GetExchangeToken(exchange string) string {
//	log.Info("GetExchangeToken")
//	for _, ex := range config.Contracts {
//		if strings.EqualFold(ex.Exchange, exchange) {
//			return ex.Token
//		}
//	}
//	return ""
//}

// GetTokenAddress get token address from config
//func GetTokenAddress(exchange string) string {
//	log.Info("GetTokenAddress")
//	for _, ex := range config.Contracts {
//		if strings.EqualFold(ex.Exchange, exchange) {
//			return ex.Token
//		}
//	}
//	return ""
//}

func GetTransferTo() []string {
	return config.Transfer.To
}

// GetFactories get facotries
func GetFactories() []common.Address {
	if factoryAddresses == nil {
		factories := make([]common.Address, len(config.Factories))
		for i, factory := range config.Factories {
			factories[i] = common.HexToAddress(factory)
		}
		log.Info("GetFactories", "factories", tools.ToJSONString(factories, !log.JSONFormat))
		factoryAddresses = factories
	}
	return factoryAddresses
}

// IsConfigedFactory is configed factory
func IsConfigedFactory(factory common.Address) bool {
	for _, fact := range GetFactories() {
		if factory == fact {
			return true
		}
	}
	return false
}

// GetAmount get config items structure
func GetAmount() float64 {
	return Amount
}

// GetConfig get config items structure
func GetConfig() *Config {
	return config
}

func GetAccount() *AccountConfig {
	return config.Account
}

// SetConfig set config items
func SetConfig(cfg *Config) {
	config = cfg
}

func UpdateConfig(keystore, password, action, pair string, value float64) {
	if action != "" {
		config.Action = action
	}
	if pair != "" {
		c, t := parsePairs(pair)
		config.Pair = fmt.Sprintf("%v/%v", t, c)
	}
	Amount = value
	GasLimit = config.Blockchain.GasLimit
	GasPrice = config.Blockchain.GasPrice
}

// LoadConfig load config
func LoadConfig(configFile string) *Config {
	//log.Info("LoadConfig", "configFile", configFile)
	if !common.FileExist(configFile) {
		log.Fatalf("LoadConfig error: config file '%v' not exist", configFile)
	}

	tmpConfig := &Config{}
	if _, err := toml.DecodeFile(configFile, &tmpConfig); err != nil {
		log.Fatalf("LoadConfig error (toml DecodeFile): %v", err)
	}

	SetConfig(tmpConfig)

	return tmpConfig
}

// AddTokenAndExchange add token and exchange
func AddTokenAndExchange(token, exchange common.Address) {
	if token == (common.Address{}) || exchange == (common.Address{}) {
		return
	}
	AllTokens[token] = struct{}{}
	AllExchanges[exchange] = struct{}{}
}

// IsInAllTokens is exchange token
func IsInAllTokens(token common.Address) bool {
	_, exist := AllTokens[token]
	return exist
}

// IsInAllExchanges is in all exchanges
func IsInAllExchanges(exchange common.Address) bool {
	_, exist := AllExchanges[exchange]
	return exist
}

// IsInAllTokenAndExchanges is in all exchanges or tokens
func IsInAllTokenAndExchanges(address common.Address) bool {
	return IsInAllTokens(address) || IsInAllExchanges(address)
}

// IsExcludedRewardAccount is excluded
//func IsExcludedRewardAccount(account common.Address) bool {
//	if account == (common.Address{}) {
//		return true
//	}
//	if IsConfigedExchange(account.String()) {
//		return true
//	}
//	return IsInAllExchanges(account)
//}

// Config config
type Config struct {
        Action string
	Pair string
	Factories []string
        Mine *mineConfig
        Exchange *exchange_struct
        Transfer *transferConfig
        Contracts []ContractConfig
        Decimal *decimal_struct
        Account *AccountConfig
        Blockchain *blockchain_struct
        Sync *SyncConfig
        Blocknumber *BlockConfig
        Blocktime *BlockConfig
        Slipage *slipageConfig
        Fee *fee_struct
}

type action_struct struct {
        Action string
        Pair string
}

type transferConfig struct {
	To []string
}

type mineConfig struct {
        TxMinAmount float64
        TxMinExpectPercent float64
        Reward int64
        RewardToken string
        SellRewardToken bool
        TxRate float64
        TxRewardPercent float64
        Rebase string
        StrategySycle string
        StrategyReward string
        Time_sycle int64
        Blocks_mine []int64
}

type BlockConfig struct {
        Blocks_day uint64
        Blocks_sycle uint64
        Block_start uint64
        AverageBlockTime uint64
}

type exchange_struct struct {
        Amount_buy uint64
        Amount_sell uint64
        Times_mine_sycle uint64
        Height uint64
        Sycles uint64
	Rounds uint64
}

type ContractConfig struct {
	Pair string
        Token string
        Exchange []string
}

type decimal_struct struct {
        Decimal_ANY int64
        Decimal_USDT int64
        Decimal_ETH int64
        Decimal_BTC int64
        Decimal_CYC int64
}

type AccountConfig struct {
        Address []string
}

type blockchain_struct struct {
        GatewayURL string
        GasLimit uint64
        GasPrice uint64
        DeadLine uint64
}

type SyncConfig struct {
        BlockStart uint64
        BlockEnd uint64
	ScanAllExchange bool
	RecordTokenAccount bool
}

type slipageConfig struct {
        Slipage float64
}

type fee_struct struct {
        Fee float64
}

