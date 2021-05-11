package main

import "fmt"

type vehicle interface {
	start() string
}

func startVehicle(v vehicle) string {
	return v.start()
}

type car struct {
	name string
}

func (c car) start() string {
	return "The car " + c.name + " has been started"
}

type motorcycle struct {
	name string
}

func (m motorcycle) start() string {
	return "The motorcycle " + m.name + " has been started"
}

func main() {

	c := car{"Gol"}
	m := motorcycle{"XPTO"}

	//sem usar interface
	fmt.Println("Sem usar interface")
	fmt.Println(c.start())
	fmt.Println(m.start())

	//usando interface
	fmt.Println("Usando interface")
	fmt.Println(startVehicle(c))
	fmt.Println(startVehicle(m))

}
