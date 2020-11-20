package testenv

import (
	"context"
	"log"
	"math/big"

	"github.com/bonedaddy/lev-vesting/bindings/levtoken"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

// DeployLevToken is used to deploy the leverage token
func DeployLevToken(testenv *Testenv, supply *big.Int) (*levtoken.Levtoken, error) {
	_, tx, contract, err := levtoken.DeployLevtoken(testenv.Auth, testenv, "LeverageToken", "LEV", supply)
	if err != nil {
		return nil, err
	}
	testenv.Commit()
	_, err = bind.WaitDeployed(context.Background(), testenv, tx)
	if err != nil {
		log.Println("failed to wait for deploy tx to be mined")
		return nil, err
	}
	return contract, nil
}
