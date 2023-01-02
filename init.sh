#!/bin/sh

VALIDATOR_NAME=egen0
CHAIN_ID=grid
KEY_NAME=grid-key
CHAINFLAG="--chain-id ${CHAIN_ID}"
TOKEN_AMOUNT="10000000000000000000000000edith"
STAKING_AMOUNT="1000000000edith"

NAMESPACE_ID=$(echo $RANDOM | md5sum | head -c 16; echo;)
echo $NAMESPACE_ID
DA_BLOCK_HEIGHT=$(curl https://rpc.limani.celestia-devops.dev/block | jq -r '.result.block.header.height')
echo $DA_BLOCK_HEIGHT

ignite chain build
gridd tendermint unsafe-reset-all
gridd init $VALIDATOR_NAME --chain-id $CHAIN_ID

gridd keys add $KEY_NAME --keyring-backend test
gridd add-genesis-account $KEY_NAME $TOKEN_AMOUNT --keyring-backend test
gridd gentx $KEY_NAME $STAKING_AMOUNT --chain-id $CHAIN_ID --keyring-backend test
gridd collect-gentxs
gridd start --rollmint.aggregator true --rollmint.da_layer celestia --rollmint.da_config='{"base_url":"http://localhost:26659","timeout":60000000000,"fee":6000,"gas_limit":6000000}' --rollmint.namespace_id $NAMESPACE_ID --rollmint.da_start_height $DA_BLOCK_HEIGHT
