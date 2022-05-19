package myMath

import (
	"fmt"
	"testing"
)

type testCase struct {
	var1     int
	var2     int
	expected int
}

func TestAdd(t *testing.T) {
	testCases := []testCase{
		{5, 10, 15},
		{45, 70, 115},
		{39, 12, 51},
		{40, 20, 60},
		{-2, -40, -42},
		{-7, 30, 23},
	}

	for _, tc := range testCases {
		if Add(tc.var1, tc.var2) != tc.expected {
			t.Errorf(
				"Add operation failed. Var1=%d, Var2=%d, Expected=%d",
				tc.var1,
				tc.var2,
				tc.expected,
			)
		}
	}
}

func TestMultiply(t *testing.T) {
	testCases := []testCase{
		{5, 10, 50},
		{3, 70, 210},
		{30, 12, 360},
		{40, 20, 800},
		{-2, -40, 80},
		{-7, 30, -210},
	}

	for _, tc := range testCases {
		if Multiply(tc.var1, tc.var2) != tc.expected {
			t.Errorf(
				"Multiply operation failed. Var1=%d, Var2=%d, Expected=%d",
				tc.var1,
				tc.var2,
				tc.expected,
			)
		}
	}
}

func BenchmarkMultiply(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Multiply(123214314, 1232194343453)
	}
}

func BenchmarkSimpleInterst(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SimpleInterest(500000, 70, 15336254)
	}
}

type SubTestCase struct {
	name     string
	params   []int
	expected int
}

func TestSimpleInterest(t *testing.T) {
	cases := []SubTestCase{
		{"small_amount", []int{100, 5, 2}, 10},
		{"medium_amount", []int{3500, 10, 3}, 1050},
		{"large_amount", []int{100000, 10, 1}, 10000},
		{"negative_amount", []int{-1000, 5, 2}, 0},
		{"zeros", []int{0, 0, 0}, 0},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := SimpleInterest(tc.params[0], tc.params[1], tc.params[2])
			if tc.expected != got {
				t.Errorf("Expected:%d, Got:%d", tc.expected, got)
			}
		})
	}
}

//ExampleSampleInterest shows how SimpleInterest can be calculated
func ExampleSimpleInterest() {
	p := 1000
	r := 10
	t := 5
	fmt.Println(SimpleInterest(p, r, t))
	// Output: 500
}
