package queriers

import (
	"strings"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/stretchr/testify/require"

	"github.com/Pylons-tech/pylons/x/pylons/handlers"
	"github.com/Pylons-tech/pylons/x/pylons/keep"
	"github.com/Pylons-tech/pylons/x/pylons/types"
)

func TestListRecipe(t *testing.T) {
	tci := keep.SetupTestCoinInput()
	tci.PlnH = handlers.NewMsgServerImpl(tci.PlnK)
	tci.PlnQ = NewQuerierServerImpl(tci.PlnK)

	sender1, _, _, _ := keep.SetupTestAccounts(t, tci, types.NewPylon(1000000), nil, nil, nil)

	// mock cookbook
	cbData := handlers.MockCookbook(tci, sender1)

	handlers.MockPopularRecipe(handlers.Rcp5BlockDelayed5xWoodcoinTo1xChaircoin, tci,
		"recipe0001", cbData.CookbookID, sender1)

	cases := map[string]struct {
		address       string
		desiredError  string
		showError     bool
		desiredRcpCnt int
		firstItemName string
	}{
		"error check when providing invalid address": {
			address:       "invalid_address",
			showError:     true,
			desiredError:  "decoding bech32 failed: invalid index of 1",
			desiredRcpCnt: 0,
		},
		"error check when not providing address": {
			address:       "",
			showError:     false,
			desiredRcpCnt: 1,
			firstItemName: "recipe0001",
		},
		"list recipe successful check": {
			address:       sender1.String(),
			showError:     false,
			desiredError:  "",
			desiredRcpCnt: 1,
			firstItemName: "recipe0001",
		},
	}
	for testName, tc := range cases {
		t.Run(testName, func(t *testing.T) {
			result, err := tci.PlnQ.ListRecipe(
				sdk.WrapSDKContext(tci.Ctx),
				&types.ListRecipeRequest{
					Address: tc.address,
				},
			)
			if tc.showError {
				require.Error(t, err)
				require.True(t, strings.Contains(err.Error(), tc.desiredError))
			} else {
				require.NoError(t, err)
				require.Equal(t, len(result.Recipes), tc.desiredRcpCnt)
				require.True(t, len(result.Recipes) > 0)
				require.Equal(t, result.Recipes[0].Name, tc.firstItemName)
			}
		})
	}
}
