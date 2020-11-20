package testenv

import (
	"math/big"

	"github.com/bonedaddy/lev-vesting/bindings/levtoken"
	"github.com/ethereum/go-ethereum/common"
)

// DeployLevToken is used to deploy the leverage token
func DeployLevToken(testenv *Testenv, supply *big.Int) (common.Address, *levtoken.Levtoken, error) {
	_, tx, contract, err := levtoken.DeployLevtoken(testenv.Auth, testenv, "LeverageToken", "LEV", supply)
	if err != nil {
		return common.Address{}, nil, err
	}
	addr, err := testenv.DoWaitDeployed(tx)
	return addr, contract, nil
}
