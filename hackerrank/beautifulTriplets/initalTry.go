package main

import (
	"fmt"
)

// worked but failed two out of 15 tests due to timeout

func main(){
	theArr := []int32{2,2,3,4,5}
	ans := beautifulTriplets(1, theArr)
	fmt.Println(ans)
}


func beautifulTriplets(d int32, arr []int32) int32 {
	index := 0
	countToTriplets :=1
	var tripletsCount int32 = 0
	superIndex := 0
	for{
		// check whether array from arr[index+1] to end contains value +d
		if superIndex > len(arr) -3{
			break
		}
		isThere, newIndex := contains(arr, arr[index] +d, index+1)
		for isThere{
			fmt.Println(arr[index],arr[newIndex])
			index = newIndex
			countToTriplets += 1
			isThere, newIndex = contains(arr, arr[index] +d, index+1)
		}
		if countToTriplets >=3 {
			fmt.Println("============", arr[index])
			countToTriplets =1
			tripletsCount +=1
		}
		superIndex +=1
		index = superIndex
		countToTriplets = 1
	}
	return tripletsCount
}


func contains(arr []int32, value int32, start int)(bool, int){
	for index, v := range arr {
		if index >= start && v == value {
			return true, index
		}
	}
	return false, -1
}


