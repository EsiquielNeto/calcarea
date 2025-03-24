package area

import "math"

const PI = 3.1416

// Circ é responsável por calcular a area do quadrado
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * PI
}

// Rect é responsável por calcular a are do retangulo
func Rect(base, altura float64) float64 {
	return base * altura
}

// Não é visível
func _TrainguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}
