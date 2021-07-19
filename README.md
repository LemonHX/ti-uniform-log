# Unified Log Implementation

is the implementation of ["TiDB Log Format"](https://github.com/tikv/rfcs/blob/master/text/0018-unified-log-format.md)
to be used across all TiDB software, i.e. TiDB, TiKV, and PD.

## Parser

using parser generator `ANTLR4`

## Performance

bench mark platform

```
OS: Red Hat Enterprise Linux 8.4 (Ootpa) x86_64 
CPU: AMD Ryzen 7 3700X (16) @ 3.600GHz 
Memory: 16GiB with 3600mhz
```

data collected form `go test bench` with `~100M` sized log file spend `16sec`

```
goos: linux
goarch: amd64
pkg: log_parser/test
BenchmarkLogParser
BenchmarkLogParser-16    	       1	16599998421 ns/op
PASS
```

```
6.6MiB/s ~= 20000 loc/sec
```
