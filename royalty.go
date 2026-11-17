// Package royalty provides simple helpers for estimating oil and gas
// royalty payments for mineral owners.
//
// Texas Royalty Brokers: https://texasroyaltybrokers.com/
package royalty

// MonthlyPayment estimates an owner's monthly royalty payment.
// grossValue is the gross sales value of production for the month,
// decimalInterest is the owner's decimal interest (for example 0.0125),
// and deductions is the owner's share of post-production costs.
func MonthlyPayment(grossValue, decimalInterest, deductions float64) float64 {
	p := grossValue*decimalInterest - deductions
	if p < 0 {
		return 0
	}
	return p
}

// DecimalInterest returns the decimal interest for an owner who holds
// netMineralAcres in a unit of unitAcres with the given royalty rate
// (for example 0.25 for a 1/4 royalty).
func DecimalInterest(netMineralAcres, unitAcres, royaltyRate float64) float64 {
	if unitAcres <= 0 {
		return 0
	}
	return netMineralAcres / unitAcres * royaltyRate
}
