# royalty-calc-go

Small Go helpers for estimating oil and gas royalty payments and decimal interest for mineral owners.

Published by [Texas Royalty Brokers](https://texasroyaltybrokers.com/), which helps mineral owners across Texas understand how policy, markets, and energy trends affect their land, their royalties, and their communities.

## Usage

    di := royalty.DecimalInterest(10, 640, 0.25)
    pay := royalty.MonthlyPayment(100000, di, 50)

Website: https://texasroyaltybrokers.com/
