package main

import "fmt"

func main() {
	// 初始化工厂
	adidasFactory, _ := getSportFactory("adidas")
	nikeFactory, _ := getSportFactory("nike")

	nikeShoe := nikeFactory.makeShoe()
	nikeShirt := nikeFactory.makeShirt()

	adidasShoe := adidasFactory.makeShoe()
	adidasShirt := adidasFactory.makeShirt()

	fmt.Println(nikeShirt, nikeShoe)
	fmt.Println(adidasShirt, adidasShoe)
}
