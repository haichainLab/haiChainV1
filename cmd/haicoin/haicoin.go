package main

import (
	_ "net/http/pprof"

	"github.com/samoslab/haicoin/src/skycoin"
	"github.com/samoslab/haicoin/src/util/logging"
	"github.com/samoslab/haicoin/src/visor"
)

var (
	// Version of the node. Can be set by -ldflags
	Version = "1.24.1"
	// Commit ID. Can be set by -ldflags
	Commit = ""
	// Branch name. Can be set by -ldflags
	Branch = ""
	// ConfigMode (possible values are "", "STANDALONE_CLIENT").
	// This is used to change the default configuration.
	// Can be set by -ldflags
	ConfigMode = ""

	logger = logging.MustGetLogger("main")

	// GenesisSignatureStr hex string of genesis signature
	GenesisSignatureStr = "368c9041de99564368fab247fa301a9769852b0c6d06c9118ba1efc4a0f4c78025d4d3972c21e73ee2d052f274bec8e30157a959ad8af107305fbc9ccbcf199000"
	// GenesisAddressStr genesis address string
	GenesisAddressStr = "ybCiFspxqcmJg9jNn1L9RFctnG6HZqPSMf"
	// BlockchainPubkeyStr pubic key string
	BlockchainPubkeyStr = "03a83ac2fb482601b81f581c6b2cd82c4f6d6f10eda9ecfecf35d2126548f2ac75"
	// BlockchainSeckeyStr empty private key string
	BlockchainSeckeyStr = ""

	// GenesisTimestamp genesis block create unix time
	GenesisTimestamp uint64 = 1542082636
	// GenesisCoinVolume represents the coin capacity
	GenesisCoinVolume uint64 = 500000000000000

	// DefaultConnections the default trust node addresses
	DefaultConnections = []string{
		"193.112.246.101:6968",
		"150.109.62.225:6968",
		"150.109.62.225:6968",
		"47.52.211.167:6968",
	}
)

func main() {
	// get node config
	nodeConfig := skycoin.NewNodeConfig(ConfigMode, skycoin.NodeParameters{
		GenesisSignatureStr: GenesisSignatureStr,
		GenesisAddressStr:   GenesisAddressStr,
		GenesisCoinVolume:   GenesisCoinVolume,
		GenesisTimestamp:    GenesisTimestamp,
		BlockchainPubkeyStr: BlockchainPubkeyStr,
		BlockchainSeckeyStr: BlockchainSeckeyStr,
		DefaultConnections:  DefaultConnections,
		PeerListURL:         "http://haicoin.io/haicoin-peers.txt",
		Port:                6968,
		WebInterfacePort:    6969,
		DataDirectory:       "$HOME/.haicoin",
		ProfileCPUFile:      "haicoin.prof",
	})

	// create a new fiber coin instance
	coin := skycoin.NewCoin(
		skycoin.Config{
			Node: *nodeConfig,
			Build: visor.BuildInfo{
				Version: Version,
				Commit:  Commit,
				Branch:  Branch,
			},
		},
		logger,
	)

	// parse config values
	coin.ParseConfig()

	// run fiber coin node
	coin.Run()
}
