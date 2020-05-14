# RocksDB
## RocksDB 库编译
### 更新rocksdb子库
    git submodule update --init
### 编译rocksdb静态库
    cd rocksdb
    ./build.sh

## 编译应用代码
### go
    cd ./go
    go build
### C
    cd ./c
    g++  c_simple_example.c -o c_simple_example -I./../rocksdb/include -L./../rocksdb -lrocksdb -lz -lpthread -lsnappy -lstdc++ -lm -ldl -O3
