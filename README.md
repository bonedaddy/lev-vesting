# LEV Vesting

Vesting contract for the $LEV token

# Vesting Stats

75% of team tokens vested for 12 months, released every month

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
* Make

## Compilation

* `make compile`
  * Compiles the Leverage token (note this is just used for testing)
  * Compiles DateTime contracts
  * Compiles vesting contract
  * Generates JSON ABI files for all contracts in `bin`

## Golang Bindings Update

Updates the golang code bindings for the smart contracts

* `make abigen`

## Flattening 

Flattens the Vesting contract into a single file, suitable for etherscan verification, usage with any solidity development environment, etc...

* `make flatten`

## All Of The Above

* `make all`

# Deployment

Take `flattened/Vesting_Flattened.sol` and use whatever preferred deployment system or development environment (truffle, hardhat, etc..). It should be compiled using Solidity 0.7.3, however you can manually adjust the compiler versiont to 0.7.x as desired.
