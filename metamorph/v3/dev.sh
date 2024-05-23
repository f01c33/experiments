#!/usr/bin/env sh
rm *_genesis.bin
# go-bindata data/*
go run main.go
gscript compile --enable-logging eval-enc-js.gs.js;
mv $TEMP/*_genesis.bin .