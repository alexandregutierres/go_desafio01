package main

import "fmt"

const ebulicaoKelvin = 373.0

func main() {

	ebulicaoAgua := ebulicaoKelvin - 273.0

	fmt.Println("Ponto de ebulição Kelvin: ", ebulicaoKelvin)
	fmt.Println("Ponto de ebulição Celcius: ", ebulicaoAgua)

}
