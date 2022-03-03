package main

import "fmt"

func create3DArray() any {
	array := [4][4][4]float32{1: {0: {
		2: 88.6,
	}}} // modify only this line
	fmt.Println()
	return array
}
