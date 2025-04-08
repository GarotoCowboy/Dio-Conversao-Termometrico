package main

import "fmt"

func main() {
	const kelvinEbulicao = 373.15
	conversao := kelvinEbulicao - 273.15

	fmt.Printf("O valor de ebulição %vK para celcius é %vºC", kelvinEbulicao, conversao)

}
