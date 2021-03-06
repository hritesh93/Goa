package main

import "fmt"

// O(n+m) time | O(1) space
func SearchInSortedMatrix(matrix [][]int, target int) []int {
	row, col := 0, len(matrix[0])-1
	for row < len(matrix) && col >= 0 {
		if matrix[row][col] > target {
			col -= 1
		} else if matrix[row][col] < target {
			row += 1
		} else {
			return []int{row, col}
		}
	}
	return []int{-1, -1}
}

func main() {
	matrix := [][]int{
		{1, 4, 7, 12, 15, 1000},
    	{2, 5, 19, 31, 32, 1001},
    	{3, 8, 24, 33, 35, 1002},
    	{40, 41, 42, 44, 45, 1003},
    	{99, 100, 103, 106, 128, 1004},
	}
	target := 44
	fmt.Println(SearchInSortedMatrix(matrix, target))
}
