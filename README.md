<p align="center">
  <img src="docs/static/go-tron.png" alt="go-tron" width="300">
</p>

<h1 align="center">go-tron</h1>

<p align="center">
<img src="https://img.shields.io/github/go-mod/go-version/dszi/go-tron?style=flat-square" alt="Go Version">
<img src="https://img.shields.io/github/license/dszi/go-tron?style=flat-square" alt="License">
</p>

<p align="center">
  <b>go-tron</b> is a Golang SDK for interacting with the TRON blockchain.
</p>

---

## Features

- **Comprehensive RPC Coverage**: Supports TRON’s gRPC API for account management, transactions, and market data.
- **Optimized for Go**: Provides idiomatic Go interfaces.
- **Secure and Efficient**: Designed for high-performance blockchain applications.

### ⚠️ Alpha Release

---

## Installation

To install go-tron, run:

```bash
go get github.com/dszi/go-tron
```

---

## Quick Start

Here’s a simple example to fetch account details:

```go
package main

import (
    "fmt"
	
    "github.com/dszi/go-tron/pkg"
    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials/insecure"
)

func main() {
    client := pkg.NewGrpcClient("grpc.nile.trongrid.io:50051",
        pkg.WithDialOptions(grpc.WithTransportCredentials(insecure.NewCredentials())),
    )

    if err := client.Start(); err != nil {
        fmt.Println("Failed to start client:", err)
        return
    }
    defer client.Stop()

    addr := "TDS7NjQwQn7iNBN1UxxpsFD5UB3pxB7msL"
    account, err := client.GetAccount(addr)
    if err != nil {
        fmt.Println("Error fetching account:", err)
        return
    }

    fmt.Printf("Account: %+v\n", account)
}
```

---

## Implemented APIs

- `GetAccount`
- `GetAccountBalance`
- `CreateTransaction`
- `BroadcastTransaction`
- `ListWitnesses`
- `GetNowBlock`
- `GetTransactionByID`
- `FreezeBalance`
- `UnfreezeBalance`
- `WithdrawBalance`
- `GetMarketPairList`
- `GetTransactionListFromPending`
- `GetBurnTrx`
- `GetBandwidthPrices`

For a complete list, see the [API documentation](./pkg/README.md).

---

## Planned Features

The following APIs are planned for future implementation:

- `GetMerkleTreeVoucherInfo`
- `ScanNoteByIvk`
- `ScanNoteByOvk`
- `GetZenPaymentAddress`
- `CreateShieldedTransactionWithoutSpendAuthSig`
- `GetShieldTransactionHash`
- `CreateSpendAuthSig`
- `CreateShieldNullifier`

For the full roadmap, visit the [development plan](./pkg/README.md).

---

## Contributing

We welcome contributions! To contribute:

1. Fork this repository.
2. Create a feature branch.
3. Submit a pull request.

---

## License

This project is licensed under the MIT License. See the [LICENSE](./LICENSE) file.

---

## Resources

- [TRON Developer Docs](https://developers.tron.network/)
- [gRPC API Reference](https://developers.tron.network/docs)
- [Nile Testnet](https://nileex.io/)

Build your TRON applications with **go-tron** today🚀!
