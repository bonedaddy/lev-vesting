package testenv

import (
	"math/big"
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

	tx, err = vestC.Prepare(testenv.Auth, utils.TeamTokensVested, testenv.Auth.From)
	require.NoError(t, err)
	testenv.DoWaitMined(tx, "prepare vest")

	receiver, err := vestC.Receiver(nil)
	require.NoError(t, err)
	require.Equal(t, receiver, testenv.Auth.From)

	// run all 12 vest cycles testing every 31 days
	type args struct {
		increaseAmt time.Duration
		doIncrease  bool
	}
	tests := []struct {
		name      string
		args      args
		wantErr   bool
		wantCycle *big.Int
	}{
		{
			"Cycle-1",
			args{((time.Hour * 24) * 31), true},
			false,
			big.NewInt(1),
		},
		{
			"Cycle-2",
			args{((time.Hour * 24) * 31), true},
			false,
			big.NewInt(2),
		},
		{
			"Cycle-3",
			args{((time.Hour * 24) * 31), true},
			false,
			big.NewInt(3),
		},
		{
			"Cycle-4",
			args{((time.Hour * 24) * 31), true},
			false,
			big.NewInt(4),
		},
		{
			"Cycle-5",
			args{((time.Hour * 24) * 31), true},
			false,
			big.NewInt(5),
		},
		{
			"Cycle-6",
			args{((time.Hour * 24) * 31), true},
			false,
			big.NewInt(6),
		},
		{
			"Cycle-7",
			args{((time.Hour * 24) * 31), true},
			false,
			big.NewInt(7),
		},
		{
			"Cycle-8",
			args{((time.Hour * 24) * 31), true},
			false,
			big.NewInt(8),
		},
		{
			"Cycle-9",
			args{((time.Hour * 24) * 31), true},
			false,
			big.NewInt(9),
		},
		{
			"Cycle-10",
			args{((time.Hour * 24) * 31), true},
			false,
			big.NewInt(10),
		},
		{
			"Cycle-11",
			args{((time.Hour * 24) * 31), true},
			false,
			big.NewInt(11),
		},
		{
			"Cycle-12",
			args{((time.Hour * 24) * 31), true},
			false,
			big.NewInt(12),
		},
		{
			"Error Test (No More Cycles)",
			args{((time.Hour * 24) * 31), false},
			true,
			big.NewInt(12),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			cycle, err := vestC.CurrentCycle(nil)
			require.NoError(t, err)
			require.Equal(t, cycle, tt.wantCycle)

			if !tt.wantErr {
				released, err := vestC.IsReleased(nil, cycle)
				require.NoError(t, err)
				require.False(t, released)
			}

			if tt.args.doIncrease {
				require.NoError(t, testenv.AdjustTime(tt.args.increaseAmt))
				testenv.Commit()
			}

			tx, err := vestC.Release(testenv.Auth)
			if (err != nil) != tt.wantErr {
				t.Errorf("Release() err %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tx != nil {
				testenv.DoWaitMined(tx, "release")
			}

			if !tt.wantErr {
				released, err := vestC.IsReleased(nil, cycle)
				require.NoError(t, err)
				require.True(t, released)
			}

		})
	}
}
