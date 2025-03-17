# go-tron RPC API Documentation

This document lists the RPC API interfaces that have been implemented in go-tron, along with those planned for future releases.

## Implemented Interfaces

### Account Management

- GetAccount
- GetAccountBalance
- GetAccountResource
- CreateAccount
- UpdateAccount
- GetRewardInfo

### Transactions

- CreateTransaction
- BroadcastTransaction
- GetTransactionByID
- GetTransactionInfoByID
- GetTransactionFromPending
- GetTransactionListFromPending
- TotalTransaction

### Resource Management

- FreezeBalance
- UnfreezeBalance
- WithdrawBalance
- UnfreezeAsset
- UnfreezeBalanceV2
- WithdrawExpireUnfreeze
- DelegateResource
- UnDelegateResource
- CancelAllUnfreezeV2

### Witnesses Management

- VoteWitnessAccount
- ListWitnesses
- CreateWitness
- UpdateWitness
- GetBrokerageInfo
- UpdateBrokerage

### Asset Management

- CreateAssetIssue
- GetAssetIssueList
- GetPaginatedAssetIssueList
- GetAssetIssueByAccount
- GetAssetIssueByName
- GetAssetIssueById
- TransferAsset
- ParticipateAssetIssue
- UpdateAsset

### Block Management

- GetNowBlock
- GetBlockByNum
- GetBlockByID
- GetNextMaintenanceTime

### Market Management

- GetMarketOrderByAccount
- GetMarketPairList
- GetMarketOrderListByPair
- GetMarketPriceByPair
- GetMarketOrderById
- GetBurnTrx

### Smart Contracts

- DeployContract
- TriggerContract

### Shielded & Privacy

- GetSpendingKey
- GetExpandedSpendingKey
- GetAkFromAsk
- GetNkFromNsk
- GetIncomingViewingKey
- GetDiversifier
- GetRcm
- GetNewShieldedAddress

### Network

- ListNodes
- GetPendingSize
- GetBandwidthPrices
- GetEnergyPrices
- GetMemoFee

## Planned Interfaces


- CreateShieldedTransaction
- GetMerkleTreeVoucherInfo
- ScanNoteByIvk
- ScanNoteByOvk
- GetZenPaymentAddress
- IsSpend
- CreateShieldedTransactionWithoutSpendAuthSig
- GetShieldTransactionHash
- CreateSpendAuthSig
- CreateShieldNullifier
- CreateShieldedContractParameters
- CreateShieldedContractParametersWithoutAsk
- ScanShieldedTRC20NotesbyIvk
- ScanShieldedTRC20NotesbyOvk
- IsShieldedTRC20ContractNoteSpent
- GetTriggerInputForShieldedTRC20Contract
- MarketSellAsset
- MarketCancelOrder
- GetBlockBalanceTrace