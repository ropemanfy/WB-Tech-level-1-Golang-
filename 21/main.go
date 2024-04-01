package main

import "fmt"

type Dungeon interface { // предположим, что у нас есть интерфейс, который используется в разных местах нашей программы
	Spank()
}

type Billy struct {
}

func NewBilly() Dungeon {
	return &Billy{}
}

func (b *Billy) Spank() {
	fmt.Println("Billy spanks")
}

type Van struct { // и в какой то момент мы решили добавить в него новый функционал из стороннего пакета
}

func (v *Van) SpankForMoney(bucks int) { // однако новые классы из пакета не реализуют наш интерфейс
	if bucks >= 300 {
		fmt.Println("Van spanks")
	}
}

type VanAdapter struct { // поэтому создаем оболочку, которая будет реализовывать наш интерфейс
	*Van
}

func (v *VanAdapter) Spank() {
	v.SpankForMoney(300) // но внутри сможет использовать свои методы
}

func NewVanAdapter() Dungeon {
	return &VanAdapter{}
}

func main() {
	billy := NewBilly()

	adapter := NewVanAdapter()

	adapter.Spank()
	billy.Spank()
}
