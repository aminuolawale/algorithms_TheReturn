//https://www.hackerrank.com/challenges/organizing-containers-of-balls/problem


package main


import (
	"fmt"
)

func main(){
	ans := organizingContainers([][]int32{[]int32{1,3,1}, []int32{2,1,2}, []int32{3,3,3}})
	fmt.Println(ans)
}

// Algorithm summary: For each type of ball, there should exist a container that contains the same number of balls

func organizingContainers(container [][]int32) string {
	size := len(container[0])
	containerSizes :=[]int{}
	for _, c := range container {
		containerSizes = append(containerSizes, sliceSum(c))
	} 

	for j:=0; j< size; j++{
		count := 0
		for i:=0; i < size; i++ {
			count += int(container[i][j])
		}
		contain, ind := contains(containerSizes, count)
		if !contain {
			return "Impossible"
		}
		containerSizes = slicePop(containerSizes, ind)
	}
	return "Possible"
}

func slicePop(slice []int, index int) []int{
	slice[index] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}

func contains( array []int, val int) (bool, int) {
	for index, a := range array {
		if a == val {
			return true, index
		}
	}
	return false, -1
}

func sliceSum(slice []int32) int {
	result :=0
	for _, val :=range slice {
		result += int(val)
	}
	return result
}