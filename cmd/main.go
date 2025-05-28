package main

import (
	"fmt"

	linkedlist "github.com/isoment/linked-list"
)

func main() {
	list := linkedlist.New[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)
	list.Append(5)

	_, err := list.Insert(6, 99)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(list.Values())
	}
}
