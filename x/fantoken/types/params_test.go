package types

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/bitsongofficial/go-bitsong/app"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestValidateParams(t *testing.T) {
	tests := []struct {
		testCase string
		Params
		expectPass bool
	}{
		{
			"Minimum value",
			Params{
				IssuePrice: sdk.NewCoin(app.BondDenom, sdk.ZeroInt()),
			},
			true,
		}, {
			"Maximum value",
			Params{
				IssuePrice: sdk.NewCoin(app.BondDenom, sdk.NewInt(math.MaxInt64)),
			},
			true,
		}, {
			"IssuePrice is negative",
			Params{
				IssuePrice: sdk.Coin{Denom: app.BondDenom, Amount: sdk.NewInt(-1)},
			},
			false,
		},
	}

	for _, tc := range tests {
		if tc.expectPass {
			require.Nil(t, ValidateParams(tc.Params), "test: %v", tc.testCase)
		} else {
			require.NotNil(t, ValidateParams(tc.Params), "test: %v", tc.testCase)
		}
	}
}
