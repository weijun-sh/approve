package utils

import (
	"github.com/weijun-sh/approve/log"
	"github.com/urfave/cli/v2"
)

var (
	// ConfigFileFlag -c|--config
	ConfigFileFlag = &cli.StringFlag{
		Name:    "config",
		Aliases: []string{"c"},
		Usage:   "config file, use toml format",
	}
	// ActionFlag -a|--action
	ActionFlag = &cli.StringFlag{
		Name:    "action",
		Aliases: []string{"a"},
		Usage:   "action, support action",
	}
	// LogFileFlag --log
	LogFileFlag = &cli.StringFlag{
		Name:  "log",
		Usage: "log file, support rotate",
	}
	// LogRotationFlag --rotate
	LogRotationFlag = &cli.Uint64Flag{
		Name:  "rotate",
		Usage: "log rotation time (unit hour)",
		Value: 24,
	}
	// LogMaxAgeFlag --maxage
	LogMaxAgeFlag = &cli.Uint64Flag{
		Name:  "maxage",
		Usage: "log max age (unit hour)",
		Value: 720,
	}
	// VerbosityFlag -v|--verbosity
	VerbosityFlag = &cli.Uint64Flag{
		Name:    "verbosity",
		Aliases: []string{"v"},
		Usage:   "log verbosity (0:panic, 1:fatal, 2:error, 3:warn, 4:info, 5:debug, 6:trace)",
		Value:   4,
	}
	// JSONFormatFlag --json
	JSONFormatFlag = &cli.BoolFlag{
		Name:  "json",
		Usage: "output log in json format",
	}
	// ColorFormatFlag --color
	ColorFormatFlag = &cli.BoolFlag{
		Name:  "color",
		Usage: "output log in color text format",
		Value: true,
	}

	// SyncFromFlag --syncfrom
	SyncFromFlag = &cli.Uint64Flag{
		Name:  "syncfrom",
		Usage: "sync start height, 0 means read from database",
		Value: 0,
	}
	// SyncToFlag --syncto
	SyncToFlag = &cli.Uint64Flag{
		Name:  "syncto",
		Usage: "sync end height (excluding end), 0 means endless",
		Value: 0,
	}

	// KeyStoreFileFlag --keystore
	KeyStoreFileFlag = &cli.StringFlag{
		Name:  "keystore",
		Usage: "keystore file path",
	}
	// PasswordFileFlag --password
	PasswordFileFlag = &cli.StringFlag{
		Name:  "password",
		Usage: "password file path",
	}
	// AmountFlag --amount
	AmountFlag = &cli.StringFlag{
		Name:  "amount",
		Usage: "amount",
	}
	// GasLimitFlag --gas
	GasLimitFlag = &cli.StringFlag{
		Name:  "gasLimit",
		Usage: "gas limit in transaction, use default if not specified",
	}
	// GasPriceFlag --gasPrice
	GasPriceFlag = &cli.StringFlag{
		Name:  "gasPrice",
		Usage: "gas price in transaction, use default if not specified",
	}
	// AccountNonceFlag --nonce
	AccountNonceFlag = &cli.StringFlag{
		Name:  "nonce",
		Usage: "nonce in transaction, use default if not specified",
	}

	// PairsFlag --pairs
	PairFlag = &cli.StringFlag{
		Name:  "pair",
		Usage: "pairs name",
	}
	// TokenFlag --token
	TokenFlag = &cli.StringFlag{
		Name:  "token",
		Usage: "token address",
	}
	// ExchangeFlag --exchange
	ExchangeFlag = &cli.StringFlag{
		Name:  "exchange",
		Usage: "exchange address",
	}
	// GatewayFlag --gateway
	GatewayFlag = &cli.StringFlag{
		Name:  "gateway",
		Usage: "gateway URL address",
	}
	// UseTimeMeasurementFlag --usetime
	UseTimeMeasurementFlag = &cli.BoolFlag{
		Name:  "usetime",
		Usage: "use timestamp instead of block height",
	}
)

// SyncArguments command line arguments
type SyncArguments struct {
	SyncStartHeight *uint64
	SyncEndHeight   *uint64
	SyncOverwrite   *bool
}

// SyncArgs sync arguments
var SyncArgs SyncArguments

// SetLogger set log level, json format, color, rotate ...
func SetLogger(ctx *cli.Context) {
	logLevel := VerbosityFlag.Value
	jsonFormat := JSONFormatFlag.Value
	colorFormat := ColorFormatFlag.Value
	log.SetLogger(uint32(logLevel), jsonFormat, colorFormat)

	logFile := ctx.String(LogFileFlag.Name)
	if logFile != "" {
		logRotation := ctx.Uint64(LogRotationFlag.Name)
		logMaxAge := ctx.Uint64(LogMaxAgeFlag.Name)
		log.SetLogFile(logFile, logRotation, logMaxAge)
	}
}

// GetConfigFilePath specified by `-c|--config`
func GetConfigFilePath(ctx *cli.Context) string {
	return ctx.String(ConfigFileFlag.Name)
}

// KeyStoreFile specified by `-k|--keystore`
func GetKeyStoreFile(ctx *cli.Context) string {
	return ctx.String(KeyStoreFileFlag.Name)
}

// PasswordFile specified by `-p|--password`
func GetPasswordFile(ctx *cli.Context) string {
	return ctx.String(PasswordFileFlag.Name)
}

// Action specified by `-a|--action`
func GetAction(ctx *cli.Context) string {
	return ctx.String(ActionFlag.Name)
}

// Pair specified by `--pair`
func GetPair(ctx *cli.Context) string {
	return ctx.String(PairFlag.Name)
}

// Amount specified by `--amount`
func GetAmount(ctx *cli.Context) float64 {
	return ctx.Float64(AmountFlag.Name)
}

// InitSyncArguments init sync arguments
func InitSyncArguments(ctx *cli.Context) {
	if ctx.IsSet(SyncFromFlag.Name) {
		start := ctx.Uint64(SyncFromFlag.Name)
		SyncArgs.SyncStartHeight = &start
	}
	if ctx.IsSet(SyncToFlag.Name) {
		end := ctx.Uint64(SyncToFlag.Name)
		SyncArgs.SyncEndHeight = &end
	}
}
