package main

import "fmt"

func main() {
	var (
		temperatura float32 = 24
		umidade     float32 = 80.0 / 100.0
		pressao     float32 = 1.05
	)
	temperaturaConfigurada := fmt.Sprintf("Temperatura em Celsius: %v°C", temperatura)
	umidadeConfigurada := fmt.Sprintf("\nPercentual de umidade: %.2f%", umidade*100)
	pressaoConfigurada := fmt.Sprintf("\nPressão atmosférica: %vPa", pressao)
	fmt.Println(temperaturaConfigurada, umidadeConfigurada, pressaoConfigurada)

}
