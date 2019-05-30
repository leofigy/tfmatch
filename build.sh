#!/usr/bin/env sh 
# strip the debugging information
go build -o matcherc -ldflags="-s -w" matcher.go
# optionals
upx --brute matcherc
