package testenv

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDeployDateTime(t *testing.T) {
	testenv, err := NewBlockchain()
	require.NoError(t, err)
	contract, err := DeployDateTime(testenv)
	require.NoError(t, err)
}
