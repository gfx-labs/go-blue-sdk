package morphoblue

import "github.com/holiman/uint256"

type adaptiveIRM struct {
	// constants

	CURVE_STEEPNESS    *uint256.Int
	TARGET_UTILIZATION *uint256.Int

	INITIAL_RATE_AT_TARGET *uint256.Int
	ADJUSTMENT_SPEED       *uint256.Int
	MIN_RATE_AT_TARGET     *uint256.Int
	MAX_RATE_AT_TARGET     *uint256.Int

	LN_2_INT         *uint256.Int
	LN_WEI_INT       *uint256.Int
	WEXP_UPPER_BOUND *uint256.Int
	WEXP_UPPER_VALUE *uint256.Int
}

var AdaptiveIRM = &adaptiveIRM{
	CURVE_STEEPNESS:    uint256.MustFromDecimal("4000000000000000000"),
	TARGET_UTILIZATION: uint256.MustFromDecimal("900000000000000000"),

	INITIAL_RATE_AT_TARGET: new(uint256.Int).Div(uint256.MustFromDecimal("4000000000000000"), SECONDS_PER_YEAR),
	ADJUSTMENT_SPEED:       new(uint256.Int).Div(uint256.MustFromDecimal("50000000000000000000"), SECONDS_PER_YEAR),
	MIN_RATE_AT_TARGET:     new(uint256.Int).Div(uint256.MustFromDecimal("1000000000000000"), SECONDS_PER_YEAR),
	MAX_RATE_AT_TARGET:     new(uint256.Int).Div(uint256.MustFromDecimal("2000000000000000000"), SECONDS_PER_YEAR),

	LN_2_INT:         uint256.MustFromDecimal("693147180559945309"),
	LN_WEI_INT:       (&uint256.Int{}).Neg(uint256.MustFromDecimal("41446531673892822312")),
	WEXP_UPPER_BOUND: uint256.MustFromDecimal("93859467695000404319"),
	WEXP_UPPER_VALUE: uint256.MustFromDecimal("57716089161558943949701069502944508345128422502756744429568"),
}

func (irm *adaptiveIRM) ExpLibWExp(x *uint256.Int) *uint256.Int {
	if x.Slt(irm.LN_WEI_INT) {
		return uint256.NewInt(0)
	}
	if !x.Slt(irm.WEXP_UPPER_BOUND) {
		return new(uint256.Int).Set(irm.WEXP_UPPER_VALUE)
	}

	// roundingAdjustment = (x < 0) ? -(LN_2_INT / 2) : (LN_2_INT / 2);
	roundingAdj := new(uint256.Int).SDiv(irm.LN_2_INT, uint256.NewInt(2))
	if x.Sign() < 0 {
		roundingAdj.Neg(roundingAdj)
	}

	// q = (x + roundingAdjustment) / LN_2_INT;
	q := new(uint256.Int).Add(x, roundingAdj)
	q.SDiv(q, irm.LN_2_INT)

	// int256 r = x - q * LN_2_INT;
	r := new(uint256.Int).Mul(q, irm.LN_2_INT)
	r.Sub(x, r)

	rSquared := new(uint256.Int).Mul(r, r)

	expR := new(uint256.Int).SDiv(rSquared, WAD)
	expR.SDiv(expR, uint256.NewInt(2))
	expR.Add(expR, r)
	expR.Add(expR, WAD)

	if q.Sign() == 0 {
		expR.Lsh(expR, uint(q.Uint64()))
		return expR
	}
	q.Neg(q)
	expR.Rsh(expR, uint(q.Uint64()))
	return expR
}

// simplified, since we only need the start rate, elapsed time, and utilization. we dont need the market id or the other market parameters
// https://github.com/morpho-org/morpho-blue-irm/blob/main/src/adaptive-curve-irm/AdaptiveCurveIrm.sol#L76C47-L76C53
func (irm *adaptiveIRM) GetBorrowRate(utilization, startRateAtTarget, elapsed *uint256.Int) (*uint256.Int, *uint256.Int) {
	//   int256 errNormFactor = utilization > ConstantsLib.TARGET_UTILIZATION
	//        ? WAD - ConstantsLib.TARGET_UTILIZATION
	//        : ConstantsLib.TARGET_UTILIZATION;
	errNormFactor := new(uint256.Int)
	if utilization.Sgt(irm.TARGET_UTILIZATION) {
		errNormFactor.Sub(WAD, irm.TARGET_UTILIZATION)
	} else {
		errNormFactor.Set(irm.TARGET_UTILIZATION)
	}

	//  int256 err = (utilization - ConstantsLib.TARGET_UTILIZATION).wDivToZero(errNormFactor);
	ERR := new(uint256.Int).Sub(utilization, irm.TARGET_UTILIZATION)
	ERR = WadDivToZero(ERR, ERR, errNormFactor)

	avgRateAtTarget := new(uint256.Int)
	endRateAtTarget := new(uint256.Int)
	if startRateAtTarget.IsZero() {
		avgRateAtTarget.Set(irm.INITIAL_RATE_AT_TARGET)
		endRateAtTarget.Set(irm.INITIAL_RATE_AT_TARGET)
	} else {
		// int256 speed = ConstantsLib.ADJUSTMENT_SPEED.wMulToZero(err);
		speed := WadMulToZero(
			new(uint256.Int),
			irm.ADJUSTMENT_SPEED,
			ERR,
		)

		// elapsed is seconds from now to the last market update.
		linearAdaptation := new(uint256.Int).Mul(speed, elapsed)

		if linearAdaptation.IsZero() {
			avgRateAtTarget.Set(startRateAtTarget)
			endRateAtTarget.Set(startRateAtTarget)
		} else {

			// https://github.com/morpho-org/morpho-blue-irm/blob/main/src/adaptive-curve-irm/AdaptiveCurveIrm.sol#L109

			// function _newRateAtTarget(int256 startRateAtTarget, int256 linearAdaptation) private pure returns (int256) {
			//     // Non negative because MIN_RATE_AT_TARGET > 0.
			//     return startRateAtTarget.wMulToZero(ExpLib.wExp(linearAdaptation)).bound(
			//         ConstantsLib.MIN_RATE_AT_TARGET, ConstantsLib.MAX_RATE_AT_TARGET
			//     );
			// }
			_newRateAtTarget := func(srat, adaptation *uint256.Int) *uint256.Int {
				expResult := irm.ExpLibWExp(adaptation)
				// WMulToZero ->  return (x * y) / WAD_INT;
				result := WadMulToZero(new(uint256.Int), srat, expResult)
				// bound between min and max rate. make sure to use signed operations
				if result.Slt(irm.MIN_RATE_AT_TARGET) {
					return new(uint256.Int).Set(irm.MIN_RATE_AT_TARGET)
				}
				if result.Sgt(irm.MAX_RATE_AT_TARGET) {
					return new(uint256.Int).Set(irm.MAX_RATE_AT_TARGET)
				}
				return result
			}
			// endRateAtTarget = _newRateAtTarget(startRateAtTarget, linearAdaptation);
			endRateAtTarget = _newRateAtTarget(startRateAtTarget, linearAdaptation)

			// int256 midRateAtTarget = _newRateAtTarget(startRateAtTarget, linearAdaptation / 2);
			halfAdaptation := new(uint256.Int).SDiv(linearAdaptation, uint256.NewInt(2))
			midRate := _newRateAtTarget(startRateAtTarget, halfAdaptation)
			midRate.Mul(midRate, uint256.NewInt(2))

			// avgRateAtTarget = (startRateAtTarget + endRateAtTarget + 2 * midRateAtTarget) / 4;
			avgRateAtTarget.Add(startRateAtTarget, endRateAtTarget)
			avgRateAtTarget.Add(avgRateAtTarget, midRate)
			avgRateAtTarget.SDiv(avgRateAtTarget, uint256.NewInt(4))
		}
	}
	//  return (uint256(_curve(avgRateAtTarget, err)), endRateAtTarget);

	//function _curve(int256 _rateAtTarget, int256 err) private pure returns (int256) {
	//    // Non negative because 1 - 1/C >= 0, C - 1 >= 0.
	//    // Non negative if _rateAtTarget >= 0 because if err < 0, coeff <= 1.
	//}

	//    int256 coeff = err < 0 ? WAD - WAD.wDivToZero(ConstantsLib.CURVE_STEEPNESS) : ConstantsLib.CURVE_STEEPNESS - WAD;
	coeff := new(uint256.Int)
	if ERR.Sign() < 0 {
		// TODO: this is a constant that we can precomputer, right?
		wadDiv := WadDivToZero(
			new(uint256.Int), WAD, irm.CURVE_STEEPNESS,
		)
		coeff.Sub(WAD, wadDiv)
	} else {
		coeff.Sub(irm.CURVE_STEEPNESS, WAD)
	}

	// return (coeff.wMulToZero(err) + WAD).wMulToZero(int256(_rateAtTarget));
	borrowRate := WadMulToZero(
		new(uint256.Int),
		coeff,
		ERR,
	)
	borrowRate.Add(borrowRate, WAD)
	borrowRate = WadMulToZero(borrowRate, borrowRate, avgRateAtTarget)

	return borrowRate, endRateAtTarget
}
