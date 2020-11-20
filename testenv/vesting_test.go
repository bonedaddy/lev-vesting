package testenv

import (
	"testing"
	"time"

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
	testenv.DoWaitMined(tx, "mint")

	vest, vestC, err := DeployVestingContract(testenv, lev, dateTime, testenv.Auth.From)
	require.NoError(t, err)

	tx, err = tokenC.Approve(testenv.Auth, vest, utils.MaxSupply)
	require.NoError(t, err)
	testenv.DoWaitMined(tx, "approve")

	tx, err = vestC.PrepareVest(testenv.Auth, utils.TeamTokensVested, testenv.Auth.From)
	require.NoError(t, err)
	testenv.DoWaitMined(tx, "prepare vest")

	receiver, err := vestC.Receiver(nil)
	require.NoError(t, err)
	require.Equal(t, receiver, testenv.Auth.From)

	type args struct {
		increaseAmt time.Duration
		doIncrease  bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"1",
			args{((time.Hour * 24) * 31), true},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.args.doIncrease {
				require.NoError(t, testenv.AdjustTime(tt.args.increaseAmt))
				testenv.Commit()
			}

			tx, err = vestC.Release(testenv.Auth)
			require.NoError(t, err)
			testenv.DoWaitMined(tx, "release")
		})
	}
}
