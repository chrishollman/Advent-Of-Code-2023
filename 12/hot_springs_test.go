package main

import (
	_ "embed"
	"testing"
)

//go:embed input_test.txt
var TestInput string

func Test_PartOneTestInput(t *testing.T) {
	got := PartOne(TestInput)
	want := 21

	if got != want {
		t.Errorf("Do better, %v != %v", got, want)
	}
}

func Test_PartOneActual(t *testing.T) {
	got := PartOne(Input)
	want := 6488

	if got != want {
		t.Errorf("Do better, %v != %v", got, want)
	}
}

func Test_PartTwoTestInput(t *testing.T) {
	got := PartTwo(TestInput)
	want := 525152

	if got != want {
		t.Errorf("Do better, %v != %v", got, want)
	}
}

func Test_PartTwoActual(t *testing.T) {
	got := PartTwo(Input)
	want := 815364548481

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
