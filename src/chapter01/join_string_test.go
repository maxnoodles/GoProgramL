package main

import (
	"strings"
	"testing"
)

func BenchmarkStringJoin1(b *testing.B) {
	// 1	1581441327 ns/op
	var a [100000]string
	for i := 0; i < 100000; i++ {
		a[i] = "1"
	}

	for i := 0; i < b.N; i++ {
		var ret string
		for _, i := range a {
			ret += i
		}
	}
}

func BenchmarkStringJoin2(b *testing.B) {
	// 2000	    844351 ns/op
	var a [100000]string
	for i := 0; i < 100000; i++ {
		a[i] = "1"
	}

	for i := 0; i < b.N; i++ {
		_ = strings.Join(a[:], "")
	}
}
