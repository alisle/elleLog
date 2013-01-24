#!/bin/sh
export GOPATH=$(pwd)
go build elleLog

if [ $? -eq 0 ]; then
	mv elleLog deploy/lib/bin

    go build elleLog-StatsServer
    if [ $? -eq 0 ]; then
            mv elleLog-StatsServer deploy/lib/bin
    fi
fi

