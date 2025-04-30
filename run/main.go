package main

import (
	"fmt"

	"github.com/bytedance/sonic"
)

type Demo struct {
	ID      int
	Name    string
	Age     int
	Address string
	Phone   string
}

func main() {

	demo := Demo{
		ID:      1,
		Name:    "jettcc",
		Age:     20,
		Address: "123 Main St",
		Phone:   "555-1234",
	}

	s, err := sonic.Marshal(demo)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(s))

	demo2 := `{"id":1,"name":"jettcc","age":20,"address":"123 Main St","phone":"555-1234"}`
	err = sonic.Unmarshal([]byte(demo2), &demo2)
	if err != nil {
		panic(err)
	}

	fmt.Println(demo2)
}
