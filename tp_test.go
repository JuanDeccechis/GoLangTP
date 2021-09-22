package main

import (
	"testing"
)

func TestContenido(t *testing.T) {
	var contenido [4]Cadena
	contenido[0] = Cadena{"TX02AB"}
	contenido[1] = Cadena{"NN100987654321"}
	contenido[2] = Cadena{"TX06ABCDE"}
	contenido[3] = Cadena{"NN04000A"}
	for i := 0; i < len(contenido); i++ {
		ValidarEstructura(contenido[i])
	}
}
