package main

import (
	"fmt" )

func main() {
	var n int
	i := 0
	s := 0
	fmt.Scan(&n)
  
	for ; i <= n; i++ {
		s += i
	}
	fmt.Println(s)
}
