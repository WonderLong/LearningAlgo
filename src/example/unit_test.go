package main

import (
	"reflect"
	"testing"
)

type A struct {
}

func F1() {
	_ = A{}
}

func F2(x interface{}) {
	_ = reflect.ValueOf(x)
}

func BenchmarkNormal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		F1()
	}
}

func BenchmarkReflection(b *testing.B) {
	var x A
	for i := 0; i < b.N; i++ {
		F2(x)
	}
}
