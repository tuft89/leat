package main

import "fmt"

func main() {
	//	strs := []string{"floze", "flower", "fluize"}
	strs := []string{"floze"}
	firststr := strs[0]
	potision := 0

	if len(strs) == 1 {
		fmt.Println("koko")
		fmt.Println(firststr)
	}

	for k := 1; k < len(firststr)+1; k++ {
		fmt.Println("firststr[:k]")
		fmt.Println(firststr[:k])
		for i := 1; i < len(strs); i++ {
			tmpstr := strs[i]
			fmt.Println("tmpstr")
			fmt.Println(tmpstr[:k])
			if len(tmpstr) < k {

			if firststr[:k] != tmpstr[:k] {
				fmt.Println("break")
				fmt.Println(tmpstr[:k])
				break
			} else {
				if i == len(strs)-1 {
					potision = k
				}
			}
		}
	}
	if potision == 0 {
		fmt.Println("nil")
	} else {
		fmt.Println("result")
		fmt.Println(firststr[:potision])
	}
}
