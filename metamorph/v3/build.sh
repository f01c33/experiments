#!/usr/bin/env sh

rm *_gscript.exe
# go-bindata data
# mv bindata.go itself
# cd itself
garble build main.go
# mv ransom.exe ../rw.exe
# cd ..
gscript compile eval-enc-js.gs.js;
mv $TEMP/*_gscript.exe .