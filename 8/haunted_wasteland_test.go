package main

import (
	_ "embed"
	"testing"
)

//go:embed input_test_one.txt
var TestInputOne string

//go:embed input_test_two.txt
var TestInputTwo string

func Test_PartOneTestInput(t *testing.T) {
	got := PartOne(TestInputOne)
	want := 2

	if got != want {
		t.Errorf("Do better, %v != %v", got, want)
	}
}

func Test_PartOneActual(t *testing.T) {
	got := PartOne(Input)
	want := 19951

	if got != want {
		t.Errorf("Do better, %v != %v", got, want)
	}
}

func Test_PartTwoTestInput(t *testing.T) {
	got := PartTwo(TestInputTwo)
	want := 6

	if got != want {
		t.Errorf("Do better, %v != %v", got, want)
	}
}

func Test_PartTwoActual(t *testing.T) {
	got := PartTwo(Input)
	want := 16342438708751

	if got != want {
		t.Errorf("Do better, %v != %v", got, want)
	}
}

func Benchmark_PartOne(b *testing.B) {
	for n := 0; n < b.N; n++ {
		PartOne(Input)
	}
}

func Benchmark_PartTwo(b *testing.B) {
	for n := 0; n < b.N; n++ {
		PartTwo(Input)
	}
}
