package main

import "fmt"

type SavingBox struct {
	money int
}

func NewBox() *SavingBox {
	return new(SavingBox)
}

func (s *SavingBox) Income(amount int) {
	s.money += amount
}

func (s *SavingBox) Break() int {
	lastMoney := s.money
	s.money = 0
	return lastMoney
}

func (s SavingBox) Peep() int {
	return s.money
}

func mainReceiverSample() {
	//box := NewBox()
	box := SavingBox{}
	fmt.Printf("%d\n", box.Peep())
	box.Income(100)
	fmt.Printf("%d\n", box.Peep())
	box.Income(200)
	fmt.Printf("%d\n", box.Peep())
	box.Income(500)
	fmt.Printf("%d\n", box.Peep())

	fmt.Printf("貯金箱を壊したら%d円出てきました。", box.Break())
}
