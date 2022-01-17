package worker

import (
	"sync"

	"github.com/fsn-dev/fsn-go-sdk/efsn/common"
	"github.com/fsn-dev/fsn-go-sdk/efsn/core/types"
)

const (
	maxParseBlocks = 1000
)

// Parse parse block and receipts
func (w *worker) Parse(block *types.Block, receipts types.Receipts) {
	msg := &message{
		block:    block,
		receipts: receipts,
	}
	w.messageChan <- msg
}

func (w *worker) startParser(wg *sync.WaitGroup) {
	defer wg.Done()
	count := 0
	wg2 := new(sync.WaitGroup)
	defer wg2.Wait()
	for {
		msg := <-w.messageChan
		if msg == nil {
			return
		}
		count++
		if !onlySyncAccount {
			wg2.Add(1)
			// parse block
			//go w.parseBlock(msg.block, wg2)
		}
		wg2.Add(1)
		// parse transactions
		go w.parseTransactions(msg.block, msg.receipts, wg2)
		if count == maxParseBlocks {
			count = 0
			wg2.Wait() // prevent memory exhausted (when blocks too large)
		}
	}
}

func (w *worker) parseTransactions(block *types.Block, receipts types.Receipts, wg *sync.WaitGroup) {
	defer wg.Done()
	wg.Add(len(block.Transactions()))
	for i, tx := range block.Transactions() {
		go w.parseTx(i, tx, block, receipts, wg)
	}
}

func (w *worker) parseTx(i int, tx *types.Transaction, block *types.Block, receipts types.Receipts, wg *sync.WaitGroup) {
	defer wg.Done()
	receipt := receipts[i]

	if onlySyncAccount {
		//parseReceipt(mt, receipt)
		return
	}

	if receipt != nil && len(receipt.Logs) != 0 {
		//parseReceipt(tx, receipt)
	}
}

func getTxSender(tx *types.Transaction) common.Address {
	signer := types.NewEIP155Signer(tx.ChainId())
	sender, _ := types.Sender(signer, tx)
	return sender
}
