package main

import "fmt"

func SomaSlice() int {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	if len(a) == 0 {
		return 0
	}

	total := 0
	for _, v := range a {
		total += v
	}
	return total
}

func main() {

	fmt.Println(SomaSlice())

}
