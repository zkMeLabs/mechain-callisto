#!/bin/bash
basedir=$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" && pwd)
project_path=$(git rev-parse --show-toplevel)
bin_name=bdjuno
bin=${project_path}/build/${bin_name}

function stop() {
    docker compose -f "${basedir}"/docker-compose.yaml down
    rm -rf "${basedir}"/data
    rm -rf "${basedir}"/bdjuno.log
    ps -ef | grep "${bin_name}" | awk '{print $2}' | xargs kill
}

function start() {
    docker compose -f "${basedir}"/docker-compose.yaml up -d
    echo "wait 30s for graphql engine start..."
    for ((i = 10; i > 0; i -= 3)); do
        echo "please wait ${i}s..."
        sleep 3
    done

    echo "Importing the Hasura metadata..."
    hasura metadata apply --project "${project_path}"/hasura --endpoint http://localhost:8080 --admin-secret mechain

    # echo "Initializing the configuration..."
    # ${bin} --home "${basedir}" parse genesis-file --genesis-file-path ./genesis.json

    echo "run BDjuno...."
    nohup "${bin}" start --home "${basedir}" >"${basedir}"/bdjuno.log 2>&1 &
}

cmd=$1
case ${cmd} in
init)
    echo "===== init ===="

    echo "===== end ===="
    ;;
start)
    echo "===== start ===="
    start
    echo "===== end ===="
    ;;
stop)
    echo "===== stop ===="
    stop
    echo "===== end ===="
    ;;
reset)
    echo "===== reset ===="

    echo "===== stop ===="
    stop
    echo "===== end ===="

    echo "===== start ===="
    start
    echo "===== end ===="

    echo "===== end ===="
    ;;
*)
    echo "Usage: localup.sh init | reset | start | stop"
    ;;
esac
