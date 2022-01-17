package params

import (
	"errors"
	"math/big"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/weijun-sh/approve/log"
	"github.com/fsn-dev/fsn-go-sdk/efsn/common"
)

const (
	defaultBlockTime uint64 = 13
)

var (
	config = &Config{}

	DecimalToken map[string]*big.Float = make(map[string]*big.Float)
	ServerURL string
	chain string
	cToken string
	cSpender []string
)

func InitConfig() {
        config := GetConfig()
        ServerURL = config.Blockchain.GatewayURL
	if err := CheckConfig(); err != nil {
		log.Fatalf("Check config failed. %v", err)
	}
	chain = config.Blockchain.Chain
	cToken = config.Contract.Token
	cSpender = config.Contract.Spender
}

func GetContracts() (string, []string) {
	return cToken, cSpender
}

func GetDecimal(token string) (*big.Float, error) {
	token = strings.ToLower(token)
	if DecimalToken[token] == nil {
		log.Info("GetDecimal", "token", token, "decimail", "nil")
		return nil, errors.New("decimail nil")
	}
	return DecimalToken[token], nil
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

// Config config
type Config struct {
        Contract *ContractConfig
        Account *AccountConfig
        Blockchain *blockchain_struct
}

type ContractConfig struct {
        Token string
        Spender []string
}

type AccountConfig struct {
        Address []string
}

type blockchain_struct struct {
	Chain string
        GatewayURL string
}

