#!/usr/bin/env sh

go build vaccine.go
mv vaccine ../data/vaccine-linux
env GOOS=windows GOARCH=386 go build vaccine.go
env GOOS=darwin go build vaccine.go
mv vaccine.exe ../data/vaccine-windows.exe
mv vaccine ../data/vaccine-osx