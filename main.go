package main

import "gotest/system"

func main() {
	system.Run()
}

//func main() {
//	var n int
//	fmt.Scanln(&n)
//	ints := make([][]int, n, n)
//	a := 0
//	b := 0
//	for i := range ints {
//		ints[i] = make([]int, n, n)
//		for j := range ints[i] {
//			fmt.Scan(&ints[i][j])
//			if i == j {
//				a += ints[i][j]
//			}
//			if len(ints[i])-1-i == j {
//				b += ints[i][j]
//			}
//		}
//	}
//	c := a - b
//	if c < 0 {
//		fmt.Println(-c)
//	} else {
//		fmt.Println(c)
//	}
//}
