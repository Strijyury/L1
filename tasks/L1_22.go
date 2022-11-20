package tasks

//Разработать программу, которая перемножает, делит, складывает, вычитает
//две числовых переменных a,b, значение которых > 2^20.

import (
	"fmt"
	"math/big"
)

// Используем пакет math/big для операций с громадными числовыми переменными
// сложение
func add(a, b *big.Int) *big.Int {
	return new(big.Int).Add(a, b)
}

// вычитание
func sub(a, b *big.Int) *big.Int {
	return new(big.Int).Sub(a, b)
}

// умножение
func multiply(a, b *big.Int) *big.Int {
	return new(big.Int).Mul(a, b)
}

// деление
func divide(a, b *big.Int) *big.Int {
	return new(big.Int).Div(a, b)
}

func Task22() {

	bigIntA := big.NewInt(int64(1 << 30))
	bigIntB := big.NewInt(int64(1 << 25))

	fmt.Println(bigIntA)
	fmt.Println(bigIntB)

	fmt.Println("Adding is equal: ", add(bigIntA, bigIntB))
	fmt.Println("Subtracting is equal: ", sub(bigIntA, bigIntB))
	fmt.Println("Multiplying is equal: ", multiply(bigIntA, bigIntB))
	fmt.Println("Devision is equal: ", divide(bigIntA, bigIntB))

}
