package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(s []string, i int) {
	// Put your code here
	s[i] = "my string value"
	os.Clearenv()
	strconv.FormatBool(true)
	size := bufio.MaxScanTokenSize
	fmt.Sprint(size)
}
