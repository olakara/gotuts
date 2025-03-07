package main

import "sync"

type MyStruct struct {
	Name string
}

func (m *MyStruct) GetName() string {
	return m.Name
}

func (m *MyStruct) SetName(name string) {
	m.Name = name
}

var (
	once     sync.Once
	instance *MyStruct
)

func GetMyStruct() *MyStruct {
	once.Do(func() {
		instance = &MyStruct{}
	})
	return instance
}

func main() {
	myStruct := GetMyStruct()
	myStruct.SetName("Hello")
	println(myStruct.GetName())
	myStruct2 := GetMyStruct()
	println(myStruct2.GetName())
}
