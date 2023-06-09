package __ironblocks__

import (
	"github.com/ethereum/go-ethereum/log"
)

var rpcClient *client
var logger log.Logger

func InitIronblocks(cfg *Config) {
	rpcClient = newClient()
	if err := rpcClient.init(cfg); err != nil {
		logger.Warn("failed to init ironblocks client", "err", err)
		return
	}
	logger.Info("Ironblocks client initialized, will use rpc endpoint", "url", cfg.getRPCURL())
}

func init() {
	logger = log.New()
}
