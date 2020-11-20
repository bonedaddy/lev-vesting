package testenv

import (
	"context"
	"log"
	"math/big"

	"github.com/bonedaddy/lev-vesting/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
)

type Testenv struct {
	Auth *bind.TransactOpts
	*backends.SimulatedBackend
}

// NewBlockchain is used to generate a simulated blockchain
func NewBlockchain() (*Testenv, error) {
	auth, _, err := utils.NewAccount()
	if err != nil {
		return nil, err
	}
	// https://medium.com/coinmonks/unit-testing-solidity-contracts-on-ethereum-with-go-3cc924091281
	balance := new(big.Int).Mul(big.NewInt(999999999999999), big.NewInt(999999999999999))
	gAlloc := map[common.Address]core.GenesisAccount{
		auth.From: {Balance: balance},
	}
	sim := backends.NewSimulatedBackend(gAlloc, 8000000)
	return &Testenv{Auth: auth, SimulatedBackend: sim}, nil
}

func (t *Testenv) DoWaitMined(tx *types.Transaction, printArgs ...string) error {
	t.Commit()
	rcpt, err := bind.WaitMined(context.Background(), t, tx)
	log.Println("gas used by transaction: ", rcpt.CumulativeGasUsed, printArgs)
	return err
}

func (t *Testenv) DoWaitDeployed(tx *types.Transaction) (common.Address, error) {
	t.Commit()
	return bind.WaitDeployed(context.Background(), t, tx)
}
