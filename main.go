package main

import "fmt"

func main() {
	s := "i sure wish this text was sorted by character code!" // input string
	r := []rune(s)                                             // slices of runes can be manipulated how i need them to be (i don't know why yet)

	fmt.Println(string(r))

	// selection sort implementation
	for i := 0; i < len(r); i++ {

		smallest := i // the smallest value should start as the first unsorted value

		// go through each value in the unsorted part of the list and find the lowest value
		for j := i + 1; j < len(r); j++ {
			if r[smallest] > r[j] {
				smallest = j
			}
		}

		// swap the lowest value to the index that we were sorting for
		r[smallest], r[i] = r[i], r[smallest]
	}

	fmt.Println(r)
	fmt.Println(string(r))
}
