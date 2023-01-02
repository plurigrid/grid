package grid

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"grid/testutil/sample"
	gridsimulation "grid/x/grid/simulation"
	"grid/x/grid/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = gridsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgSubmitPower = "op_weight_msg_submit_power"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSubmitPower int = 100

	opWeightMsgSubmitEstimate = "op_weight_msg_submit_estimate"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSubmitEstimate int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	gridGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&gridGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgSubmitPower int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSubmitPower, &weightMsgSubmitPower, nil,
		func(_ *rand.Rand) {
			weightMsgSubmitPower = defaultWeightMsgSubmitPower
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSubmitPower,
		gridsimulation.SimulateMsgSubmitPower(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSubmitEstimate int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSubmitEstimate, &weightMsgSubmitEstimate, nil,
		func(_ *rand.Rand) {
			weightMsgSubmitEstimate = defaultWeightMsgSubmitEstimate
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSubmitEstimate,
		gridsimulation.SimulateMsgSubmitEstimate(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
