package main

import (
	"fmt"
	"go-circle/utils"
)

func main(){
	radius :=  10.0
	fmt.Println("Luas Lingkaran : ", utils.AreaCircle(radius))
	fmt.Println("Keliling Lingkaran : ", utils.CircumferenceCircle(radius))
}