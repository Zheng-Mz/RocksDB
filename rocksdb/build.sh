#!/bin/bash

cd code
make static_lib
cp librocksdb.a ../
cp ./include/rocksdb ../include -r
