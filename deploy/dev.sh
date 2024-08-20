#!/bin/bash

basedir=$(
    cd "$(dirname "$0")" || exit
    pwd
)

echo "${basedir}"

project_path=${basedir}/../
bin=${project_path}/build/bdjuno

#GENESIS_PATH=/Users/lcq/Code/evmos/evmos/build/nodes/node0/evmosd/config/genesis.json

cmd=$1

case ${cmd} in
build)
    echo "===== build ===="

    echo "===== end ===="
    ;;
init)
    echo "===== init ===="

    echo "===== end ===="
    ;;
start)
    echo "===== start ===="

    echo "===== end ===="
    ;;
stop)
    echo "===== stop ===="

    echo "===== end ===="
    ;;
reset)
    echo "===== reset ===="
    echo "Start the postgres and graphql-engine..."
    cd "${project_path}"/deploy || exit
    docker compose -f docker-compose.yaml down
    rm -rf ./data
    docker compose -f docker-compose.yaml up -d

    echo "wait 30s for graphql engine start..."
    for ((i = 20; i > 0; i -= 3)); do
        echo "please wait ${i}s..."
        sleep 3
    done

    echo "Importing the Hasura metadata..."
    cd "$project_path"/hasura || exit
    hasura metadata apply --endpoint http://localhost:8080 --admin-secret mechain

    # echo "rebuild bdjuno..."
    # killall bdjuno

    # cd "${project_path}" || exit
    # make build

    # echo "Initializing the configuration..."
    # cd "$project_path"/deploy || exit
    # cp config.yaml ./data/config.yaml
    # ${bin} parse genesis-file --genesis-file-path ./genesis.json --home ./data

    # echo "run BDjuno...."
    #nohup ${bin} start --home ./data >./bdjuno.log 2>&1 &
    #    ${bin} start --home ./data

    echo "===== end ===="
    ;;
*)
    echo "Usage: dev.sh build | init | reset | start | stop"
    ;;
esac
