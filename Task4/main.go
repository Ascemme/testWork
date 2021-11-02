package main

import "fmt"

var name string

type inStorage struct {
	itemsIn string
}

type itemsinStore struct {
	circle    string
	rectangle string
	square    string
	triangle  string
}

type madeItem struct {
	shape     string
	matearial string
}

type inStorageIs []inStorage

func main() {
	start()
}
func start() {
	fmt.Println("Вы заказчик = 1  или налоговая = 2")
	var input string
	fmt.Scanln(&input)
	if input == "1" {
		orderForShape()

	}
	if input == "2" {
		orderForDitals()
	} else {
		fmt.Println("нет такого значения, ведите еще раз")
		start()
	}
}
func orderForShape() {
	fmt.Println("На какое имя сделать заказ")
	fmt.Println("круг =1, прямоугольник =2 квадрат =3 ,равносторонний  треугольник =4")
	var input string
	fmt.Scanln(&input)
	if input == "1" || input == "2" || input == "3" || input == "4" {
		choosingMaterial(input)
	} else {
		orderForShape()
	}
}
func choosingMaterial(shape string) {
	var material1 string
	fmt.Println("какой материал выберим")
	fmt.Println("жести =1, пластика =2, карбона=3")
	fmt.Scanln(&material1)
	if material1 == "1" || material1 == "2" || material1 == "3" {
		working(shape, material1)
	} else {
		choosingMaterial(shape)
	}

}
func working(shape string, material string) {
	if shape == "1" {
		shape = "круг"
	}
	if shape == "2" {
		shape = "прямоугольник"
	}
	if shape == "3" {
		shape = "квадрат"
	}
	if shape == "4" {
		shape = "равносторонний  треугольник"
	}

	if material == "1" {
		material = "жести"
	}
	if material == "2" {
		material = "пластика"
	}
	if material == "3" {
		material = "карбона"
	}
	itemsMade := madeItem{"мы сделали уже " + shape, "из " + material}
	fmt.Println(itemsMade)
}

func orderForDitals() {

}
