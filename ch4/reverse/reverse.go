package main

import "fmt"

func reverse(array *[5]int) {
	for i, j := 0, len(array)-1; i < j; i, j = i+1, j-1 {
		array[i], array[j] = array[j], array[i]
	}
}

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	reverse(&a)
	fmt.Println(a)
}
