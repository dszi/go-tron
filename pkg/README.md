# go-tron RPC API Documentation

This document lists the RPC API interfaces that have been implemented in go-tron, along with those planned for future releases.

## Implemented Interfaces

- GetAccount
- GetAccountBalance
- GetAccountResource
- CreateTransaction
- BroadcastTransaction
- CreateAccount
- UpdateAccount
- VoteWitnessAccount
- GetBrokerageInfo
- GetRewardInfo
- UpdateBrokerage
- CreateAssetIssue
- ListWitnesses
- CreateWitness
- UpdateWitness
- TransferAsset
- ParticipateAssetIssue
- ListNodes
- GetAssetIssueList
- GetAssetIssueByAccount
- GetAssetIssueByName
- GetAssetIssueById
- GetNowBlock
- GetBlockByNum
- TotalTransaction
- GetTransactionByID
- FreezeBalance
- UnfreezeBalance
- WithdrawBalance
- UnfreezeAsset
- GetNextMaintenanceTime
- GetTransactionInfoByID
- GetBlockByID
- GetMarketOrderByAccount
- GetMarketPairList
- GetMarketOrderListByPair
- GetMarketPriceByPair
- GetMarketOrderById
- GetBurnTrx
- UnfreezeBalanceV2
- WithdrawExpireUnfreeze
- DelegateResource
- UnDelegateResource
- GetTransactionFromPending
- GetTransactionListFromPending
- GetPendingSize
- CancelAllUnfreezeV2
- GetBandwidthPrices
- GetEnergyPrices
- GetMemoFee

## Planned Interfaces

- getTransactionsFromThis
- getTransactionsToThis
- UpdateAsset
- GetPaginatedAssetIssueList
- DeployContract
- TriggerContract
- CreateShieldedTransaction
- GetMerkleTreeVoucherInfo
- ScanNoteByOvk
- GetSpendingKey
- GetExpandedSpendingKey
- GetAkFromAsk
- GetNkFromNsk
- GetIncomingViewingKey
- GetDiversifier
- GetZenPaymentAddress
- GetRcm
- IsSpend
- CreateShieldedTransactionWithoutSpendAuthSig
- GetShieldTransactionHash
- CreateSpendAuthSig
- CreateShieldNullifier
- GetNewShieldedAddress
- CreateShieldedContractParameters
- CreateShieldedContractParametersWithoutAsk
- ScanShieldedTRC20NotesbyIvk
- ScanShieldedTRC20NotesbyOvk
- IsShieldedTRC20ContractNoteSpent
- GetTriggerInputForShieldedTRC20Contract
- MarketSellAsset
- MarketCancelOrder
- GetBlockBalanceTrace

---

