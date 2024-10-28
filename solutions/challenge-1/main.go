package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	jsonFile, err := os.ReadFile("../../files/hard.json")
	if err != nil {
		fmt.Println(err)
	}
	var data [][]int
	err = json.Unmarshal(jsonFile, &data)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(MaxSum(data))
}

func MaxSum(arr [][]int) int {
	m := make(map[[2]int]int)
	return traverse(m, arr, 0, 0)
}

func traverse(m map[[2]int]int, arr [][]int, r, c int) int {
	if r >= len(arr) || c >= len(arr) {
		return 0
	}
	if v, ok := m[[2]int{r, c}]; ok {
		return v
	}
	left := traverse(m, arr, r+1, c)
	right := traverse(m, arr, r+1, c+1)
	output := arr[r][c]
	if left > right {
		output += left
	} else {
		output += right
	}
	m[[2]int{r, c}] = output
	return output
}
