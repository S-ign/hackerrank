package main

//d=4
//m=4
//[1,1,3,4,2,0]

func birthday(s []int32, d int32, m int32) int32 {
	bars := [][]int32{}
	var t int32
	for _, a := range s {
		if t <= d {
			t += a
		}
	}
}

func main() {
}
