# LEV Vesting

Vesting contracts for the $LEV token

# Vesting Stats

75% of team tokens vested for 12 months, released every month


We will need a contract vesting 75% of team tokens for 12 months, released monthly

# Date Time Contract

The contract uses a newer version of [BokkyPooBahsDateTimeLibrary](https://github.com/bokkypoobah/BokkyPooBahsDateTimeLibrary) than what is currently deployed on mainnet. This means the contract will need to be deployed.

# Gas Costs

The following values are logged from unit tests which use go-ethereum simulated backend

| Action | Cost |
|--------|------|
| Date Time deployment | 1090857 |
| Vesting deployment | 492694 |
| Vest preparation | 455530 |
| Vest release | 77157 |

# Development

## Dependencies

* Linux 64-bit OS
* Go 1.15
* 

## Compilation

# Deployment

Take `flattened/Vesting_Flattened.sol` and use whatever preferred deployment system (truffle, hardhat, etc..)
