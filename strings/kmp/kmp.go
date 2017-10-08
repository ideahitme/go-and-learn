package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Buffer(make([]byte, 4096*1024), 4096*1028)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < n; i++ {
		scanner.Scan()
		a := scanner.Text()
		scanner.Scan()
		b := scanner.Text()
		res := kmp(b, a)
		if len(res) == 0 {
			fmt.Println("Not Found")
			continue
		}
		fmt.Println(len(res))
		for i := 0; i < len(res); i++ {
			fmt.Printf("%d ", res[i]+1)
		}
		fmt.Printf("\n")
	}
}

func kmp(p, s string) []int {
	occ := make([]int, 0)
	m := p + "$" + s
	n := len(p)
	pfx := prefix(m)
	for i := n; i < len(m); i++ {
		if pfx[i] == n {
			occ = append(occ, (i - 2*n))
		}
	}
	return occ
}

func prefix(m string) []int {
	border := 0
	n := len(m)
	s := make([]int, n)

	for i := 1; i < n; i++ {
		for border > 0 && m[border] != m[i] {
			border = s[border-1]
		}
		if m[border] == m[i] {
			border++
		} else {
			border = 0
		}
		s[i] = border
	}

	return s
}
