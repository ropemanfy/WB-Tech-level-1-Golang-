package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewInt(0) // создаем переменные типа big.Int
	b := big.NewInt(0)

	a.SetString("3298523985235983025724263434162", 10) // устанавливаем их значение
	b.SetString("720375209357230570235239023", 10)

	sum := new(big.Int)
	sum.Add(a, b) // находим сумму

	sub := new(big.Int)
	sub.Sub(a, b) // разность

	mul := new(big.Int)
	mul.Mul(a, b) // произведение

	div := new(big.Int)
	div.Div(a, b) // частное

	fmt.Println(sum)
	fmt.Println(sub)
	fmt.Println(mul)
	fmt.Println(div)
}
