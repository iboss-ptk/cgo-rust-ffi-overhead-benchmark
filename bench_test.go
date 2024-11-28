package ffibench_test

import (
	"encoding/json"
	"io"
	"strings"
	"testing"

	"github.com/iboss-ptk/ffibench"
)

func BenchmarkAddFFI(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			ffibench.Add()
		}
	})
}

func BenchmarkFibonacciFFI(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			ffibench.Fibonacci(100000)
		}
	})
}

// use marhalling `1` as a baseline comparison for the overhead of calling CGO
// There are 2 versions of this benchmark:
// - `BenchmarkLowLevelJSONCParseIntBaseline` use lower level primitives to parse the JSON string
// - `BenchmarkJSONMarshalCParseIntBaseline` use `json.Marshal` to parse the JSON string

func BenchmarkLowLevelJSONCParseIntBaseline(b *testing.B) {
	msg := `1`

	b.RunParallel(func(pb *testing.PB) {
		var dst int
		r := strings.NewReader(msg)
		dec := json.NewDecoder(r)
		for pb.Next() {
			r.Seek(0, io.SeekStart)
			if err := dec.Decode(&dst); err != nil {
				panic(err)
			}
		}
	})
}

func BenchmarkJSONMarshalCParseIntBaseline(b *testing.B) {
	msg := `1`

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			json.Marshal(msg)
		}
	})
}
