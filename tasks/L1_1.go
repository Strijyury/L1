package tasks

//Дана структура Human (с произвольным набором полей и методов).
//Реализовать встраивание методов в структуре Action от родительской
//структуры Human (аналог наследования).

import "fmt"

// определяем структуру Human
type Human struct {
	Name  string
	Speed int
}

// определяем метод структуры Human
func (h Human) PrintWalkTime(meters int) {
	time := meters / h.Speed
	fmt.Printf("%s прошел %d метров за %d секунд\n", h.Name, meters, time)
}

// определяем структуру Action с встраиванием типа Human
type Action struct {
	Human
}

// определение метода структуры Action со встроенными параметрами структуры Human
func (a Action) PrintRunTime(meters int) {
	time := meters / (a.Speed * 2)
	fmt.Printf("%s пробежал %d метров за %d секунд\n", a.Name, meters, time)
}

// реализация методов Action
func Task1() {
	Runner := Action{Human{Name: "Юра", Speed: 10}}
	fmt.Printf("Type: %T value: %#v\n", Runner, Runner)
	fmt.Println(Runner.Human.Speed)
	fmt.Println(Runner.Speed)

	var meters int = 20
	Runner.PrintRunTime(meters)
	Runner.PrintWalkTime(meters)
}
