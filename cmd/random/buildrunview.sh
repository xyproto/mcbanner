#!/bin/sh
go get -d
go build
./compare > out.png
eog out.png
