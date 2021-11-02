package main

import "fmt"

func main() {
	x := []int{17, 18, 5, 4, 6, 1}
	fmt.Println(x)
	fmt.Println(logic(x))
	y := []int{4, 0, 0}
	fmt.Println(logic(y))

	M := []int{15, 23, 46, 7, 20, 30}
	//mOut{46,46,30,30,30,-1}
	fmt.Println(logic(M))
}

func maxValue(x []int) (i int) {
	var maxV int
	for _, value := range x {

		if maxV < value {
			maxV = value
		}
	}
	return maxV
}

func logic(x []int) []int {
	var newArray []int
	var y []int
	for i := range x {
		if i != len(x)-1 {
			newArray = x[i+1 : len(x)]
			if maxValue(newArray) != 0 {
				y = append(y, maxValue(newArray))
			}
		} else {
			y = append(y, -1)
		}

	}
	return y
}

//Input: arr = [17,18,5,4,6,1]

//Output: [18,6,6,6,1,-1]

// index 0 --> the greatest element to the right of index 0 is index 1 (18).
// index 1 --> the greatest element to the right of index 1 is index 4 (6).
// index 2 --> the greatest element to the right of index 2 is index 4 (6).
// index 3 --> the greatest element to the right of index 3 is index 4 (6).
// index 4 --> the greatest element to the right of index 4 is index 5 (1).
// index 5 --> there are no elements to the right of index 5, so we put -1.
