package ffibench

/*
#cgo LDFLAGS: -L${SRCDIR}/rust/target/release -lrustffi -ldl
#include "rust/target/release/librustffi.h"
*/
import "C"

func Add() int {
	return int(C.add(1, 2))
}

func Fibonacci(n int) int {
	return int(C.fibonacci(C.uint(n)))
}
