package main 

import "fmt"

//C = K - 273

var ebulicaoC = 100.0


func main(){
	var ebulicaoK = float64(ebulicaoC + 273.0)

	fmt.Printf("O ponto de ebulição da água em Celsius é %g , e o ponto de ebulição da água em Kelvin é %g", ebulicaoC, ebulicaoK)


}