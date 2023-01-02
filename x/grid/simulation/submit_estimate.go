package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"grid/x/grid/keeper"
	"grid/x/grid/types"
)

func SimulateMsgSubmitEstimate(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgSubmitEstimate{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the SubmitEstimate simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "SubmitEstimate simulation not implemented"), nil, nil
	}
}
