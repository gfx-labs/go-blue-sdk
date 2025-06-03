package morphoblue

import "errors"

// Morpho Blue error definitions based on ErrorsLib.sol
var (
	// Ownership errors
	ErrorNotOwner = errors.New("not owner")

	// Market configuration errors
	ErrorMaxLltvExceeded       = errors.New("max LLTV exceeded")
	ErrorMaxFeeExceeded        = errors.New("max fee exceeded")
	ErrorAlreadySet            = errors.New("already set")
	ErrorIrmNotEnabled         = errors.New("IRM not enabled")
	ErrorLltvNotEnabled        = errors.New("LLTV not enabled")
	ErrorMarketAlreadyCreated  = errors.New("market already created")
	ErrorNoCode                = errors.New("no code")
	ErrorMarketNotCreated      = errors.New("market not created")

	// Input validation errors
	ErrorInconsistentInput = errors.New("inconsistent input")
	ErrorZeroAssets        = errors.New("zero assets")
	ErrorZeroAddress       = errors.New("zero address")

	// Authorization errors
	ErrorUnauthorized = errors.New("unauthorized")

	// Position errors
	ErrorInsufficientCollateral = errors.New("insufficient collateral")
	ErrorInsufficientLiquidity  = errors.New("insufficient liquidity")
	ErrorHealthyPosition        = errors.New("position is healthy")

	// Signature errors
	ErrorInvalidSignature = errors.New("invalid signature")
	ErrorSignatureExpired = errors.New("signature expired")
	ErrorInvalidNonce     = errors.New("invalid nonce")

	// Transfer errors
	ErrorTransferReverted        = errors.New("transfer reverted")
	ErrorTransferReturnedFalse   = errors.New("transfer returned false")
	ErrorTransferFromReverted    = errors.New("transferFrom reverted")
	ErrorTransferFromReturnedFalse = errors.New("transferFrom returned false")

	// Overflow errors
	ErrorMaxUint128Exceeded = errors.New("max uint128 exceeded")

	// Math errors
	ErrorDivideByZero    = errors.New("divide by zero")
	ErrorUint256Overflow = errors.New("uint256 overflow")
)