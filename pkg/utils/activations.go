package utils

import "math"

// note: the "apply" functions are usable in gonum matrices

// tanh
func Tanh(x float64) float64 {
	return math.Tanh(x)
}

func TanhPrime(x float64) float64 {
	return 1 - math.Pow(Tanh(x), 2)
}

func ApplyTanh(_, _ int, x float64) float64 {
	return Tanh(x)
}

func ApplyTanhPrime(_, _ int, x float64) float64 {
	return TanhPrime(x)
}

// silu
func Silu(x float64) float64 {
	return x / (1 + math.Exp(-x))
}

func SiluPrime(x float64) float64 {
	exp_negx := math.Exp(-x)
	return (1 + x*(1+exp_negx)/(1+exp_negx)) / (1 + exp_negx)
}

func ApplySilu(_, _ int, x float64) float64 {
	return Silu(x)
}

func ApplySiluPrime(_, _ int, x float64) float64 {
	return SiluPrime(x)
}

// relu
func Relu(x float64) float64 {
	return math.Max(0, x)
}

func ReluPrime(x float64) float64 {
	if x > 0 {
		return 1
	} else {
		return 0
	}
}

func ApplyRelu(_, _ int, x float64) float64 {
	return Relu(x)
}

func ApplyReluPrime(_, _ int, x float64) float64 {
	return ReluPrime(x)
}

// sigmoid
func Sigmoid(x float64) float64 {
	return 1 / (1 + math.Exp(-x))
}

func SigmoidPrime(x float64) float64 {
	return Sigmoid(x) * (1 - Sigmoid(x))
}

func ApplySigmoid(_, _ int, x float64) float64 {
	return Sigmoid(x)
}

func ApplySigmoidPrime(_, _ int, x float64) float64 {
	return SigmoidPrime(x)
}
