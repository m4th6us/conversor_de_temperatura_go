package main

import "fmt"

func main(){

	var kelvin float32

	fmt.Println("Digite a temperatura em kelvin, e será convertido para celsius: ")
	fmt.Scan(&kelvin)

	celsius := conversor(kelvin)

	fmt.Println("A temperatura convertida para celsius é: ", celsius)

}

func conversor(kelvin float32) float32 {

	return kelvin - 273

}