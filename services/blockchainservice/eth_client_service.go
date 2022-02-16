package blockchainservice

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"sync"
	"time"
	"worker-transaction/appconfig"
	"worker-transaction/consts"

	"github.com/avast/retry-go"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	ethClient           sync.Once
	blockchainClient    *ethclient.Client
	newBlockchainClient *MyClient
	req                 uint64 = 0
	lockReq                    = sync.Mutex{}
)

type MyClient struct {
	cli *ethclient.Client
}

func (this *MyClient) GetClient() *ethclient.Client {
	return this.cli
}

func getAttempt(attempts []uint) uint {
	if len(attempts) > 0 {
		return attempts[0]
	}

	return consts.DEFAULT_ATTEMPT
}

func (this *MyClient) TransactionReceipt(ctx context.Context, txHash common.Hash, attempts ...uint) (receipt *types.Receipt, err error) {
	retry.Do(func() error {
		receipt, err = this.cli.TransactionReceipt(ctx, txHash)
		return err
	},
		retry.OnRetry(func(n uint, err error) {
			log.Println("func TransactionReceipt time attempts: ", n, ", err: ", err)
		}),
		retry.RetryIf(func(err error) bool {
			return err != nil
		}),
		retry.Attempts(getAttempt(attempts)),
		retry.MaxJitter(time.Millisecond))

	return receipt, err
}

func (this *MyClient) Close() {
	this.cli.Close()
}

func (this *MyClient) ChainID(ctx context.Context, attempts ...uint) (num *big.Int, err error) {
	retry.Do(func() error {
		num, err = this.cli.ChainID(ctx)
		return err
	},
		retry.OnRetry(func(n uint, err error) {
			log.Println("func ChainID time attempts: ", n, ", err: ", err)
		}),
		retry.RetryIf(func(err error) bool {
			return err != nil
		}),
		retry.Attempts(getAttempt(attempts)),
		retry.MaxJitter(time.Millisecond))

	return
}

func (this *MyClient) BlockByHash(ctx context.Context, hash common.Hash, attempts ...uint) (b *types.Block, err error) {
	retry.Do(func() error {
		b, err = this.cli.BlockByHash(ctx, hash)
		return err
	},
		retry.OnRetry(func(n uint, err error) {
			log.Println("func BlockByHash time attempts: ", n, ", err: ", err)
		}),
		retry.RetryIf(func(err error) bool {
			return err != nil
		}),
		retry.Attempts(getAttempt(attempts)),
		retry.MaxJitter(time.Millisecond))

	return
}

func (this *MyClient) BlockByNumber(ctx context.Context, number *big.Int, attempts ...uint) (b *types.Block, err error) {
	retry.Do(func() error {
		b, err = this.cli.BlockByNumber(ctx, number)
		return err
	}, retry.OnRetry(func(n uint, err error) {
		log.Println("func BlockByNumber time attempts: ", n, ", err: ", err)
	}),
		retry.RetryIf(func(err error) bool {
			return err != nil
		}),
		retry.Attempts(getAttempt(attempts)),
		retry.MaxJitter(time.Millisecond))

	return
}

func (this *MyClient) BlockNumber(ctx context.Context, attempts ...uint) (n uint64, err error) {
	retry.Do(func() error {
		n, err = this.cli.BlockNumber(ctx)
		return err
	},
		retry.OnRetry(func(n uint, err error) {
			log.Println("func BlockNumber time attempts: ", n, ", err: ", err)
		}),
		retry.RetryIf(func(err error) bool {
			return err != nil
		}),
		retry.Attempts(getAttempt(attempts)),
		retry.MaxJitter(time.Millisecond))

	return
}

func (this *MyClient) HeaderByHash(ctx context.Context, hash common.Hash, attempts ...uint) (h *types.Header, err error) {
	retry.Do(func() error {
		h, err = this.cli.HeaderByHash(ctx, hash)
		return err
	},
		retry.OnRetry(func(n uint, err error) {
			log.Println("func HeaderByHash time attempts: ", n, ", err: ", err)
		}),
		retry.RetryIf(func(err error) bool {
			return err != nil
		}),
		retry.Attempts(getAttempt(attempts)),
		retry.MaxJitter(time.Millisecond))

	return
}

func (this *MyClient) HeaderByNumber(ctx context.Context, number *big.Int, attempts ...uint) (h *types.Header, err error) {
	retry.Do(func() error {
		h, err = this.cli.HeaderByNumber(ctx, number)
		return err
	},
		retry.OnRetry(func(n uint, err error) {
			log.Println("func HeaderByNumber time attempts: ", n, ", err: ", err)
		}),
		retry.RetryIf(func(err error) bool {
			return err != nil
		}),
		retry.Attempts(getAttempt(attempts)),
		retry.MaxJitter(time.Millisecond))

	return
}
func (this *MyClient) TransactionByHash(ctx context.Context, hash common.Hash, attempts ...uint) (tx *types.Transaction, isPending bool, err error) {
	retry.Do(func() error {
		tx, isPending, err = this.cli.TransactionByHash(ctx, hash)
		return err
	},
		retry.OnRetry(func(n uint, err error) {
			log.Println("func TransactionByHash time attempts: ", n, ", err: ", err)
		}),
		retry.RetryIf(func(err error) bool {
			return err != nil
		}),
		retry.Attempts(getAttempt(attempts)),
		retry.MaxJitter(time.Millisecond))

	return
}
func (this *MyClient) TransactionSender(ctx context.Context, tx *types.Transaction, block common.Hash, index uint, attempts ...uint) (addr common.Address, err error) {
	retry.Do(func() error {
		addr, err = this.cli.TransactionSender(ctx, tx, block, index)
		return err
	},
		retry.OnRetry(func(n uint, err error) {
			log.Println("func TransactionSender time attempts: ", n, ", err: ", err)
		}),
		retry.RetryIf(func(err error) bool {
			return err != nil
		}),
		retry.Attempts(getAttempt(attempts)),
		retry.MaxJitter(time.Millisecond))

	return
}
func (this *MyClient) TransactionCount(ctx context.Context, blockHash common.Hash, attempts ...uint) (c uint, err error) {
	retry.Do(func() error {
		c, err = this.cli.TransactionCount(ctx, blockHash)
		return err
	},
		retry.OnRetry(func(n uint, err error) {
			log.Println("func TransactionCount time attempts: ", n, ", err: ", err)
		}),
		retry.RetryIf(func(err error) bool {
			return err != nil
		}),
		retry.Attempts(getAttempt(attempts)),
		retry.MaxJitter(time.Millisecond))

	return
}

func (this *MyClient) TransactionInBlock(ctx context.Context, blockHash common.Hash, index uint, attempts ...uint) (t *types.Transaction, err error) {
	retry.Do(func() error {
		t, err = this.cli.TransactionInBlock(ctx, blockHash, index)
		return err
	},
		retry.OnRetry(func(n uint, err error) {
			log.Println("func TransactionInBlock time attempts: ", n, ", err: ", err)
		}),
		retry.RetryIf(func(err error) bool {
			return err != nil
		}),
		retry.Attempts(getAttempt(attempts)),
		retry.MaxJitter(time.Millisecond))

	return
}
func (this *MyClient) SyncProgress(ctx context.Context, attempts ...uint) (info *ethereum.SyncProgress, err error) {
	retry.Do(func() error {
		info, err = this.cli.SyncProgress(ctx)
		return err
	},
		retry.OnRetry(func(n uint, err error) {
			log.Println("func SyncProgress time attempts: ", n, ", err: ", err)
		}),
		retry.RetryIf(func(err error) bool {
			return err != nil
		}),
		retry.Attempts(getAttempt(attempts)),
		retry.MaxJitter(time.Millisecond))

	return
}
func (this *MyClient) SubscribeNewHead(ctx context.Context, ch chan<- *types.Header, attempts ...uint) (sub ethereum.Subscription, err error) {
	retry.Do(func() error {
		sub, err = this.cli.SubscribeNewHead(ctx, ch)
		return err
	},
		retry.OnRetry(func(n uint, err error) {
			log.Println("func SubscribeNewHead time attempts: ", n, ", err: ", err)
		}),
		retry.RetryIf(func(err error) bool {
			return err != nil
		}),
		retry.Attempts(getAttempt(attempts)),
		retry.MaxJitter(time.Millisecond))

	return
}

func (this *MyClient) NetworkID(ctx context.Context, attempts ...uint) (id *big.Int, err error) {
	retry.Do(func() error {
		id, err = this.cli.NetworkID(ctx)
		return err
	},
		retry.OnRetry(func(n uint, err error) {
			log.Println("func NetworkID time attempts: ", n, ", err: ", err)
		}),
		retry.RetryIf(func(err error) bool {
			return err != nil
		}),
		retry.Attempts(getAttempt(attempts)),
		retry.MaxJitter(time.Millisecond))

	return
}

func (this *MyClient) BalanceAt(ctx context.Context, account common.Address, blockNumber *big.Int, attempts ...uint) (bal *big.Int, err error) {
	retry.Do(func() error {
		bal, err = this.cli.BalanceAt(ctx, account, blockNumber)
		return err
	},
		retry.OnRetry(func(n uint, err error) {
			log.Println("func BalanceAt time attempts: ", n, ", err: ", err)
		}),
		retry.RetryIf(func(err error) bool {
			return err != nil
		}),
		retry.Attempts(getAttempt(attempts)),
		retry.MaxJitter(time.Millisecond))

	return
}

func (this *MyClient) StorageAt(ctx context.Context, account common.Address, key common.Hash, blockNumber *big.Int, attempts ...uint) (arrB []byte, err error) {
	retry.Do(func() error {
		arrB, err = this.cli.StorageAt(ctx, account, key, blockNumber)
		return err
	},
		retry.OnRetry(func(n uint, err error) {
			log.Println("func StorageAt time attempts: ", n, ", err: ", err)
		}),
		retry.RetryIf(func(err error) bool {
			return err != nil
		}),
		retry.Attempts(getAttempt(attempts)),
		retry.MaxJitter(time.Millisecond))

	return
}

func (this *MyClient) CodeAt(ctx context.Context, account common.Address, blockNumber *big.Int, attempts ...uint) (arrB []byte, err error) {
	retry.Do(func() error {
		arrB, err = this.cli.CodeAt(ctx, account, blockNumber)
		return err
	},
		retry.OnRetry(func(n uint, err error) {
			log.Println("func CodeAt time attempts: ", n, ", err: ", err)
		}),
		retry.RetryIf(func(err error) bool {
			return err != nil
		}),
		retry.Attempts(getAttempt(attempts)),
		retry.MaxJitter(time.Millisecond))

	return
}

func (this *MyClient) NonceAt(ctx context.Context, account common.Address, blockNumber *big.Int, attempts ...uint) (n uint64, err error) {
	retry.Do(func() error {
		n, err = this.cli.NonceAt(ctx, account, blockNumber)
		return err
	},
		retry.OnRetry(func(n uint, err error) {
			log.Println("func NonceAt time attempts: ", n, ", err: ", err)
		}),
		retry.RetryIf(func(err error) bool {
			return err != nil
		}),
		retry.Attempts(getAttempt(attempts)),
		retry.MaxJitter(time.Millisecond))

	return
}
func (this *MyClient) FilterLogs(ctx context.Context, q ethereum.FilterQuery, attempts ...uint) (aLog []types.Log, err error) {
	retry.Do(func() error {
		aLog, err = this.cli.FilterLogs(ctx, q)
		return err
	},
		retry.OnRetry(func(n uint, err error) {
			log.Println("func FilterLogs time attempts: ", n, ", err: ", err)
		}),
		retry.RetryIf(func(err error) bool {
			return err != nil
		}),
		retry.Attempts(getAttempt(attempts)),
		retry.MaxJitter(time.Millisecond))

	return
}
func (this *MyClient) SubscribeFilterLogs(ctx context.Context, q ethereum.FilterQuery, ch chan<- types.Log, attempts ...uint) (sub ethereum.Subscription, err error) {
	retry.Do(func() error {
		sub, err = this.cli.SubscribeFilterLogs(ctx, q, ch)
		return err
	},
		retry.OnRetry(func(n uint, err error) {
			log.Println("func SubscribeFilterLogs time attempts: ", n, ", err: ", err)
		}),
		retry.RetryIf(func(err error) bool {
			return err != nil
		}),
		retry.Attempts(getAttempt(attempts)),
		retry.MaxJitter(time.Millisecond))

	return
}
func (this *MyClient) PendingBalanceAt(ctx context.Context, account common.Address, attempts ...uint) (bal *big.Int, err error) {
	retry.Do(func() error {
		bal, err = this.cli.PendingBalanceAt(ctx, account)
		return err
	},
		retry.OnRetry(func(n uint, err error) {
			log.Println("func PendingBalanceAt time attempts: ", n, ", err: ", err)
		}),
		retry.RetryIf(func(err error) bool {
			return err != nil
		}),
		retry.Attempts(getAttempt(attempts)),
		retry.MaxJitter(time.Millisecond))

	return
}
func (this *MyClient) PendingStorageAt(ctx context.Context, account common.Address, key common.Hash, attempts ...uint) (bArr []byte, err error) {
	retry.Do(func() error {
		bArr, err = this.cli.PendingStorageAt(ctx, account, key)
		return err
	},
		retry.OnRetry(func(n uint, err error) {
			log.Println("func PendingStorageAt time attempts: ", n, ", err: ", err)
		}),
		retry.RetryIf(func(err error) bool {
			return err != nil
		}),
		retry.Attempts(getAttempt(attempts)),
		retry.MaxJitter(time.Millisecond))

	return
}
func (this *MyClient) PendingCodeAt(ctx context.Context, account common.Address, attempts ...uint) (bArr []byte, err error) {
	retry.Do(func() error {
		bArr, err = this.cli.PendingCodeAt(ctx, account)
		return err
	},
		retry.OnRetry(func(n uint, err error) {
			log.Println("func PendingCodeAt time attempts: ", n, ", err: ", err)
		}),
		retry.RetryIf(func(err error) bool {
			return err != nil
		}),
		retry.Attempts(getAttempt(attempts)),
		retry.MaxJitter(time.Millisecond))

	return
}
func (this *MyClient) PendingNonceAt(ctx context.Context, account common.Address, attempts ...uint) (n uint64, err error) {
	retry.Do(func() error {
		n, err = this.cli.PendingNonceAt(ctx, account)
		return err
	},
		retry.OnRetry(func(n uint, err error) {
			log.Println("func PendingNonceAt time attempts: ", n, ", err: ", err)
		}),
		retry.RetryIf(func(err error) bool {
			return err != nil
		}),
		retry.Attempts(getAttempt(attempts)),
		retry.MaxJitter(time.Millisecond))

	return
}

func (this *MyClient) PendingTransactionCount(ctx context.Context, attempts ...uint) (n uint, err error) {
	retry.Do(func() error {
		n, err = this.cli.PendingTransactionCount(ctx)
		return err
	},
		retry.OnRetry(func(n uint, err error) {
			log.Println("func PendingTransactionCount time attempts: ", n, ", err: ", err)
		}),
		retry.RetryIf(func(err error) bool {
			return err != nil
		}),
		retry.Attempts(getAttempt(attempts)),
		retry.MaxJitter(time.Millisecond))

	return
}

func (this *MyClient) CallContract(ctx context.Context, msg ethereum.CallMsg, blockNumber *big.Int, attempts ...uint) (bArr []byte, err error) {
	retry.Do(func() error {
		bArr, err = this.cli.CallContract(ctx, msg, blockNumber)
		return err
	},
		retry.OnRetry(func(n uint, err error) {
			log.Println("func CallContract time attempts: ", n, ", err: ", err)
		}),
		retry.RetryIf(func(err error) bool {
			return err != nil
		}),
		retry.Attempts(getAttempt(attempts)),
		retry.MaxJitter(time.Millisecond))

	return
}

func (this *MyClient) PendingCallContract(ctx context.Context, msg ethereum.CallMsg, attempts ...uint) (bArr []byte, err error) {
	retry.Do(func() error {
		bArr, err = this.cli.PendingCallContract(ctx, msg)
		return err
	},
		retry.OnRetry(func(n uint, err error) {
			log.Println("func PendingCallContract time attempts: ", n, ", err: ", err)
		}),
		retry.RetryIf(func(err error) bool {
			return err != nil
		}),
		retry.Attempts(getAttempt(attempts)),
		retry.MaxJitter(time.Millisecond))

	return
}

func (this *MyClient) SuggestGasPrice(ctx context.Context, attempts ...uint) (p *big.Int, err error) {
	retry.Do(func() error {
		p, err = this.cli.SuggestGasPrice(ctx)
		return err
	},
		retry.OnRetry(func(n uint, err error) {
			log.Println("func SuggestGasPrice time attempts: ", n, ", err: ", err)
		}),
		retry.RetryIf(func(err error) bool {
			return err != nil
		}),
		retry.Attempts(getAttempt(attempts)),
		retry.MaxJitter(time.Millisecond))

	return
}

func (this *MyClient) SuggestGasTipCap(ctx context.Context, attempts ...uint) (p *big.Int, err error) {
	retry.Do(func() error {
		p, err = this.cli.SuggestGasTipCap(ctx)
		return err
	},
		retry.OnRetry(func(n uint, err error) {
			log.Println("func SuggestGasTipCap time attempts: ", n, ", err: ", err)
		}),
		retry.RetryIf(func(err error) bool {
			return err != nil
		}),
		retry.Attempts(getAttempt(attempts)),
		retry.MaxJitter(time.Millisecond))

	return
}

func (this *MyClient) EstimateGas(ctx context.Context, msg ethereum.CallMsg, attempts ...uint) (g uint64, err error) {
	retry.Do(func() error {
		g, err = this.cli.EstimateGas(ctx, msg)
		return err
	},
		retry.OnRetry(func(n uint, err error) {
			log.Println("func EstimateGas time attempts: ", n, ", err: ", err)
		}),
		retry.RetryIf(func(err error) bool {
			return err != nil
		}),
		retry.Attempts(getAttempt(attempts)),
		retry.MaxJitter(time.Millisecond))

	return
}

func (this *MyClient) SendTransaction(ctx context.Context, tx *types.Transaction, attempts ...uint) (err error) {
	retry.Do(func() error {
		err = this.cli.SendTransaction(ctx, tx)
		return err
	},
		retry.OnRetry(func(n uint, err error) {
			log.Println("func SendTransaction time attempts: ", n, ", err: ", err)
		}),
		retry.RetryIf(func(err error) bool {
			return err != nil
		}),
		retry.Attempts(getAttempt(attempts)),
		retry.MaxJitter(time.Millisecond))

	return
}

func newClient(cli *ethclient.Client) *MyClient {
	return &MyClient{
		cli: cli,
	}
}

func GetEthClientService(chain string) *ethclient.Client {
	ethClient.Do(func() {
		fmt.Println(getEndpoint(chain))
		client, err := ethclient.Dial(getEndpoint(chain))
		if err != nil {
			log.Println(err.Error(), "err.Error() eth_client_service.go:574")
			log.Println("can't connect client")
		}
		blockchainClient = client

	})

	return blockchainClient
}

func GetNewEthClientService(chain string) *MyClient {
	client, err := ethclient.Dial(getEndpointNew(chain))
	if err != nil {
		log.Println(err.Error(), "err.Error() services/blockchainservice/eth_client_service.go:35")
		log.Println("can't connect client")
	}
	newBlockchainClient = newClient(client)

	return newBlockchainClient
}

func getEndpoint(chain string) string {
	switch chain {
	case consts.BSC_CHAIN:
		return appconfig.Config.BnbRPC
	case consts.POLYGON_CHAIN:
		return appconfig.Config.PolygonRPC
	default:
		return appconfig.Config.BnbRPC
	}
}

func getEndpointNew(chain string) string {
	switch chain {
	case consts.BSC_CHAIN:
		return getRandomInListBsc()
	case consts.POLYGON_CHAIN:
		return getRandomInListPolygon()
	default:
		return appconfig.Config.BnbRPC
	}
}

func getRandomInListBsc() string {
	lockReq.Lock()
	defer lockReq.Unlock()
	req++
	if req%1000 == 0 {
		log.Println("req: ", req)
	}
	return appconfig.Config.BnbRPCS[req%uint64(len(appconfig.Config.BnbRPCS))]
}

func getRandomInListPolygon() string {
	lockReq.Lock()
	defer lockReq.Unlock()
	req++
	if req%1000 == 0 {
		log.Println("req: ", req)
	}
	return appconfig.Config.PolygonRPCs[req%uint64(len(appconfig.Config.PolygonRPCs))]
}

func getGas(input []byte, price *big.Int, fromAddress, contractAddress common.Address) (
	gasLimit uint64, gasPrice *big.Int, err error) {
	//gasLimit, err = GetEthClientService().EstimateGas(context.Background(), ethereum.CallMsg{
	//From:  fromAddress,
	//To:    &contractAddress,
	//Data:  input,
	//Value: price,
	//})
	//if err != nil {
	//	log.Println(err.Error(), "err.Error() getGas")
	//	return
	//}
	gasLimit = consts.BSC_GAS_LIMIT
	gasPrice, err = GetEthClientService(consts.BSC_CHAIN).SuggestGasPrice(context.Background())
	if err != nil {
		gasPrice, err = GetNewEthClientService(consts.BSC_CHAIN).SuggestGasPrice(context.Background())
		if err != nil {
			log.Println(err.Error(), "err.Error() getGas")
			return
		}
	}
	return
}
