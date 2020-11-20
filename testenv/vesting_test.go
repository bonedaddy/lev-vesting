package testenv

import (
	"testing"

	"github.com/bonedaddy/lev-vesting/utils"
	"github.com/stretchr/testify/require"
)

func TestDeployVestingContract(t *testing.T) {
	testenv, err := NewBlockchain()
	require.NoError(t, err)

	dateTime, _, err := DeployDateTime(testenv)
	require.NoError(t, err)

	lev, tokenC, err := DeployLevToken(testenv, utils.MaxSupply)
	require.NoError(t, err)

	tx, err := tokenC.Mint(testenv.Auth, testenv.Auth.From, utils.MaxSupply)
	require.NoError(t, err)

	testenv.DoWaitMined(tx)

	vest, vestC, err := DeployVestingContract(testenv, lev, dateTime, testenv.Auth.From)
	require.NoError(t, err)

	tx, err = tokenC.Approve(testenv.Auth, vest, utils.MaxSupply)
	require.NoError(t, err)
	testenv.DoWaitMined(tx)

	tx, err = vestC.PrepareVest(testenv.Auth, utils.TeamTokensVested, testenv.Auth.From)
	require.NoError(t, err)
	testenv.DoWaitMined(tx)

}
