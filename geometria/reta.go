package geometria

import "math"

// Ponto representa uma coordenada no plano cartesiano
type Ponto struct {
	X float64
	Y float64
}

//Catetos pega os modulos dos catetos do triangulo formado por 2 pontos
func Catetos(pontoA, pontoB Ponto) (cx, cy float64) {
	cx = math.Abs(pontoB.X - pontoA.X)
	cy = math.Abs(pontoB.Y - pontoA.Y)

	return 
}

// Distancia linear entre 2 pontos
func Distancia(pontoA,pontoB Ponto) float64 {
	catetoX, catetoY := Catetos(pontoA, pontoB)
	return math.Sqrt(math.Pow(catetoX, 2) + math.Pow(catetoY, 2))
}

