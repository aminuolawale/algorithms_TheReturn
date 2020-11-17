
package main

import ("fmt")


func main(){
	cases :=[][]int32{{0 ,3} ,{4, 6}, {6, 7} ,{3, 5}, {0, 7}}
	width := []int32{2, 3, 1, 2, 3, 2, 3, 3}
	answer := serviceLane(8, cases, width)
	fmt.Println(answer)
}


func serviceLane(n int32, cases [][]int32, width [] int32) []int32 {
	var answers []int32
	for _, c := range cases {
		answers = append(answers, min(width[c[0]:c[1]+1]))
	}
	return answers
}



func min(array [] int32)int32{
	min := array[0]
	for _, val := range array {
		if val < min {
			min = val
		}
	}
	return min
}