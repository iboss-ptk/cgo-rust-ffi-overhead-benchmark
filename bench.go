package ffibench

/*
#cgo LDFLAGS: ${SRCDIR}/rust/target/release/librust.a -ldl
#include "rust/target/release/librust.h"
*/
import "C"

func Add() int {
	return int(C.add(1, 2))
}

func Fibonacci(n int) int {
	return int(C.fibonacci(C.uint(n)))
}
