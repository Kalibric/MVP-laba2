package main

import (
	"fmt"
	"math"
)

type m struct {
	a float64
	b float64
	c float64
	d float64
}

func main() {
	var (
		n         uint16
		n_counter         = []uint16{10}
		a         float64 = 0
		b         float64 = math.Pi

		M       []m
		resultM []float64
		lambda  []float64
		mu      []float64
	)
	for _, n = range n_counter {
		M = genM(a, b, n)
		lambda, mu = forwardSweep(M)
		resultM = backwardSweep(M, lambda, mu)

		fmt.Println(resultM)
	}

	fmt.Println(S3(1, resultM, a, b, n))
}

func d2f(value float64) float64 {
	return -math.Sin(value)
}

func genM(a float64, b float64, n uint16) []m {
	h := (b - a) / float64(n)
	//startOutput(a, b, n, h)
	M := genStartValues(make([]m, n+1))
	M = genValues(M, a, h)
	return M
}

func startOutput(a float64, b float64, n uint16, h float64) {
	fmt.Println("---------------------------")
	fmt.Println("Количество отрезков:", n)
	fmt.Println("[ a / b ]:", "[", a, "/", b, "]")
	fmt.Println("Шаг:", h)
	fmt.Println("---------------------------")
}

func genStartValues(M []m) []m {
	if len(M) < 1 {
		return M
	}
	M[0].c, M[0].a, M[0].b, M[0].d = 0, 1, 0, 0
	M[len(M)-1].c, M[len(M)-1].a, M[len(M)-1].b, M[len(M)-1].d = 0, 1, 0, 0
	return M
}

func genValues(M []m, a float64, h float64) []m {
	n := len(M) - 1
	for i := 1; i < n; i++ {
		//fmt.Println("i =", i)
		M[i].c = h / 6
		M[i].a = (h + h) / 3
		M[i].b = h / 6
		M[i].d = (f(a+h*float64(i+1))-f(a+h*float64(i)))/h - (f(a+h*float64(i))-f(a+h*float64(i-1)))/h

		//fmt.Println(M[i].c, M[i].a, M[i].b, M[i].d)
	}
	return M
}

func f(value float64) float64 {
	return math.Sin(value)
}

func forwardSweep(M []m) ([]float64, []float64) {
	lambda := make([]float64, len(M))
	lambda[0] = 0
	mu := make([]float64, len(M))
	mu[0] = 0
	for i := 1; i < len(M); i++ {
		lambda[i] = -(M[i].b) / (M[i].a + M[i].c*lambda[i-1])
		mu[i] = (M[i].d - M[i].c*mu[i-1]) / (M[i].a + M[i].c*lambda[i-1])
	}
	return lambda, mu
}

func backwardSweep(M []m, lambda []float64, mu []float64) []float64 {
	n := len(M) - 1
	result := make([]float64, len(M))
	result[n] = mu[n]
	for i := n - 1; i >= 0; i-- {
		result[i] = lambda[i]*result[i+1] + mu[i]
	}

	return result
}

func S3(x float64, M []float64, a float64, b float64, n uint16) float64 {
	if x > b || x < a {
		panic("err")
	}
	var (
		i   uint16
		x_1 = a
		x_2 float64
		y_1 float64
		y_2 float64
		m_1 float64
		m_2 float64
		h   = (b - a) / float64(n)
	)
	for i = 0; i < n-1 && x >= a+float64(i+1)*h; i++ {
	}
	x_1, x_2 = a+float64(i)*h, a+float64(i+1)*h
	m_1, m_2 = M[i], M[i+1]
	y_1, y_2 = f(x_1), f(x_2)
	return (m_1 * math.Pow(x_2-x, 3) / (h * 6)) +
		(m_2*math.Pow(x-x_1, 3))/(h*6) +
		(y_1-(m_1*h*h)/6)*((x_2-x)/h) +
		(y_2-(m_2*h*h)/6)*((x-x_1)/h)
}
