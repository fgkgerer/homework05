package main

import (
	"context"
	"crypto/ecdsa"
	"flag"
	"fmt"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	/**
	    * 任务 1：区块链读写 任务目标
	使用 Sepolia 测试网络实现基础的区块链交互，包括查询区块和发送交易。 具体任务

	环境搭建
	安装必要的开发工具，如 Go 语言环境、 go-ethereum 库。
	注册 Infura 账户，获取 Sepolia 测试网络的 API Key。
	查询区块
	编写 Go 代码，使用 ethclient 连接到 Sepolia 测试网络。
	实现查询指定区块号的区块信息，包括区块的哈希、时间戳、交易数量等。
	输出查询结果到控制台。
	发送交易
	准备一个 Sepolia 测试网络的以太坊账户，并获取其私钥。
	编写 Go 代码，使用 ethclient 连接到 Sepolia 测试网络。
	构造一笔简单的以太币转账交易，指定发送方、接收方和转账金额。
	对交易进行签名，并将签名后的交易发送到网络。
	输出交易的哈希值。
	*/

	sendMode := flag.Bool("send", false, "enable send transaction mode")
	toAddrHex := flag.String("to", "", "recipient address (required for send mode)")
	amountEth := flag.Float64("amount", 0, "amount in ETH (required for send mode)")

	blockNumberFlag := flag.Uint64("number", 0, "block number to query (0 means skip)")
	flag.Parse()

	rpcURL := os.Getenv("ETH_RPC_URL")
	if rpcURL == "" {
		log.Fatal("ETH_RPC_URL is not set")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	client, err := ethclient.DialContext(ctx, rpcURL)
	if err != nil {
		log.Fatalf("failed to connect to Ethereum node: %v", err)
	}
	defer client.Close()
	log.Println("connected to Ethereum node")

	if *sendMode {
		//发送交易
		log.Println("sending transaction")
		if *toAddrHex == "" || *amountEth <= 0 {
			log.Fatal("send mode requires --to and --amount flags")
		}
		privKeyHex := os.Getenv("SENDER_PRIVATE_KEY")
		if privKeyHex == "" {
			log.Fatal("SENDER_PRIVATE_KEY is not set (required for send mode)")
		}

		// 解析私钥
		privKey, err := crypto.HexToECDSA(trim0x(privKeyHex))
		if err != nil {
			log.Fatalf("invalid private key: %v", err)
		}

		// 获取发送方地址
		publicKey := privKey.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			log.Fatal("error casting public key to ECDSA")
		}
		fromAddr := crypto.PubkeyToAddress(*publicKeyECDSA)
		toAddr := common.HexToAddress(*toAddrHex)
		// 获取链 ID
		chainID, err := client.ChainID(ctx)
		if err != nil {
			log.Fatalf("failed to get chain id: %v", err)
		}

		// 获取 nonce
		nonce, err := client.PendingNonceAt(ctx, fromAddr)
		if err != nil {
			log.Fatalf("failed to get nonce: %v", err)
		}

		// 获取建议的 Gas 价格（使用 EIP-1559 动态费用）
		gasTipCap, err := client.SuggestGasTipCap(ctx)
		if err != nil {
			log.Fatalf("failed to get gas tip cap: %v", err)
		}

		// 获取 base fee，计算 fee cap
		header, err := client.HeaderByNumber(ctx, nil)
		if err != nil {
			log.Fatalf("failed to get header: %v", err)
		}
		baseFee := header.BaseFee
		if baseFee == nil {
			// 如果不支持 EIP-1559，使用传统 gas price
			gasPrice, err := client.SuggestGasPrice(ctx)
			if err != nil {
				log.Fatalf("failed to get gas price: %v", err)
			}
			baseFee = gasPrice
		}
		// fee cap = base fee * 2 + tip cap（简单策略）
		gasFeeCap := new(big.Int).Add(
			new(big.Int).Mul(baseFee, big.NewInt(2)),
			gasTipCap,
		)

		// 估算 Gas Limit（普通转账固定为 21000）
		gasLimit := uint64(21000)

		// 转换 ETH 金额为 Wei
		// amountEth * 1e18
		amountWei := new(big.Float).Mul(
			big.NewFloat(*amountEth),
			big.NewFloat(1e18),
		)
		valueWei, _ := amountWei.Int(nil)

		// 检查余额是否足够
		balance, err := client.BalanceAt(ctx, fromAddr, nil)
		if err != nil {
			log.Fatalf("failed to get balance: %v", err)
		}
		// 计算总费用：value + gasFeeCap * gasLimit
		totalCost := new(big.Int).Add(
			valueWei,
			new(big.Int).Mul(gasFeeCap, big.NewInt(int64(gasLimit))),
		)

		if balance.Cmp(totalCost) < 0 {
			log.Fatalf("insufficient balance: %s < %s", balance, totalCost)
		}

		tx := types.NewTx(&types.DynamicFeeTx{
			ChainID:   chainID,
			Nonce:     nonce,
			GasTipCap: gasTipCap,
			GasFeeCap: gasFeeCap,
			Gas:       gasLimit,
			To:        &toAddr,
			Value:     valueWei,
			Data:      nil,
		})

		// 签名交易
		signer := types.NewLondonSigner(chainID)
		signedTx, err := types.SignTx(tx, signer, privKey)
		if err != nil {
			log.Fatalf("failed to sign transaction: %v", err)
		}

		// 发送交易
		if err := client.SendTransaction(ctx, signedTx); err != nil {
			log.Fatalf("failed to send transaction: %v", err)
		}

		// 输出交易信息
		fmt.Println("=== Transaction Sent ===")
		fmt.Printf("From       : %s\n", fromAddr.Hex())
		fmt.Printf("To         : %s\n", toAddr.Hex())
		fmt.Printf("Value      : %s ETH (%s Wei)\n", fmt.Sprintf("%.6f", amountEth), valueWei.String())
		fmt.Printf("Gas Limit  : %d\n", gasLimit)
		fmt.Printf("Gas Tip Cap: %s Wei\n", gasTipCap.String())
		fmt.Printf("Gas Fee Cap: %s Wei\n", gasFeeCap.String())
		fmt.Printf("Nonce      : %d\n", nonce)
		fmt.Printf("Tx Hash    : %s\n", signedTx.Hash().Hex())
		// fmt.Println("\nTransaction is pending. Use --tx flag to query status:")
		// fmt.Printf("  go run main.go --tx %s\n", signedTx.Hash().Hex())

	} else {
		//查询区块
		log.Println("querying block number", *blockNumberFlag)

		//查询指定区块
		if *blockNumberFlag > 0 {
			num := big.NewInt(0).SetUint64(*blockNumberFlag)
			block, err := client.BlockByNumber(ctx, num)
			if err != nil {
				log.Fatalf("failed to get block: %v", err)
			}
			printBlock(fmt.Sprintf("Block %d", *blockNumberFlag), block)
		} else {
			//查询最新区块
			block, err := client.BlockByNumber(ctx, nil)
			if err != nil {
				log.Fatalf("failed to get latest block: %v", err)
			}
			printBlock("Latest Block", block)
		}
	}

}

// trim0x 移除十六进制字符串前缀 "0x"
func trim0x(s string) string {
	if len(s) >= 2 && s[:2] == "0x" {
		return s[2:]
	}
	return s
}

//打印指定区块号的区块信息，包括区块的哈希、时间戳、交易数量等。
func printBlock(title string, block *types.Block) {
	fmt.Println("======================================")
	fmt.Println(title)
	fmt.Println("======================================")
	fmt.Printf("Block: %+v\n", block)

	//基本信息
	fmt.Printf("Number: %d\n", block.NumberU64()) //或者 block.Number().Uint64()
	fmt.Printf("Hash: %s\n", block.Hash().Hex())
	fmt.Printf("Parent Hash  : %s\n", block.ParentHash().Hex())

	//时间信息
	blockTime := time.Unix(int64(block.Time()), 0)
	fmt.Printf("Time: %s\n", blockTime.Format(time.DateTime))
	fmt.Printf("Time (Local) : %s\n", blockTime.Local().Format("2006-01-02 15:04:05 MST"))
	fmt.Printf("Time: %s\n", blockTime.Format(time.RFC3339))

	//交易信息
	transactions := block.Transactions()
	txCount := len(transactions)
	fmt.Printf("Tx Count     : %d\n", txCount)
}
