package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(s1, s2 []string) []string {
	// put your code here
	s1 = append(s1, s2...)

	os.Clearenv()
	fmt.Sprint(bufio.MaxScanTokenSize)
	return s1
}
