package main

import (
	"fmt"
	"reflect"
	"sort"
	"strconv"
)

type matrixType [][]int32

//func (m matrixType) convertMatrixInt32ToInt() (intMatrix [][]int) {
//	matrixCopy := make([][]int32, len(m))
//	copy(matrixCopy, m)
//	intMatrix = make([][]int, len(matrixCopy))
//	//fmt.Println(matrixCopy)
//	for r := 0; r < len(matrixCopy); r++ {
//		intMatrix[r] = make([]int, len(matrixCopy))
//		for c := 0; c < len(matrixCopy); c++ {
//			intMatrix[r][c] = int(matrixCopy[r][c])
//		}
//	}
//	//fmt.Println(intMatrix)
//	return
//}

func (m matrixType) flattenMatrix() (flatten []int32) {
	for _, r := range m {
		for _, c := range r {
			flatten = append(flatten, c)
		}
	}
	return
}

// flattenAndSortDecending flattens a 2d slice into a 1d slice and sorts
// them in decending order [[1,2],[3,4]] -> [4,3,2,1]
func (m *matrixType)flattenAndSortDecending() (arr []int32) {
	arr = m.flattenMatrix()
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})	
	return arr
}

//func (m matrixType) mapPosition() (mappedPositions map[int][][]int) {
//	mappedPositions = make(map[int][][]int, len(m))
//	for _, v := range m.flattenAndSortDecending() {
//		//fmt.Println(v)
//		for ri, r := range m {
//			for ci, c := range r {
//				if v == int(c) {
//					//fmt.Println(c)
//					if !isInMatrix([]int{ri,ci}, mappedPositions[v]){
//						mappedPositions[v] = append(mappedPositions[v], []int{ri, ci})
//					}
//				}
//			}
//		}
//	}
//	return
//}

func (m matrixType) mapQuadrentPosition(decending []int32) (mappedPositions map[int32][][]int32) {
	mappedPositions = make(map[int32][][]int32, len(m))
	for _, v := range decending {
		//fmt.Println(v)
		for ri, r := range m {
			for ci, c := range r {
				if v == c {
					ctqp := convertToQuadrantPos([]int32{int32(ri),int32(ci)}, int32(len(m)))
					//fmt.Println(c)
					if !isInMatrix(ctqp, mappedPositions[v]){
						mappedPositions[v] = append(mappedPositions[v], ctqp)
					}
				}
			}
		}
	}
	return
}

func isInMatrix(arr []int32, matrix [][]int32) bool {
	for _, v := range matrix {
		if reflect.DeepEqual(arr, v) {
			return true
		}
	}
	return false
}

func isInStringArr(s string, sa []string) bool {
	for _, v := range sa {
		if s == v {
			return true
		}
	}
	return false
}

func convertToQuadrantPos(pos []int32, legnthOfMatrix int32) []int32 {
	for i, v := range pos {
		if v > (legnthOfMatrix/2)-1 {
			pos[i] = (legnthOfMatrix-1) - v
		}
	}
	return pos
}

func flippingMatrix(matrix [][]int32) (largest int32) {
	var m matrixType
	m = matrix
	largestNumbers :=  m.flattenAndSortDecending()
	quadrantPositions := m.mapQuadrentPosition(largestNumbers)
	checkPositions := []string{}

	for _, v := range largestNumbers {
		for _, r := range quadrantPositions[v] {
			s := strconv.Itoa(int(r[0]))+strconv.Itoa(int(r[1]))
			if !isInStringArr(s, checkPositions) {
				checkPositions = append(checkPositions, s)
				//fmt.Println(v)
				largest += v
			}
		}
	}

	return
}

func main() {
	//m := new(matrixType)
	//m = &matrixType{
	//	{107, 54, 128, 15},
	//	{12, 75, 110, 138},
	//	{100, 96, 34, 85},
	//	{75, 15, 28, 112}, //488
	//}
	// 112, 128, 138, 110

	m:= [][]int32{
		{112, 42, 83, 119},
		{56, 125, 56, 49},
		{15, 78, 101, 43,},
		{62, 98, 114, 108},
	}

	matrix := [][]int32{
		{107, 54, 128, 15},
		{12, 75, 110, 138},
		{100, 96, 34, 85},
		{75, 15, 28, 112}, //488
	}

	fmt.Println(flippingMatrix(m))
	fmt.Println()
	fmt.Println(flippingMatrix(matrix))

}
