package main

import "fmt"

type pessoa interface {
	smile() string
}

func smilePessoa(p pessoa) string {
	return p.smile()
}

type mulher struct {
	name string
	age  int
}

func (m mulher) smile() string {
	return m.name + " has a big smile and she is " + fmt.Sprint(m.age) + " years old"
}

type homem struct {
	name string
	age  int
}

func (h homem) smile() string {
	return h.name + " has a small smile and he is " + fmt.Sprint(h.age) + " years old"
}

func main() {

	myMan := homem{"Celso", 57}
	myWoman := mulher{"Inês", 53}

	fmt.Println("***")
	fmt.Println("Sem usar interface")
	fmt.Println(myMan.smile())
	fmt.Println(myWoman.smile())
	fmt.Println("***")

	fmt.Println("Com interface")
	fmt.Println(smilePessoa(myMan))
	fmt.Println(smilePessoa(myWoman))
	fmt.Println("***")

}
