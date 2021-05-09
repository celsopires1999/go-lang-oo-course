package main

import "fmt"

func main() {

	fmt.Println("append")
	slice := make([]int, 5)
	slice = append(slice, 1, 10, 40, 40)
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Printf("Fim- append\n\n")

	fmt.Println("cap")
	slice1 := make([]int, 2, 5)
	slice1 = append(slice1, 10, 2, 5, 40)
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))
	fmt.Printf("Fim - cap\n\n")

	fmt.Println("dobrando cap")
	slice2 := make([]int, 2, 5)
	fmt.Println(slice2)
	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))

	for i := 0; i < 20; i++ {
		slice2 = append(slice2, i)
		fmt.Println("Len: ", len(slice2), "Cap: ", cap(slice2))
	}
	fmt.Printf("Fim - dobrando cap\n\n")

	fmt.Println("slices se separam")
	slice3 := make([]int, 2, 5)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	sliceTest := slice3
	slice3[0] = 99
	fmt.Println("slice3: ", slice3)
	fmt.Println("sliceTest: ", sliceTest)

	slice3 = append(slice3, 1, 2, 3, 5)
	slice3[0] = 10
	fmt.Println("slice3: ", slice3)
	fmt.Println("sliceTest: ", sliceTest)

	fmt.Printf("Fim - slices se separam\n\n")

	fmt.Println("slice string")

	sliceString := []string{
		"Hello",
		"World",
		"Much",
		"Better",
		"Better2",
	}

	fmt.Println(sliceString)
	fmt.Println(sliceString[0])   // Hello
	fmt.Println(sliceString[:2])  // Hello World
	fmt.Println(sliceString[1:2]) // World
	fmt.Println(sliceString[2:4]) // Much Better
	fmt.Println(sliceString[2:])  // Much Better

	fmt.Printf("Fim - slice string\n\n")

}
