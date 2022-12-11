#!/bin/sh

cd ex1
go build -o $1
echo Built ex1.
cd ../ex2
go build -o $1
echo Built ex2.
cd ..
