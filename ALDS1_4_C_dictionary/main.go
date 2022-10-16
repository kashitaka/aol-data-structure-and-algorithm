package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var sc = bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	set := make(map[string]struct{}, n)
	for i := 0; i < n; i++ {
		sc.Scan()
		command := sc.Text()
		sc.Scan()
		str := sc.Text()
		switch command {
		case "insert":
			set[str] = struct{}{}
		case "find":
			_, ok := set[str]
			if ok {
				fmt.Println("yes")
			} else {
				fmt.Println("no")
			}
		}
	}
}
