package utils

var (
	// MaxSupply of LEV
	MaxSupply = ToWei(10000000.0, 18)
	// TeamTokens are the amount of tokens allocated to team
	TeamTokens = ToWei(7500.0, 18)
	// TeamTokensVested are the amount of tokens team is vesting
	TeamTokensVested = ToWei(5625.0, 18)
	// VestingPeriodMonths denotes the total vesting time
	VestingPeriodMonths = 12
	// VestingCycleMonths denotes the cyle at which individual tranches are released
	VestingCycleMonths = 1
)
