package types_test

import (
	"testing"

	"github.com/crypto-org-chain/chain-main/v2/x/nft/types"
	"github.com/stretchr/testify/require"
)

func TestValidateDenomId(t *testing.T) {
	// valid denom id without hyphens
	err := types.ValidateDenomID("abc")
	require.NoError(t, err)

	// valid denom id with hyphens
	err = types.ValidateDenomID("abc-def")
	require.NoError(t, err)

	// invalid denom id with underscore
	err = types.ValidateDenomID("abc_def")
	require.Error(t, err)

	// too short denom id
	err = types.ValidateDenomID("ab")
	require.Error(t, err)

	// invalid denom id that starts with a digit
	err = types.ValidateDenomID("1abc")
	require.Error(t, err)

	// too long denom id
	err = types.ValidateDenomID("token-677578fa480df7daa517233a6a8ac2ec5a5b88eec9a32e1764574bf97c140ffd677578fa480df7daa517233a6a8ac2ec5a5b88eec9a32e1764574bf97c140ffd")
	require.Error(t, err)
}
