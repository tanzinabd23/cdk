package config

// DefaultValues is the default configuration
const DefaultValues = `
ForkUpgradeBatchNumber = 0
ForkUpgradeNewForkId = 0

[Etherman]
	URL="http://localhost:8545"
	ForkIDChunkSize=100
	[Etherman.EthermanConfig]
		URL="http://localhost:8545"
		MultiGasProvider=false
		L1ChainID=1337
		HTTPHeaders=[]
		[Etherman.EthermanConfig.Etherscan]
			ApiKey=""
			Url="https://api.etherscan.io/api?module=gastracker&action=gasoracle&apikey="

[Common]
NetworkID = 1
IsValidiumMode = false
ContractVersions = "banana"

[Log]
Environment = "development" # "production" or "development"
Level = "info"
Outputs = ["stderr"]

[SequenceSender]
WaitPeriodSendSequence = "15s"
LastBatchVirtualizationTimeMaxWaitPeriod = "10s"
L1BlockTimestampMargin = "30s"
MaxTxSizeForL1 = 131072
L2Coinbase = "0xfa3b44587990f97ba8b6ba7e230a5f0e95d14b3d"
PrivateKey = {Path = "./test/sequencer.keystore", Password = "testonly"}
SequencesTxFileName = "sequencesender.json"
GasOffset = 80000
WaitPeriodPurgeTxFile = "15m"
MaxPendingTx = 1
MaxBatchesForL1 = 300
BlockFinality = "FinalizedBlock"
RPCURL = "localhost:8123"
GetBatchWaitInterval = "10s"
	[SequenceSender.EthTxManager]
		FrequencyToMonitorTxs = "1s"
		WaitTxToBeMined = "2m"
		GetReceiptMaxTime = "250ms"
		GetReceiptWaitInterval = "1s"
		PrivateKeys = [
			{Path = "./test/sequencer.keystore", Password = "testonly"},
		]
		ForcedGas = 0
		GasPriceMarginFactor = 1
		MaxGasPriceLimit = 0
		StoragePath = "ethtxmanager.db"
		ReadPendingL1Txs = false
		SafeStatusL1NumberOfBlocks = 0
		FinalizedStatusL1NumberOfBlocks = 0
			[SequenceSender.EthTxManager.Etherman]
				URL = "http://127.0.0.1:8545"
				MultiGasProvider = false
				L1ChainID = 1337
[Aggregator]
Host = "0.0.0.0"
Port = 50081
RetryTime = "5s"
VerifyProofInterval = "10s"
TxProfitabilityCheckerType = "acceptall"
TxProfitabilityMinReward = "1.1"
ProofStatePollingInterval = "5s"
SenderAddress = ""
CleanupLockedProofsInterval = "2m"
GeneratingProofCleanupThreshold = "10m"
BatchProofSanityCheckEnabled = true
ForkId = 9
GasOffset = 0
WitnessURL = "localhost:8123"
UseL1BatchData = true
UseFullWitness = false
SettlementBackend = "l1"
AggLayerTxTimeout = "5m"
AggLayerURL = ""
MaxWitnessRetrievalWorkers = 2
SyncModeOnlyEnabled = false
SequencerPrivateKey = {}
	[Aggregator.DB]
		Name = "aggregator_db"
		User = "aggregator_user"
		Password = "aggregator_password"
		Host = "cdk-aggregator-db"
		Port = "5432"
		EnableLog = false	
		MaxConns = 200
	[Aggregator.Log]
		Environment = "development" # "production" or "development"
		Level = "info"
		Outputs = ["stderr"]
	[Aggregator.StreamClient]
		Server = "localhost:6900"
	[Aggregator.EthTxManager]
		FrequencyToMonitorTxs = "1s"
		WaitTxToBeMined = "2m"
		GetReceiptMaxTime = "250ms"
		GetReceiptWaitInterval = "1s"
		PrivateKeys = [
			{Path = "/pk/aggregator.keystore", Password = "testonly"},
		]
		ForcedGas = 0
		GasPriceMarginFactor = 1
		MaxGasPriceLimit = 0
		StoragePath = ""
		ReadPendingL1Txs = false
		SafeStatusL1NumberOfBlocks = 0
		FinalizedStatusL1NumberOfBlocks = 0
			[Aggregator.EthTxManager.Etherman]
				URL = ""
				L1ChainID = 11155111
				HTTPHeaders = []
	[Aggregator.Synchronizer]
		[Aggregator.Synchronizer.Log]
			Environment = "development" # "production" or "development"
			Level = "info"
			Outputs = ["stderr"]
		[Aggregator.Synchronizer.SQLDB]
			DriverName = "sqlite3"
			DataSourceName = "file:/tmp/aggregator_sync_db.sqlite"
		[Aggregator.Synchronizer.Synchronizer]
			SyncInterval = "10s"
			SyncChunkSize = 1000
			GenesisBlockNumber = 5511080
			SyncUpToBlock = "finalized"
			BlockFinality = "finalized"
			OverrideStorageCheck = false
		[Aggregator.Synchronizer.Etherman]
			L1URL = "http://localhost:8545"
			ForkIDChunkSize = 100
			L1ChainID = 0
			[Aggregator.Synchronizer.Etherman.Validium]
				Enabled = false
				TrustedSequencerURL = ""
				RetryOnDACErrorInterval = "1m"
				DataSourcePriority = ["trusted", "external"]
			[Aggregator.Synchronizer.Etherman.Validium.Translator]
				FullMatchRules = []
			[Aggregator.Synchronizer.Etherman.Validium.RateLimit]
				NumRequests = 900
				Interval = "1s"
[ReorgDetectorL1]
DBPath = "/tmp/reorgdetectorl1"

[ReorgDetectorL2]
DBPath = "/tmp/reorgdetectorl2"

[L1InfoTreeSync]
DBPath = "/tmp/L1InfoTreeSync.sqlite"
GlobalExitRootAddr="0x8464135c8F25Da09e49BC8782676a84730C318bC"
SyncBlockChunkSize=10
BlockFinality="LatestBlock"
URLRPCL1="http://test-aggoracle-l1:8545"
WaitForNewBlocksPeriod="100ms"
InitialBlock=0

[AggOracle]
TargetChainType="EVM"
URLRPCL1="http://test-aggoracle-l1:8545"
BlockFinality="FinalizedBlock"
WaitPeriodNextGER="100ms"
	[AggOracle.EVMSender]
		GlobalExitRootL2="0x8464135c8F25Da09e49BC8782676a84730C318bC"
		URLRPCL2="http://test-aggoracle-l2:8545"
		ChainIDL2=1337
		GasOffset=0
		WaitPeriodMonitorTx="100ms"
		SenderAddr="0x70997970c51812dc3a010c7d01b50e0d17dc79c8"
		[AggOracle.EVMSender.EthTxManager]
				FrequencyToMonitorTxs = "1s"
				WaitTxToBeMined = "2s"
				GetReceiptMaxTime = "250ms"
				GetReceiptWaitInterval = "1s"
				PrivateKeys = [
					{Path = "/app/keystore/aggoracle.keystore", Password = "testonly"},
				]
				ForcedGas = 0
				GasPriceMarginFactor = 1
				MaxGasPriceLimit = 0
				StoragePath = "/tmp/ethtxmanager-sequencesender.db"
				ReadPendingL1Txs = false
				SafeStatusL1NumberOfBlocks = 5
				FinalizedStatusL1NumberOfBlocks = 10
					[AggOracle.EVMSender.EthTxManager.Etherman]
						URL = "http://test-aggoracle-l2"
						MultiGasProvider = false
						L1ChainID = 1337
						HTTPHeaders = []

[RPC]
Host = "0.0.0.0"
Port = 5576
ReadTimeout = "2s"
WriteTimeout = "2s"
MaxRequestsPerIPAndSecond = 10

[ClaimSponsor]
DBPath = "/tmp/claimsopnsor"
Enabled = true
SenderAddr = "0xfa3b44587990f97ba8b6ba7e230a5f0e95d14b3d"
BridgeAddrL2 = "0xB7098a13a48EcE087d3DA15b2D28eCE0f89819B8"
MaxGas = 200000
RetryAfterErrorPeriod = "1s"
MaxRetryAttemptsAfterError = -1
WaitTxToBeMinedPeriod = "3s"
WaitOnEmptyQueue = "3s"
GasOffset = 0
	[ClaimSponsor.EthTxManager]
		FrequencyToMonitorTxs = "1s"
		WaitTxToBeMined = "2s"
		GetReceiptMaxTime = "250ms"
		GetReceiptWaitInterval = "1s"
		PrivateKeys = [
			{Path = "/app/keystore/claimsopnsor.keystore", Password = "testonly"},
		]
		ForcedGas = 0
		GasPriceMarginFactor = 1
		MaxGasPriceLimit = 0
		StoragePath = "/tmp/ethtxmanager-claimsponsor.db"
		ReadPendingL1Txs = false
		SafeStatusL1NumberOfBlocks = 5
		FinalizedStatusL1NumberOfBlocks = 10
			[ClaimSponsor.EthTxManager.Etherman]
				URL = "http://test-aggoracle-l2"
				MultiGasProvider = false
				L1ChainID = 1337
				HTTPHeaders = []

[BridgeL1Sync]
DBPath = "/tmp/bridgel1sync"
BlockFinality = "LatestBlock"
InitialBlockNum = 0
BridgeAddr = "0xB7098a13a48EcE087d3DA15b2D28eCE0f89819B8"
SyncBlockChunkSize = 100
RetryAfterErrorPeriod = "1s"
MaxRetryAttemptsAfterError = -1
WaitForNewBlocksPeriod = "3s"

[BridgeL2Sync]
DBPath = "/tmp/bridgel2sync"
BlockFinality = "LatestBlock"
InitialBlockNum = 0
BridgeAddr = "0xB7098a13a48EcE087d3DA15b2D28eCE0f89819B8"
SyncBlockChunkSize = 100
RetryAfterErrorPeriod = "1s"
MaxRetryAttemptsAfterError = -1
WaitForNewBlocksPeriod = "3s"

[LastGERSync]
DBPath = "/tmp/lastgersync"
BlockFinality = "LatestBlock"
InitialBlockNum = 0
GlobalExitRootL2Addr = "0xa40d5f56745a118d0906a34e69aec8c0db1cb8fa"
RetryAfterErrorPeriod = "1s"
MaxRetryAttemptsAfterError = -1
WaitForNewBlocksPeriod = "1s"
DownloadBufferSize = 100

[NetworkConfig.L1]
L1ChainID = 0
PolAddr = "0x0000000000000000000000000000000000000000"
ZkEVMAddr = "0x0000000000000000000000000000000000000000"
RollupManagerAddr = "0x0000000000000000000000000000000000000000"
GlobalExitRootManagerAddr = "0x0000000000000000000000000000000000000000"

`
