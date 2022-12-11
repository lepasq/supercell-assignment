#!/bin/sh

echo -e "======= TESTING EXERCISE 1 =======\n"
cd ex1
go test -v .

echo -e "\n======= TESTING EXERCISE 2 =======\n"
cd ../ex2
go test -v .
cd ..
