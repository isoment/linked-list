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
	v, _ := list.GetByIndex(3)
	fmt.Println(v)
}
