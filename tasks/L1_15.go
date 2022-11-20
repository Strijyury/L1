package tasks

import "fmt"

/*
К каким негативным последствиям может привести данный фрагмент кода,
и как это исправить? Приведите корректный пример реализации.
*/

//var justString string
// Объявление глобальной переменной считается плохой практикой, поскольку потом можно будет
// легко запутаться в том, какая переменная за что отвечает, могут происходить конфликты и коллизии
// Масштабировать код становится сложнее. При объявлении глобаных переменных ревьюить код станет затруднительно,
// поведение программы может стать непредсказуемым и определить алгоритм выполнения процессов будет сложно

//
//func someFunc() {
//	v := createHugeString(1 << 10)
//	justString = v[:100]
//}
//
//func main() {
//	someFunc()
//}

func createHugeString() []byte {
	hugeString := `It is a long established fact that a reader will be distracted by the readable content of a page
	when looking at its layout. The point of using Lorem Ipsum is that it has a more-or-less normal distribution of letters,
	as opposed to using "Content here, content here", making it look like readable English.
	Many desktop publishing packages and web page editors now use Lorem Ipsum as their default model text,
	and a search for "lorem ipsum" will uncover many web sites still in their infancy.Various versions have evolved
	over the years, sometimes by accident, sometimes on purpose (injected humour and the like).`

	hugeStringByte := []byte(hugeString)
	sliceForLen := hugeStringByte[:100]        // Зададим срез для создания интервала(len) будущего слайса
	newSlice := make([]byte, len(sliceForLen)) // Создадим байтовый срез длинной равной sliceForLen
	copy(newSlice, hugeStringByte)             // Копируем новый срез из нашего текста
	return newSlice
}

/*
Итак, зачем все эти танцы с бубном, для чего все это делалось.
В Го есть сборщик мусора, который очищает память любых сущностей Го, на которых мы перестаем ссылаться.
В нашем случае, у нас есть огрооомная строка, на которую мы ссылаемся когда делаем из нее срез.
С этого момента в памяти хранятся и большая строка и наш новый срез, который на нее ссылается.
Чтобы избежать хранения в памяти ненужных сущностей и оставить только полезные, наш newSlice,
нужно сделать его копию и тогда ссылки на hugeString уже будут не нужны и сборщик мусора Го освободит память,
оставив только newSlice.
*/

func Task15() {
	fmt.Printf("Срез из огромной строки равен: %c\n", createHugeString())
}
