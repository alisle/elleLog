#!/bin/sh
export GOPATH=$(pwd)
go build elleLog
if [ $? -eq 0 ]; then
	mv elleLog deploy/lib/bin
fi
