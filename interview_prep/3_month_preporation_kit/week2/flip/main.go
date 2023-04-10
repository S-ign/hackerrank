package main

import (
	"fmt"
)

func convertToQuadrantPos(pos []int32, legnthOfMatrix int32) []int32 {
	for i, v := range pos {
		if v > (legnthOfMatrix/2)-1 {
			pos[i] = (legnthOfMatrix-1) - v
		}
	}
	return pos
}

func sumArr(arr []int32) (sum int32) {
	for _, v := range arr {
		sum += v
	}
	return
}

func sumMatrix(matrix [][]int32) (sum int32) {
	for _, arr := range matrix {
		sum += sumArr(arr)
	}
	return
}

func flippingMatrix(matrix [][]int32) int32 {
	hl := len(matrix)/2
	upperLeftQuardant := make([][]int32, hl)
	for i := 0; i < hl; i++ {
		upperLeftQuardant[i] = make([]int32, hl)
	}

	for r:=0; r<len(matrix); r++ {
		for c:=0; c<len(matrix); c++ {
			pos := convertToQuadrantPos([]int32{int32(r),int32(c)}, int32(len(matrix)))
			fmt.Printf("ULQ[%d][%d]: %d\nmatrix[%d][%d]: %d\n\n", pos[0], pos[1],
				upperLeftQuardant[pos[0]][pos[1]], r, c, matrix[r][c])
			if upperLeftQuardant[pos[0]][pos[1]] < matrix[r][c] {
				upperLeftQuardant[pos[0]][pos[1]] = matrix[r][c]
			}
		}
	}
	return sumMatrix(upperLeftQuardant)
}

func main() {
	m:= [][]int32{
		{112, 42, 83, 119},
		{56, 125, 56, 49},
		{15, 78, 101, 43,},
		{62, 98, 114, 108},
	}

	//matrix := [][]int32{
	//	{107, 54, 128, 15},
	//	{12, 75, 110, 138},
	//	{100, 96, 34, 85},
	//	{75, 15, 28, 112}, //488
	//}

	fmt.Println(flippingMatrix(m))
	//fmt.Println()
	//fmt.Println(flippingMatrix(matrix))

}
