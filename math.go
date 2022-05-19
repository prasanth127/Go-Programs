package myMath

func Add(a, b int) int {
	return a + b
}

func Multiply(a, b int) int {
	return a * b
}

func SimpleInterest(p, r, t int) int {
	if p < 0 {
		return 0
	}
	return (p * r * t) / 100
}
