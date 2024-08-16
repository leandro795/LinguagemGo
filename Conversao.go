package main

import (
	"fmt"
)

func kelvinToCelsius(kelvin float64) float64 {
	return kelvin - 273.15
}

func main() {
	var kelvin float64

	fmt.Print("Digite a temperatura em Kelvin: ")
	_, err := fmt.Scanf("%f", &kelvin)
	if err != nil {
		fmt.Println("Erro ao ler a entrada:", err)
		return
	}

	celsius := kelvinToCelsius(kelvin)
	fmt.Printf("A temperatura em graus Celsius é: %.2f°C\n", celsius)
}
