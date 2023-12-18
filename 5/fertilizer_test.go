package main

import (
	_ "embed"
	"testing"
)

//go:embed input_test.txt
var TestInput string

func Test_PartOneTestInput(t *testing.T) {
	got := PartOne(TestInput)
	want := 35

	if got != want {
		t.Errorf("Do better, %v != %v", got, want)
	}
}

func Test_PartOneActual(t *testing.T) {
	got := PartOne(Input)
	want := 621354867

	if got != want {
		t.Errorf("Do better, %v != %v", got, want)
	}
}

func Test_PartTwoTestInput(t *testing.T) {
	got := PartTwo(TestInput)
	want := 46

	if got != want {
		t.Errorf("Do better, %v != %v", got, want)
	}
}

func Test_PartTwoActual(t *testing.T) {
	got := PartTwo(Input)
	want := 15880236

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
