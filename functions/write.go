package functions

import "fmt"

func Write(arr [][]string) {
	startP := 0
	for index, value := range arr {
		if len(value) == 1 && startP == index {
			fmt.Println()
			startP += 1
		} else if len(value) == 1 {
			printer(arr, startP, index)
			startP = index + 1
		} else if index+1 == len(arr) {
			printer(arr, startP, index+1)
		}
	}
}

func printer(array [][]string, startP int, endP int) {
	for i := 1; i <= 8; i++ {
		for j := startP; j < endP; j++ {
			fmt.Print(array[j][i])
		}
		fmt.Println()
	}
}
