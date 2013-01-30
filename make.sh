#!/bin/sh
export PATH=/usr/local/go/bin:$PATH
export GOPATH=$(pwd):$GOPATH
go build elleLog

if [ $? -eq 0 ]; then
	mv elleLog deploy/lib/bin

    go build elleLog-StatsServer
    if [ $? -eq 0 ]; then
            mv elleLog-StatsServer deploy/lib/bin
    fi
fi

