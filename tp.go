package main

import (
	"fmt"
	"os"
	"strconv"
)

type Cadena struct {
	Input string
}

type Estructura struct {
	tipo      string
	tamaño    int
	contenido string
}

func (c *Cadena) GenerarEstructura(contenido string) Estructura {
	tipo := string(contenido[0]) + string(contenido[1])

	tamañoD, err := strconv.Atoi(string(contenido[2])) // + string(args.Input[3])
	tamañoU, err := strconv.Atoi(string(contenido[3])) // + string(args.Input[3])
	tamañoF := 0

	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	} else {
		tamañoD = tamañoD * 10
		tamañoF = tamañoD + tamañoU
	}

	value := ""
	for i := 4; i < len(contenido); i++ {
		value = value + string(contenido[i])
	}

	est := Estructura{tipo, tamañoF, value}
	return est
}

func (e *Estructura) ValidarTamaño() bool {
	return e.tamaño == len(e.contenido)
}

func (e *Estructura) ValidarNumero(tipo string) bool {
	for i := 0; i < len(tipo); i++ {
		value, err := strconv.Atoi(string(tipo[i]))
		if err != nil {
			fmt.Println("ERROR, el valor " + tipo + " no es un número ")
			fmt.Println(err)
		} else {
			if !((0 <= value) && (value <= 9)) {
				fmt.Println("ERROR, el valor " + tipo + " no es un número ")
				fmt.Println(err)
			}
		}
	}
	fmt.Println("tipo válido para el valor: " + tipo)
	return true
}

func ValidarEstructura(contenido Cadena) {
	estructura := contenido.GenerarEstructura(contenido.Input)
	validaTipo := estructura.ValidarNumero(strconv.Itoa(estructura.tamaño))
	if validaTipo {
		fmt.Println("tipo size OK")
	} else {
		fmt.Println("err: size no es del tipo correcto")
	}
	if estructura.tipo == "NN" {
		validaTipoContenido := estructura.ValidarNumero(string(estructura.contenido))
		if validaTipoContenido {
			fmt.Println("tipo size OK")
		} else {
			fmt.Println("err: size no es del tipo correcto")
		}
	}
	validaTamaño := estructura.ValidarTamaño()
	if validaTamaño {
		fmt.Println("tamaño OK")
	} else {
		fmt.Println("err: size y tamaño de contenido no coinciden")
	}

	fmt.Println(estructura)
}

func main() {
	contenido := Cadena{os.Args[1]}
	ValidarEstructura(contenido)
}
