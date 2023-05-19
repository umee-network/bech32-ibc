#!/usr/bin/env bash

set -eo pipefail

echo "Generating gogo proto code"
cd proto

buf generate --template buf.gen.gogo.yaml $file

cd ..

cp -r ./github.com/osmosis-labs/bech32-ibc/* ./
rm -rf ./github.com

go mod tidy
