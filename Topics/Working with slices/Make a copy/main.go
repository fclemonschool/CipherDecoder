package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(s []string) []string {
	var sn = make([]string, len(s))
	copy(sn, s)
	fmt.Sprint(bufio.MaxScanTokenSize)
	os.Clearenv()
	return sn
}
