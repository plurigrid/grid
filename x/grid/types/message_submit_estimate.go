package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSubmitEstimate = "submit_estimate"

var _ sdk.Msg = &MsgSubmitEstimate{}

func NewMsgSubmitEstimate(creator string, watt string) *MsgSubmitEstimate {
	return &MsgSubmitEstimate{
		Creator: creator,
		Watt:    watt,
	}
}

func (msg *MsgSubmitEstimate) Route() string {
	return RouterKey
}

func (msg *MsgSubmitEstimate) Type() string {
	return TypeMsgSubmitEstimate
}

func (msg *MsgSubmitEstimate) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSubmitEstimate) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSubmitEstimate) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
