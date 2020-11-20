package testenv

import (
	"context"

	"github.com/bonedaddy/lev-vesting/bindings/datetime"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func DeployDateTime(testenv *Testenv) (common.Address, *datetime.Datetime, error) {
	_, tx, contract, err := datetime.DeployDatetime(testenv.Auth, testenv)
	if err != nil {
		return common.Address{}, nil, err
	}
	testenv.Commit()
	addr, err := bind.WaitDeployed(context.Background(), testenv, tx)
	if err != nil {
		return common.Address{}, nil, err
	}
	return addr, contract, nil
}
