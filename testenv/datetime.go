package testenv

import (
	"github.com/bonedaddy/lev-vesting/bindings/datetime"
	"github.com/ethereum/go-ethereum/common"
)

func DeployDateTime(testenv *Testenv) (common.Address, *datetime.Datetime, error) {
	_, tx, contract, err := datetime.DeployDatetime(testenv.Auth, testenv)
	if err != nil {
		return common.Address{}, nil, err
	}
	addr, err := testenv.DoWaitDeployed(tx)

	return addr, contract, nil
}
