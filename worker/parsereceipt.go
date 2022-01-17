package worker

import (
	"time"

	"github.com/fsn-dev/fsn-go-sdk/efsn/common"
)

var (
	topicTokenPurchase   = common.HexToHash("0xcd60aa75dea3072fbc07ae6d7d856b5dc5f4eee88854f5b4abf7b680ef8bc50f")
	topicEthPurchase     = common.HexToHash("0x7f4091b46c33e918a0f3aa42307641d17bb67029427a5369e54b353984238705")
	topicAddLiquidity    = common.HexToHash("0x06239653922ac7bea6aa2b19dc486b9361821d37712eb796adfd38d81de278ca")
	topicRemoveLiquidity = common.HexToHash("0x0fbf06c058b90cb038a618f8c2acbf6145f8b3570fd1fa56abb8f0f3f05b36e8")
	topicTransfer        = common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")
	topicApproval        = common.HexToHash("0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925")
	topicCreateExchange  = common.HexToHash("0x9d42cb017eb05bd8944ab536a8b35bc68085931dd5f4356489801453923953f9")
)

const secondsPerDay = 24 * 3600

func getDayBegin(timestamp uint64) uint64 {
	return timestamp - timestamp%secondsPerDay
}

func timestampToDate(timestamp uint64) string {
	return time.Unix(int64(timestamp), 0).Format("2006-01-02 15:04:05")
}

