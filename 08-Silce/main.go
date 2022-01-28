package main

import "fmt"

func main() {
	s := make([]string, 5)
	s[0] = "a"
	s = append(s, "d")
	fmt.Println(s)

	s2d := make([][]int, 3)
	fmt.Println(s2d)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		s2d[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			s2d[i][j] = i + j
		}
	}
	fmt.Println(s2d)
}
