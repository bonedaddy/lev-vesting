package testenv

import (
	"context"

	"github.com/bonedaddy/lev-vesting/bindings/datetime"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func DeployDateTime(testenv *Testenv) (*datetime.Datetime, error) {
	_, tx, contract, err := datetime.DeployDatetime(testenv.Auth, testenv)
	if err != nil {
		return nil, err
	}
	testenv.Commit()
	_, err = bind.WaitDeployed(context.Background(), testenv, tx)
	if err != nil {
		return nil, err
	}
	return contract, nil
}
