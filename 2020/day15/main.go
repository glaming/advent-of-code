package main

import "fmt"

func main() {
	numbers := []int{1,0,15,2,10,13}

	lastSpoken := make(map[int]int, 0)
	for i, n := range numbers {
		lastSpoken[n] = i+1
	}

	for i := len(numbers); i < 30000000; i++ {
		prevNum := numbers[i-1]
		n, ok := lastSpoken[prevNum]

		spokenNum := 0
		if ok {
			spokenNum = i - n
		}
		lastSpoken[prevNum] = i
		numbers = append(numbers, spokenNum)
	}

	fmt.Println(numbers[29999999])
}
