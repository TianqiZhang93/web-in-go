#!/usr/bin/env bash

cd "$(cd "$(dirname "$0")";pwd)"

env=$1

bin=go_web

set -e
export GO111MODULE=on
export GOPROXY=https://goproxy.cn


function go_test() {
    mkdir -p coverage
    go test -coverprofile=coverage/cover.out ./...
    if [[ $? -ne 0 ]]; then
        echo "unit test failed!!!"
        exit 1
    fi
    go tool cover -html=coverage/cover.out -o coverage/coverage.html
}

echo ["Unit Test"]
go_test

echo
echo [build]
go build -o ./bin/${bin}

rm -rf ./dist ./*.tar.gz
mkdir -p ./dist/${bin}/log ./dist/${bin}/conf

cp -r ./bin ./dist/${bin}/
cp -r ./bin/server.sh ./dist/${bin}/bin/
cp -r ./env_configs/${env}/server.yaml ./dist/${bin}/conf/

echo
echo -e "\033[32m编译成功\033[0m"
echo
