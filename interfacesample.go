package main

import "fmt"

type Animal interface {
	Cry()
}

type Dog struct{}

func (d *Dog) Cry() {
	fmt.Println("わんわん")
}

type Cat struct{}

func (c *Cat) Cry() {
	fmt.Println("ニャーニャー")
}

func MakeAnimalCry(a Animal) {
	fmt.Println("鳴け！")
	a.Cry()
}

type Butterfly struct{}

func MakeSomeoneCry(someone interface{}) {
	fmt.Println("鳴け！")
	a, ok := someone.(Animal)
	if !ok {
		fmt.Println("動物では無いので鳴けません。")
		return
	}
	a.Cry()
}

func mainInterfaceSample() {
	dog := new(Dog)
	cat := new(Cat)
	MakeAnimalCry(dog)
	MakeAnimalCry(cat)
	butterfly := new(Butterfly)
	MakeSomeoneCry(dog)
	MakeSomeoneCry(cat)
	MakeSomeoneCry(butterfly)
}
