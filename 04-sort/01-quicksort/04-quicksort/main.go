package main

import "fmt"

func quicksort(list []int) []int {
	if len(list) < 2 {
		return list
	} else {
		pivot := list[0]

		less := []int{}
		greater := []int{}

		for _, num := range list[1:] {
			if pivot > num {
				less = append(less, num)
			} else {
				greater = append(greater, num)
			}
		}
		less = append(quicksort(less), pivot)
		greater = quicksort(greater)
		return append(less, greater...)
	}
}

func main() {
	s := []int{10, 5, 2, 3}
	fmt.Println(s)
	fmt.Println(quicksort(s))
}
