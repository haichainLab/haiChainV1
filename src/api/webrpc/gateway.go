package webrpc

import (
	"github.com/samoslab/haicoin/src/cipher"
	"github.com/samoslab/haicoin/src/coin"
	"github.com/samoslab/haicoin/src/daemon"
	"github.com/samoslab/haicoin/src/visor"
	"github.com/samoslab/haicoin/src/visor/historydb"
)

//go:generate goautomock -template=testify Gatewayer

// Gatewayer provides interfaces for getting skycoin related info.
type Gatewayer interface {
	GetLastBlocks(num uint64) (*visor.ReadableBlocks, error)
	GetBlocks(start, end uint64) (*visor.ReadableBlocks, error)
	GetBlocksInDepth(vs []uint64) (*visor.ReadableBlocks, error)
	GetUnspentOutputs(filters ...daemon.OutputsFilter) (*visor.ReadableOutputSet, error)
	GetTransaction(txid cipher.SHA256) (*visor.Transaction, error)
	InjectBroadcastTransaction(tx coin.Transaction) error
	GetAddrUxOuts(addr []cipher.Address) ([]*historydb.UxOut, error)
	GetTimeNow() uint64
}
