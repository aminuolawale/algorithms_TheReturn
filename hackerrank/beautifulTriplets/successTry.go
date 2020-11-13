package main

import (
	"fmt"
)

//

func main(){
	theArr := []int32{1 ,2 ,4 ,5 ,7 ,8 ,10}
	ans := beautifulTriplets(3, theArr)
	fmt.Println(ans)
}


func beautifulTriplets(d int32, arr []int32) int32 {
	count := 0
	for index, val := range arr {
		if contains2(arr, val, d, index+1){
			count += 1
		}
	}
	return int32(count)
}


func contains2(arr []int32, value int32, d int32, start int)bool{
	count:=0
	for index, v := range arr {
		if index >= start && v == value +d {
			count += 1
			value +=d
			if count >=2 {
				return true
			}
		}
	}
	return false
}


