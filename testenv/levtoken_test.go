package testenv

import (
	"context"
	"testing"

	"github.com/bonedaddy/lev-vesting/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/stretchr/testify/require"
)

func TestDeployLevToken(t *testing.T) {
	testenv, err := NewBlockchain()
	require.NoError(t, err)

	supply := utils.ToWei(125000000.0, 18)

	contract, err := DeployLevToken(testenv, supply)
	require.NoError(t, err)

	tx, err := contract.Mint(testenv.Auth, testenv.Auth.From, supply)
	require.NoError(t, err)

	testenv.Commit()

	_, err = bind.WaitMined(context.Background(), testenv, tx)
	require.NoError(t, err)

	retSupply, err := contract.TotalSupply(nil)
	require.NoError(t, err)
	require.Equal(t, supply, retSupply)

	balance, err := contract.BalanceOf(nil, testenv.Auth.From)
	require.NoError(t, err)
	require.Equal(t, supply, balance)
}
