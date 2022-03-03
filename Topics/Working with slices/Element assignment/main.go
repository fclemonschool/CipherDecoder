package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(s []string, i int) []string {
	// put your code here
	var sn = make([]string, len(s))
	copy(sn, s)
	sn[i] = "my string value"

	fmt.Sprint(bufio.MaxScanTokenSize)
	os.Clearenv()
	strconv.FormatBool(true)
	return sn
}
