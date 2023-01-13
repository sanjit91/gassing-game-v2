// Write a function that takes a slice of integers and returns a new slice with only the even numbers.
//    Input: [1, 2, 3, 4, 5]
//    Output: [2, 4]

package main

import "fmt"

func main() {

	fmt.Println(evenNumbers([]int{1, 2, 3, 4, 5})) // output: [2, 4]

}

func evenNumbers(nums []int) []int {
	evens := []int{}
	for _, num := range nums {
		if num%2 == 0 {
			evens = append(evens, num)
		}
	}
	return evens
}
