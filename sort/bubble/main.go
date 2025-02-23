package main

import (
	"log"
)

var data = []int{5, 4, 1, 7, 6, 9}

func bubble() []int {

	datalen := len(data)
	var swapped = make([]int, datalen)
	copy(swapped, data)

	for index := 0; index < datalen; index++ {

		if index+1 < datalen-1 {

			if swapped[index] > swapped[index+1] {
				tmp := swapped[index]
				swapped[index] = swapped[index+1]
				swapped[index+1] = tmp
			}
		}

		if index+1 == datalen-1 {
			if swapped[index] > swapped[index+1] {
				tmp := swapped[index]
				swapped[index] = swapped[index+1]
				swapped[index+1] = tmp
			}
		}
	}

	return swapped
}

func main() {
	log.Printf("Before sort: %v, after sort: %v", data, bubble())
}
