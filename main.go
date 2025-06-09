package main

import "fmt"

func main(){
	radius :=  10.0
	fmt.Printf("Luas Lingkaran : %v\n", areaCircle(radius))
	fmt.Printf("Keliling Lingkaran : %v\n", circumferenceCircle(radius))
}

func areaCircle(radius float64) float64{
	area := 3.14 * radius * radius
	return area
}
func circumferenceCircle(radius float64) float64{
	circumference := 2 * 3.14 * radius
	return circumference
}