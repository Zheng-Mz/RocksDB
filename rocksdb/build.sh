#!/bin/bash

cd code
make static_lib
cp librocksdb.a ../
cp ./include ../include -r
