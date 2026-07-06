# abigen 生成 Go 绑定代码 — 操作步骤与结果

## 环境信息

| 项目 | 值 |
|---|---|
| 合约 | `BeggingContract.sol` |
| 合约地址 | `0xc70a18ea7d59ba4b1757df851f7c5e05006aa516` |
| 网络 | Sepolia 测试网 |
| Solidity 版本 | ^0.8.13 |
| Go 版本 | 1.26.2 |
| go-ethereum 版本 | v1.17.4 |
| abigen 路径 | `/d/gopath/bin/abigen` |
| Forge 产物路径 | `D:\meta\homework\solidity\out\BeggingContract.sol\BeggingContract.json` |

---

## 操作步骤

### 1. 编译 Solidity 合约

```bash
cd D:\meta\homework\solidity
forge build
```

产物生成在 `out/BeggingContract.sol/BeggingContract.json`，包含 ABI 和 bytecode。

### 2. 用 jq 提取 ABI 和字节码

```bash
cd D:\meta\homework\homework05\task2

# 提取 ABI
jq '.abi' \
  "D:/meta/homework/solidity/out/BeggingContract.sol/BeggingContract.json" \
  > BeggingContract.abi

# 提取字节码 (bytecode.object)
jq -r '.bytecode.object' \
  "D:/meta/homework/solidity/out/BeggingContract.sol/BeggingContract.json" \
  > BeggingContract.bin
```

- ABI 文件大小: 5,109 bytes
- Bytecode 文件大小: 13,362 bytes

### 3. 使用 abigen 生成 Go 绑定代码

```bash
abigen \
  --abi BeggingContract.abi \
  --bin BeggingContract.bin \
  --pkg main \
  --type BeggingContract \
  --out BeggingContract.go
```

### 4. 清理临时文件

```bash
rm -f BeggingContract.abi BeggingContract.bin
```

---

## 生成的文件

| 文件 | 说明 |
|---|---|
| `BeggingContract.go` | abigen 自动生成的 Go 绑定代码，package `main` |

生成的关键类型和方法：

```go
// 结构体
type BeggingContractDonor struct {
    DonorAddress common.Address
    Amount       *big.Int
}

// 实例化已部署的合约
func NewBeggingContract(address common.Address, backend bind.ContractBackend) (*BeggingContract, error)

// 部署新合约
func DeployBeggingContract(auth *bind.TransactOpts, backend bind.ContractBackend,
    initialOwner common.Address, durationDays *big.Int) (common.Address, *types.Transaction, *BeggingContract, error)

// 只读方法 (Caller)
func (_BeggingContract *BeggingContractCaller) Beg(opts *bind.CallOpts) (string, error)
func (_BeggingContract *BeggingContractCaller) Owner(opts *bind.CallOpts) (common.Address, error)
func (_BeggingContract *BeggingContractCaller) DEADLINE(opts *bind.CallOpts) (*big.Int, error)
func (_BeggingContract *BeggingContractCaller) GetDonation(opts *bind.CallOpts, donor common.Address) (*big.Int, error)
func (_BeggingContract *BeggingContractCaller) GetTopDonors(opts *bind.CallOpts) ([3]BeggingContractDonor, error)
func (_BeggingContract *BeggingContractCaller) Donations(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error)
func (_BeggingContract *BeggingContractCaller) TopDonors(opts *bind.CallOpts, arg0 *big.Int) (struct{...}, error)

// 写入方法 (Transactor)
func (_BeggingContract *BeggingContractTransactor) Donate(opts *bind.TransactOpts) (*types.Transaction, error)
func (_BeggingContract *BeggingContractTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error)
func (_BeggingContract *BeggingContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error)
func (_BeggingContract *BeggingContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error)

// 事件过滤 (Filterer)
func (_BeggingContract *BeggingContractFilterer) FilterDonation(...)
func (_BeggingContract *BeggingContractFilterer) WatchDonation(...)
func (_BeggingContract *BeggingContractFilterer) FilterOwnershipTransferred(...)
func (_BeggingContract *BeggingContractFilterer) FilterWithdraw(...)
```

### 5. 编写 main.go 调用代码

```go
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
```

### 6. 运行验证

```bash
ETH_RPC_URL="https://sepolia.infura.io/v3/..." \
CONTRACT_ADDR="0xc70a18ea7d59ba4b1757df851f7c5e05006aa516" \
go run .
```

---

## 执行结果

```
2026/07/06 17:07:03 beg() => Please give me some money!
2026/07/06 17:07:03 owner() => 0xDD87b63dd7b73f501992F92a25DeeE79445869e5
2026/07/06 17:07:04 DEADLINE() => 1780582896
2026/07/06 17:07:04 Top Donors:
2026/07/06 17:07:04   #1: 0x28a79B67AD5903656c83c289fe8ca9FeDBfD6533 — 10000000000000000 wei
2026/07/06 17:07:04   #2: 0xDD87b63dd7b73f501992F92a25DeeE79445869e5 — 2000000000000000 wei
```

| 方法 | 返回值 | 说明 |
|---|---|---|
| `beg()` | `"Please give me some money!"` | pure 函数，直接返回字符串 |
| `owner()` | `0xDD87b63dd7b73f501992F92a25DeeE79445869e5` | 合约所有者地址 |
| `DEADLINE()` | `1780582896` | 募资截止时间戳 |
| `getTopDonors()` | #1 `0x28a79B67...` 捐赠 0.01 ETH | 前三名捐赠者（按金额降序） |
| | #2 `0xDD87b63dd...` 捐赠 0.002 ETH | |

---

## 常用命令速查

```bash
# 从 Forge 产物生成 Go 绑定
jq '.abi' out/Contract.sol/Contract.json > Contract.abi
jq -r '.bytecode.object' out/Contract.sol/Contract.json > Contract.bin
abigen --abi Contract.abi --bin Contract.bin --pkg main --type Contract --out Contract.go
rm -f Contract.abi Contract.bin

# 编译 & 运行
go build -o task2.exe .
ETH_RPC_URL="<rpc-url>" CONTRACT_ADDR="<address>" go run .
```
