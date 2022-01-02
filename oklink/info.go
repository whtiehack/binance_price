package oklink


type Info struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	DetailMsg string `json:"detailMsg"`
	Data Data `json:"data"`
}
//type Market struct {
//	Symbol string `json:"symbol"`
//	MarketSymbol string `json:"marketSymbol"`
//	Price float64 `json:"price"`
//	Volume24H float64 `json:"volume24h"`
//	PercentChange1H float64 `json:"percentChange1h"`
//	PercentChange24H float64 `json:"percentChange24h"`
//	PercentChange7D float64 `json:"percentChange7d"`
//	MarketCap float64 `json:"marketCap"`
//	Timestamp int `json:"timestamp"`
//}
//type Address struct {
//	ValidAddressCount int `json:"validAddressCount"`
//	ValidAddressCountDiff int `json:"validAddressCountDiff"`
//	NewAddressCount24H int `json:"newAddressCount24h"`
//	ActiveAddressCount int `json:"activeAddressCount"`
//	ActiveAddressCountDiff int `json:"activeAddressCountDiff"`
//}
//type Block struct {
//	Height int `json:"height"`
//	FirstExchangeHistoricalTime int `json:"firstExchangeHistoricalTime"`
//	FirstBlockTime int64 `json:"firstBlockTime"`
//	FirstBlockHeight int `json:"firstBlockHeight"`
//	AvgBlockInterval int `json:"avgBlockInterval"`
//	AvgBlockSize24H float64 `json:"avgBlockSize24h"`
//	AvgBlockSize24HPercent float64 `json:"avgBlockSize24hPercent"`
//	MediaBlockSize float64 `json:"mediaBlockSize"`
//	HalveTime int `json:"halveTime"`
//	MinerIncome float64 `json:"minerIncome"`
//	LatestBlockBaseFee float64 `json:"latestBlockBaseFee"`
//}
//type Fee struct {
//	BestFeePerKbyte interface{} `json:"bestFeePerKbyte"`
//	BestFeePerKwu interface{} `json:"bestFeePerKwu"`
//	BestFeePerKvbyte interface{} `json:"bestFeePerKvbyte"`
//	BestGasPrice interface{} `json:"bestGasPrice"`
//}
//type GlobalDifficulty struct {
//	CurrentDiffculty string `json:"currentDiffculty"`
//	RealCurrentDiffculty float64 `json:"realCurrentDiffculty"`
//	CurrentDiffcultyPercentChange float64 `json:"currentDiffcultyPercentChange"`
//	CurrentDiffcultyChangeTime int `json:"currentDiffcultyChangeTime"`
//	CurrentDiffcultyChangeHeight int `json:"currentDiffcultyChangeHeight"`
//	NextDiffculty interface{} `json:"nextDiffculty"`
//	RealNextDiffculty float64 `json:"realNextDiffculty"`
//	NextDiffcultyChangeTime int `json:"nextDiffcultyChangeTime"`
//	NextDiffcultyChangeHeight int `json:"nextDiffcultyChangeHeight"`
//	NextDiffcultyPercentChange float64 `json:"nextDiffcultyPercentChange"`
//	NextDifficultyChangeBlock int `json:"nextDifficultyChangeBlock"`
//}
//type Hashes struct {
//	GlobalHashes string `json:"globalHashes"`
//	GlobalHashesPercentChange24H float64 `json:"globalHashesPercentChange24h"`
//}
//type Mine struct {
//	AvgMineReward24H float64 `json:"avgMineReward24h"`
//	MinerIncomePerUnit float64 `json:"minerIncomePerUnit"`
//	MinerIncomePerUnitAddFee float64 `json:"minerIncomePerUnitAddFee"`
//	MinerIncomePerUnitCoin float64 `json:"minerIncomePerUnitCoin"`
//}
//type ReduceReward struct {
//	NextReduceRewardTime int `json:"nextReduceRewardTime"`
//	NextReduceRewardHeight int `json:"nextReduceRewardHeight"`
//}



type Transaction struct {
	PendingTransactionCount float64 `json:"pendingTransactionCount"`
	PendingTransactionSize float64 `json:"pendingTransactionSize"`
	TransactionValue24H float64 `json:"transactionValue24h"`
	TransactionCount24H float64 `json:"transactionCount24h"`
	TotalTransactionCount float64 `json:"totalTransactionCount"`
	TranRate float64 `json:"tranRate"`
	AvgTransactionCount24H float64 `json:"avgTransactionCount24h"`
	AvgTransactionCount24HPercent float64 `json:"avgTransactionCount24hPercent"`
}
//type Usdt struct {
//	Fee interface{} `json:"fee"`
//	OmiUsdtTotalSupply float64 `json:"omiUsdtTotalSupply"`
//	TrxUsdtTotalSupply float64 `json:"trxUsdtTotalSupply"`
//	EthUsdtTotalSupply float64 `json:"ethUsdtTotalSupply"`
//	WeekAddCoin float64 `json:"weekAddCoin"`
//	WeekDestoryCoin float64 `json:"weekDestoryCoin"`
//}
//type Eth struct {
//	ValueCoinAmount int `json:"valueCoinAmount"`
//	ContractAmount int `json:"contractAmount"`
//	Erc20Amount int `json:"erc20Amount"`
//	InternalTransactionAmount int `json:"internalTransactionAmount"`
//	TokenAmount int `json:"tokenAmount"`
//	UncleBlockAmount int `json:"uncleBlockAmount"`
//	NewErc20Amount int `json:"newErc20Amount"`
//	GasFee float64 `json:"gasFee"`
//	Erc20Value float64 `json:"erc20Value"`
//	Erc721Count int `json:"erc721Count"`
//	NewErc721Count int `json:"newErc721Count"`
//	TotalBurnt float64 `json:"totalBurnt"`
//	Erc1155Count int `json:"erc1155Count"`
//	NewErc1155Count int `json:"newErc1155Count"`
//}


type Data struct {
	ID int `json:"id"`
	CoinType string `json:"coinType"`
	Name string `json:"name"`
	Symbol string `json:"symbol"`
	CoinName string `json:"coinName"`
	//FullName interface{} `json:"fullName"`
	//WebsiteSlug string `json:"websiteSlug"`
	//IconPath interface{} `json:"iconPath"`
	//Rank int `json:"rank"`
	//Mineable bool `json:"mineable"`
	//Algorithm string `json:"algorithm"`
	//ProofType interface{} `json:"proofType"`
	//FullyPremined bool `json:"fullyPremined"`
	//PreMinedValue interface{} `json:"preMinedValue"`
	//NetHashesPerSecond interface{} `json:"netHashesPerSecond"`
	//BlockReward interface{} `json:"blockReward"`
	//BlockPeriodTime interface{} `json:"blockPeriodTime"`
	//CirculatingSupply float64 `json:"circulatingSupply"`
	//TotalSupply float64 `json:"totalSupply"`
	//MaxSupply float64 `json:"maxSupply"`
	//TokenAddress interface{} `json:"tokenAddress"`
	//FirstBlockTime int64 `json:"firstBlockTime"`
	//FirstBlockHeight interface{} `json:"firstBlockHeight"`
	//FirstHistoricalData interface{} `json:"firstHistoricalData"`
	//LastHistoricalData int64 `json:"lastHistoricalData"`
	//LastSyncTime int64 `json:"lastSyncTime"`
	//CreateTime int64 `json:"createTime"`
	//UpdateTime int64 `json:"updateTime"`
	//IcoPrice string `json:"icoPrice"`
	//Market Market `json:"market"`
	//Address Address `json:"address"`
	//Block Block `json:"block"`
	//Fee Fee `json:"fee"`
	//GlobalDifficulty GlobalDifficulty `json:"globalDifficulty"`
	//Hashes Hashes `json:"hashes"`
	//Mine Mine `json:"mine"`
	//ReduceReward ReduceReward `json:"reduceReward"`
	Transaction Transaction `json:"transaction"`
	//Okchain interface{} `json:"okchain"`
	//Usdt Usdt `json:"usdt"`
	//Eth Eth `json:"eth"`
	//TrxUsdtTotalSupply int `json:"trxUsdtTotalSupply"`
	//EthUsdtTotalSupply int `json:"ethUsdtTotalSupply"`
	//CoreAlgorithm string `json:"coreAlgorithm"`
	//MasterNodeCount int `json:"masterNodeCount"`
	//DiffChangeRemainBlock int `json:"diffChangeRemainBlock"`
	//PlatformID interface{} `json:"platformId"`
	//Platform interface{} `json:"platform"`
}