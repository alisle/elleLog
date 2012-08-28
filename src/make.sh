#!/bin/sh
export GOPATH=$(pwd)/..
go build ellelog.go
if [ $? -eq 0 ]; then
	cp ellelog ../deploy/lib/bin
fi
