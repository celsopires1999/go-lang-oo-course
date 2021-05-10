package main

import "fmt"

type Car struct {
	Name  string
	Year  int
	Color string
}

type SuperCar struct {
	Car
	CanFly bool
}

func (c Car) info() string {
	return fmt.Sprintf("Name: %s Year: %d Color: %s", c.Name, c.Year, c.Color)
}

func (c SuperCar) info() string {
	return fmt.Sprintf("Name: %s Year: %d Color: %s CanFly: %t", c.Name, c.Year, c.Color, c.CanFly)
}

func main() {

	car1 := Car{"Corolla", 2017, "White"}
	car2 := Car{"BMW 320i", 2018, "Black"}
	sCar := SuperCar{
		Car: Car{
			Name:  "Fusca",
			Year:  2030,
			Color: "Blue",
		},
		CanFly: true}

	fmt.Println("P1")
	fmt.Println(car1)
	fmt.Println(car2)

	fmt.Println("P2")
	fmt.Println(car1.Name)
	fmt.Println(car2.Name)

	fmt.Println("P3")
	fmt.Println(car1.info())
	fmt.Println(car2.info())

	fmt.Println("P4")
	fmt.Println(sCar)
	fmt.Println(sCar.info())

}
