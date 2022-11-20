package tasks

//Разработать программу, которая в рантайме способна определить тип
//переменной: int, string, bool, channel из переменной типа interface{}.

import "fmt"

// С помощью switch в консоль выводится case соответствующий v.type
func typeOf(v interface{}) {
	switch v.(type) {
	case string:
		fmt.Printf("%s is String type(%T)\n", v, v)
	case int:
		fmt.Printf("%d is Integer type(%T)\n", v, v)
	case bool:
		fmt.Printf("%v is Boolean type(%T)\n", v, v)
	case chan int:
		fmt.Printf("%v is Channel int type(%T)\n", v, v)
	}
}

func Task14() {
	str := "Hello"
	num := 42
	boolean := true
	ch := make(chan int)

	typeOf(str)
	typeOf(num)
	typeOf(boolean)
	typeOf(ch)
}
