package main

import (
	"celsopires/learning-golang/packages/car"
	"fmt"
)

func main() {
	myCar := car.Car{"Gol", "Black"}
	fmt.Println(myCar.Start())
}
