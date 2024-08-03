#!/bin/bash

BDJUNO_PATH=/Users/lcq/Code/zkme/callisto
GENESIS_PATH=/Users/lcq/Code/evmos/evmos/build/nodes/node0/evmosd/config/genesis.json

killall bdjuno

cd $BDJUNO_PATH
make build

cd $BDJUNO_PATH/build/data
rm -rf ./db_data
docker-compose -f docker-compose.yaml down
docker-compose -f docker-compose.yaml up -d

sleep 30

cd $BDJUNO_PATH/hasura
hasura metadata apply --endpoint http://localhost:8080 --admin-secret 123456

cd $BDJUNO_PATH/build
./bdjuno parse genesis-file --genesis-file-path $GENESIS_PATH --home ./data
#nohup ./bdjuno start --home ./data >./bdjuno.log 2>&1 &
./bdjuno start --home ./data