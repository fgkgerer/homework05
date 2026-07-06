package main

import (
	"context"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

/**
*
任务 2：合约代码生成 任务目标
使用 abigen 工具自动生成 Go 绑定代码，用于与 Sepolia 测试网络上的智能合约进行交互。 具体任务

编写智能合约
使用 Solidity 编写一个简单的智能合约，例如一个计数器合约。
编译智能合约，生成 ABI 和字节码文件。
使用 abigen 生成 Go 绑定代码
安装 abigen 工具。
使用 abigen 工具根据 ABI 和字节码文件生成 Go 绑定代码。
使用生成的 Go 绑定代码与合约交互
编写 Go 代码，使用生成的 Go 绑定代码连接到 Sepolia 测试网络上的智能合约。
调用合约的方法，例如增加计数器的值。
输出调用结果。
*/

func main() {
	rpcURL := os.Getenv("ETH_RPC_URL")
	contractAddr := os.Getenv("CONTRACT_ADDR")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	client, err := ethclient.DialContext(ctx, rpcURL)
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer client.Close()

	contract, err := NewBeggingContract(common.HexToAddress(contractAddr), client)
	if err != nil {
		log.Fatalf("failed to instantiate contract: %v", err)
	}

	// 调用 beg()
	msg, _ := contract.Beg(&bind.CallOpts{Context: ctx})
	log.Printf("beg() => %s", msg)

	// 调用 owner()
	owner, _ := contract.Owner(&bind.CallOpts{Context: ctx})
	log.Printf("owner() => %s", owner.Hex())

	// 调用 DEADLINE()
	deadline, _ := contract.DEADLINE(&bind.CallOpts{Context: ctx})
	log.Printf("DEADLINE() => %s", deadline.String())

	// 调用 getTopDonors()
	topDonors, _ := contract.GetTopDonors(&bind.CallOpts{Context: ctx})
	for i, d := range topDonors {
		if d.Amount.Cmp(big.NewInt(0)) > 0 {
			log.Printf("TopDonor #%d: %s — %s wei", i+1, d.DonorAddress.Hex(), d.Amount.String())
		}
	}
}
