# CGO Overhead Benchmark

This repo contains a benchmark for the overhead of calling Rust functions from Go using FFI.


## Setup
```sh
cargo build --release --manifest-path rust/Cargo.toml
```

```sh
go install golang.org/x/perf/cmd/benchstat@latest
```

## Run Benchmarks

To replicate the results in the README, run the following commands:

```sh
# single-threaded benchmark for direct comparison
go test -bench=. -cpu=1 
```

With benchhstat
```sh
benchstat <(go test -bench=. -cpu=1 -count=10)
```

```sh
cargo bench --manifest-path rust/Cargo.toml
```

## Results

This benchmark measures parsing a JSON string with `1` as an integer to compare with the overhead of calling a Rust function from Go using FFI.

- `BenchmarkLowLevelJSONCParseIntBaseline` use lower level primitives to parse the JSON string
- `BenchmarkJSONMarshalCParseIntBaseline` use `json.Marshal` to parse the JSON string

The result is as follows:


Baseline
```
                              │   sec/op    │
LowLevelJSONCParseIntBaseline   64.02n ± 3%
JSONMarshalCParseIntBaseline    91.06n ± 5%
```


Go to Rust FFI
```
                              │   sec/op    │
AddFFI                          48.40n ± 1%
FibonacciFFI                    68.28µ ± 2%
```

Pure Rust
```
add 1 + 2               time:   [978.88 ps 987.08 ps 1.0002 ns]
Found 1 outliers among 100 measurements (1.00%)
  1 (1.00%) high severe

fib 100000              time:   [65.621 µs 66.162 µs 67.067 µs]
Found 4 outliers among 100 measurements (4.00%)
  1 (1.00%) low mild
  2 (2.00%) high mild
  1 (1.00%) high severe
```

If we were to write in the same style as go results:

```
                              │   sec/op    │
AddFFI                          988.72p ± 1%
FibonacciFFI                    66.162μ ± 1%
```


So simple add diff is +47.41128 ns, this would be the most direct comparison for the overhead of calling a Rust function from Go using FFI.

To compare the `FibonacciFFI` benchmark (from Go to Rust via FFI) with the pure Rust `fib 100000` benchmark, we considered their respective mean execution times and uncertainties. Here's how they compare:

- **FibonacciFFI:** `68.28µ ± 2%`, giving a range of `[66.9144, 69.6456]`.
- **fib 100000:** `66.162µ ± 1%`, giving a range of `[65.50038, 66.82362]`.

The ranges indicate no overlap, suggesting that `FibonacciFFI` consistently takes longer than the pure Rust `fib 100000`. To quantify the significance of this difference, we calculated the **z-score**:
\[
z = \frac{\text{Difference between means}}{\sqrt{\text{Variance 1} + \text{Variance 2}}}
\]
The result was `z ≈ 1.40`, which corresponds to a `p-value ≈ 16%`. This means there is a 16% probability that the observed difference is due to random variation, which is not statistically significant at common confidence thresholds (e.g., 95%).

While the `FibonacciFFI` benchmark exhibits slightly higher execution times compared to `fib 100000`, the difference is relatively minor and likely reflects the overhead of calling Rust functions from Go using FFI in which it's not statistically significant.


## Build & Deployment process

To address build and deployment process, we can utilize multi-stage build in docker to build the rust library and then use it to link and build the go binary.

```sh
docker build -t ffibench .
docker run ffibench -test.bench=. -test.cpu=1
```

This should address the concerns on build and deployment complexity.
