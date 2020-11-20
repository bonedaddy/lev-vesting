package testenv

import (
	"github.com/bonedaddy/lev-vesting/bindings/vesting"
	"github.com/ethereum/go-ethereum/common"
)

// DeployVestingContract is used to deploy the vesting contract
func DeployVestingContract(testenv *Testenv, levToken common.Address, dateTime common.Address, receiver common.Address) (common.Address, *vesting.Vesting, error) {
	_, tx, contract, err := vesting.DeployVesting(testenv.Auth, testenv, levToken, dateTime, testenv.Auth.From)
	if err != nil {
		return common.Address{}, nil, err
	}
	addr, err := testenv.DoWaitDeployed(tx, "vesting")
	return addr, contract, nil
}
