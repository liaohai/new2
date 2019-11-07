package main

import (
	"fmt"
	"go5/01区块链/test/get/calc"
)

func main() {
	add := calc.Add(1, 10)
	fmt.Println("Add(a, b int)",add)
	fmt.Println("++++")

	res :=calc.Sub(30 ,20)
	fmt.Println("++++",res)
	fmt.Println("====")
}
