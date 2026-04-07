package morphosdk

// TransactionType represents the type of transaction
type TransactionType string

const (
	TransactionTypeSupply             TransactionType = "Supply"
	TransactionTypeSupplyCollateral   TransactionType = "Supply Collateral"
	TransactionTypeWithdraw           TransactionType = "Withdraw"
	TransactionTypeWithdrawCollateral TransactionType = "Withdraw Collateral"
	TransactionTypeBorrow             TransactionType = "Borrow"
	TransactionTypeRepay              TransactionType = "Repay"
)

// CapacityLimitReason represents the reason for a capacity limit
type CapacityLimitReason string

const (
	CapacityLimitReasonLiquidity  CapacityLimitReason = "Liquidity"
	CapacityLimitReasonBalance    CapacityLimitReason = "Balance"
	CapacityLimitReasonPosition   CapacityLimitReason = "Position"
	CapacityLimitReasonCollateral CapacityLimitReason = "Collateral"
	CapacityLimitReasonCap        CapacityLimitReason = "Cap"

	// V2 vault capacity limits
	CapacityLimitReasonVaultV2AbsoluteCap CapacityLimitReason = "VaultV2_AbsoluteCap"
	CapacityLimitReasonVaultV2RelativeCap CapacityLimitReason = "VaultV2_RelativeCap"
)
