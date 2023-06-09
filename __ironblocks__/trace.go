package __ironblocks__

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
)

type TraceTransactionResult map[string]interface{}

func TraceTransaction(ctx context.Context, hash common.Hash) error {
	if rpcClient == nil {
		logger.Warn("Ironblocks detector is not initialized,skipping")
		return nil
	}
	result := TraceTransactionResult{}
	err := rpcClient.do(ctx, &result, "debug_traceTransaction", hash, nil)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to trace transaction %s", hash.Hex()), "err", err)
		return nil

	}
	logger.Info(fmt.Sprintf("transaction %s trace result", hash.Hex()), "result", result)
	return nil
}
