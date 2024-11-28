# CGO Overhead Benchmark

This repo contains a benchmark for the overhead of calling Rust functions from Go using FFI.


## Setup
```sh
cargo build --release --manifest-path rust/Cargo.toml
```

## Run Benchmarks

To replicate the results in the README, run the following commands:

```sh
# single-threaded benchmark for direct comparison
go test -bench=. -cpu=1 
```

```sh
cargo bench --manifest-path rust/Cargo.toml
```

## Results

This benchmark measures parsing a JSON string with `1` as an integer to compare with the overhead of calling a Rust function from Go using FFI.

- `BenchmarkLowLevelJSONCParseIntBaseline` use lower level primitives to parse the JSON string
- `BenchmarkJSONMarshalCParseIntBaseline` use `json.Marshal` to parse the JSON string

The result is as follows:

```
BenchmarkLowLevelJSONCParseIntBaseline  18930902                64.25 ns/op
BenchmarkJSONMarshalCParseIntBaseline   13860427                88.82 ns/op
```


Go to Rust FFI
```
BenchmarkAddFFI                         25235221                48.33 ns/op
BenchmarkFibonacciFFI                      17535             68534 ns/op
```

Pure Rust
```
add 1 + 2               time:   [998.68 ps 1.0026 ns 1.0068 ns]
fib 100000              time:   [66.927 µs 67.167 µs 67.391 µs]
```

So simple add diff is +47.3274 ns in this run, this would be the most direct comparison for the overhead of calling a Rust function from Go using FFI.
Fibonacci diff is +1367 ns in this run (but most likely due to variance).

## Conclusion

Comparing to baseline, CGO overhead is pretty negligible.
