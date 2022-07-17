package main

import (
	"fmt"

	"github.com/vitalvas/uxid"
)

func main() {
	id := uxid.New().String()

	fmt.Println("id:", id)
	fmt.Println("id len:", len(id))
}
