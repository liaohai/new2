package main

import (
	"fmt"
	"go5/01区块链/test/get/calc"
)

func main() {
	add := calc.Add(1, 10)
	fmt.Println("Add(1+ 10 int)",add)
	fmt.Println("++++")

	res :=calc.Sub(30 ,20)
	fmt.Println("++++",res)



	res =calc.Multi(30 ,20)
	fmt.Println("====",res)


	res =calc.Multi(30 ,20)
	fmt.Println("====",res)


	fmt.Println("Sub(30 - 20)",res)

	res =calc.Multi(30 ,20)
	fmt.Println("Multi(30 * 20)",res)

	res =calc.Div(30 ,20)
	fmt.Println("Div(30/20)",res)


}
