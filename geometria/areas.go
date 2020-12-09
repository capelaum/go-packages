package geometria

import "math"

// Circ calcula a área de uma circunferência
func areaCirc(raio float64) float64 {
	return math.Pow(raio, 2) * math.Pi
}

// Rect calcula a área de um retangulo
func areaRetangulo(base, altura float64) float64 {
	return base * altura
}

// TrianguloEq calcula a área de um triangulo equilatero
func areaTrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}
