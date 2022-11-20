package tasks

//Реализовать паттерн «адаптер» на любом примере.

import "fmt"

// Создаем интерфейс телефона с разъемом наушников, который имплементирует структуру Samsung.
// В структуру IphoneAdapter встраиваем структуру Iphone и имплементируем интерфейс телефона с разъемом
// для наушников. По итогу логика структуры Iphone не нарушена, в него все также вставляются наушники
// через порт зарядки, но теперь с адаптером в него можно вставить и те наушники, для которых порт
// зарядке не подходит

type phone interface {
	insertIntoHeadphonePort()
}

type client struct{}

func (c *client) insertHeadphoneIntoPhone(p phone) {
	fmt.Println("Client insert headphone into phone")
	p.insertIntoHeadphonePort()
}

type samsung struct{}

func (s *samsung) insertIntoHeadphonePort() {
	fmt.Println("Headphone is plugged into Samsung's headphone port")
}

type iPhone struct{}

func (i *iPhone) insertIntoRechargePort() {
	fmt.Println("Headphone is plugged into iPhone's recharge port")
}

type iPhoneAdapter struct {
	iPhone *iPhone
}

func (ia *iPhoneAdapter) insertIntoHeadphonePort() {
	fmt.Println("Adapter turn iPhone's recharge port into Samsung's headphone port")
	ia.iPhone.insertIntoRechargePort()
}

func Task21() {
	client := &client{}
	samsung := &samsung{}
	
	client.insertHeadphoneIntoPhone(samsung)

	iPhone := &iPhone{}
	iPhoneAdapter := &iPhoneAdapter{iPhone}

	client.insertHeadphoneIntoPhone(iPhoneAdapter)

}
