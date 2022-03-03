package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(s []int) {
	// put your code here
	for _, element := range s {
		element *= 2
		fmt.Println(element)
	}
	os.Clearenv()
	strconv.FormatBool(true)
	fmt.Sprint(bufio.MaxScanTokenSize)
}
