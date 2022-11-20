package tasks

//Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

import "fmt"

// Реализуем множества через map
// Если данные из слайса будут совпадать, то в мапе останется только один ключ со значением true
// В пустой слайс добавляем из мапы ключи множества
func set(s []string) []string {
	var res []string
	setOfString := make(map[string]bool, len(s))

	for _, val := range s {
		setOfString[val] = true
	}

	for key := range setOfString {
		res = append(res, key)
	}

	return res
}

func Task12() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}

	setArr := set(arr)

	fmt.Printf("Arr: %v\nSet: %v\n", arr, setArr)
}
