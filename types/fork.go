package types

//default hard fork block height for bityuan real network
var (
	ForkV1               int64 = 1
	ForkV2AddToken       int64 = 1
	ForkV3               int64 = 1
	ForkV4AddManage      int64 = 1
	ForkV5Retrive        int64 = 1
	ForkV6TokenBlackList int64 = 1
	ForkV7BadTokenSymbol int64 = 1
	ForkBlockHash        int64 = 1
	ForkV9               int64 = 1
	ForkV10TradeBuyLimit int64 = 1
	ForkV11ManageExec    int64 = 100000
	ForkV12TransferExec  int64 = 100000
	ForkV13ExecKey       int64 = 200000
	ForkV14TxGroup       int64 = 200000
	ForkV15ResetTx0      int64 = 200000
	ForkV16Withdraw      int64 = 200000
	ForkV17EVM           int64 = 250000
	ForkV18Relay         int64 = 500000
	ForkV19TokenPrice    int64 = 300000
	ForkV20EVMState      int64 = 350000
)

//bityuan test net fork
func SetTestNetFork() {
	ForkV1 = 75260
	ForkV2AddToken = 100899
	ForkV3 = 110000
	ForkV4AddManage = 120000
	ForkV5Retrive = 180000
	ForkV6TokenBlackList = 190000
	ForkV7BadTokenSymbol = 184000
	ForkBlockHash = 208986 + 200
	ForkV9 = 350000
	ForkV10TradeBuyLimit = 301000
	ForkV11ManageExec = 400000
	ForkV12TransferExec = 408400
	ForkV13ExecKey = 408400
	ForkV14TxGroup = 408400
	ForkV15ResetTx0 = 453400
	ForkV16Withdraw = 480000
	ForkV17EVM = 500000
	ForkV18Relay = 570000
	ForkV19TokenPrice = 560000
	ForkV20EVMState = 650000
}

func SetForkToOne() {
	ForkV1 = 1
	ForkV2AddToken = 1
	ForkV3 = 1
	ForkV4AddManage = 1
	ForkV5Retrive = 1
	ForkV6TokenBlackList = 1
	ForkV7BadTokenSymbol = 1
	ForkBlockHash = 1
	ForkV9 = 1
	ForkV10TradeBuyLimit = 1
	ForkV11ManageExec = 1
	ForkV12TransferExec = 1
	ForkV13ExecKey = 1
	ForkV14TxGroup = 1
	ForkV15ResetTx0 = 1
	ForkV16Withdraw = 1
	ForkV17EVM = 1
	ForkV18Relay = 1
	ForkV19TokenPrice = 1
	ForkV20EVMState = 1
}

//paraName not used currently
func SetForkForPara(paraName string) {
	ForkV1 = 1
	ForkV2AddToken = 1
	ForkV3 = 1
	ForkV4AddManage = 1
	ForkV5Retrive = 1
	ForkV6TokenBlackList = 1
	ForkV7BadTokenSymbol = 1
	ForkBlockHash = 1
	ForkV9 = 1
	ForkV10TradeBuyLimit = 1
	ForkV11ManageExec = 1
	ForkV12TransferExec = 1
	ForkV13ExecKey = 1
	ForkV14TxGroup = 1
	ForkV15ResetTx0 = 1
	ForkV16Withdraw = 1
	ForkV17EVM = 1
	ForkV18Relay = 1
	ForkV19TokenPrice = 1
	ForkV20EVMState = 1
}

func IsMatchFork(height int64, fork int64) bool {
	if height == -1 || height >= fork {
		return true
	}
	return false
}