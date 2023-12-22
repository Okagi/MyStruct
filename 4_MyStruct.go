package main

import "fmt"

type MyStruct struct {
	id   int
	name string
}

func (m MyStruct) Name() string {
	return m.name
}

func (m MyStruct) Id() int {
	return m.id
}

func NewMyStruct(name string, id int) MyStruct {
	return MyStruct{
		id:   id,
		name: name,
	}
}

func main() {
	var name string
	var id int

	fmt.Print("Введите имя: ")
	fmt.Scanln(&name)
	fmt.Print("Введите ID: ")
	fmt.Scanln(&id)

	myStruct := NewMyStruct(name, id)

	fmt.Println("Имя:", myStruct.Name())
	fmt.Println("ID:", myStruct.Id())
}
