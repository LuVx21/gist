
```bash
CGO_CFLAGS="-I/opt/homebrew/Cellar/rocksdb/8.9.1/include" \
CGO_LDFLAGS="-L/opt/homebrew/Cellar/rocksdb/8.9.1/lib -lrocksdb -lstdc++ -lm -lz -lbz2 -L/opt/homebrew/Cellar/snappy/1.1.10/lib -L/opt/homebrew/Cellar/lz4/1.9.4/lib -L/opt/homebrew/Cellar/zstd/1.5.5/lib"
go get github.com/tecbot/gorocksdb
```
