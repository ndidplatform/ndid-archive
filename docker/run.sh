#!/bin/bash
MODE=$1

if [ "$MODE" = "idp" ]
then
	echo "Runnig IdP mode"
	/usr/bin/tendermint init
	sed -i "s|127.0.0.1:46658|ndid-abci-idp:46000|g" /tendermint/config/config.toml
	sed -i "s|46657|45000|g" /tendermint/config/config.toml
	sed -i "s|46656|47000|g" /tendermint/config/config.toml
	/usr/bin/tendermint node --consensus.create_empty_blocks=false
elif [ "$MODE" = "rp" ]
then
	echo "Running RP mode"
	/usr/bin/tendermint init
	sed -i "s|127.0.0.1:46658|ndid-abci-rp:46001|g" /tendermint/config/config.toml
	sed -i "s|46657|45001|g" /tendermint/config/config.toml
	sed -i "s|46656|47001|g" /tendermint/config/config.toml
	/usr/bin/tendermint node --consensus.create_empty_blocks=false
else
	echo "Unknown option."
	echo "Usage:"
	echo "$0 idp"
	echo "or"
	echo "$0 rp"
fi

