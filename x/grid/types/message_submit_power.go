package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSubmitPower = "submit_power"

var _ sdk.Msg = &MsgSubmitPower{}

func NewMsgSubmitPower(creator string, watt string) *MsgSubmitPower {
	return &MsgSubmitPower{
		Creator: creator,
		Watt:    watt,
	}
}

func (msg *MsgSubmitPower) Route() string {
	return RouterKey
}

func (msg *MsgSubmitPower) Type() string {
	return TypeMsgSubmitPower
}

func (msg *MsgSubmitPower) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSubmitPower) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSubmitPower) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
