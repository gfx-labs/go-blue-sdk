package morphosdk

// TransactionType represents the type of transaction
type TransactionType string

const (
	TransactionTypeSupply             TransactionType = "Supply"
	TransactionTypeSupplyCollateral   TransactionType = "SupplyCollateral"
	TransactionTypeWithdraw           TransactionType = "Withdraw"
	TransactionTypeWithdrawCollateral TransactionType = "WithdrawCollateral"
	TransactionTypeBorrow             TransactionType = "Borrow"
	TransactionTypeRepay              TransactionType = "Repay"
)

// CapacityLimitReason represents the reason for a capacity limit
type CapacityLimitReason string

const (
	CapacityLimitReasonLiquidityLimit   CapacityLimitReason = "liquidityLimit"
	CapacityLimitReasonUtilizationLimit CapacityLimitReason = "utilizationLimit"
	CapacityLimitReasonSupplyCapLimit   CapacityLimitReason = "supplyCapLimit"
	CapacityLimitReasonBorrowCapLimit   CapacityLimitReason = "borrowCapLimit"

	// V2 vault capacity limits
	CapacityLimitReasonVaultV2AbsoluteCapLimit CapacityLimitReason = "vaultV2AbsoluteCapLimit"
	CapacityLimitReasonVaultV2RelativeCapLimit CapacityLimitReason = "vaultV2RelativeCapLimit"
)
