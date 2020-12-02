package main

import (
	"fmt"
)


func main(){
	fmt.Println(absolutePermutation(10,1))
}

// https://www.hackerrank.com/challenges/bomber-man/problem
func absolutePermutation(n int32, k int32) []int32 {
	var results = make([]int32, n)
	for val:=1; val<= int(n); val++{
		ind1 := int32(val)-k
		ind2 := int32(val) + k
		if ind1 >0 && ind1 <= n && results[ind1-1]==0{
			results[ind1-1]= int32(val)
		} else if ind2 >0 && ind2 <= n  && results[ind2-1]==0{
			results[ind2-1]= int32(val)
		} else {
			return []int32{-1}
		}
	}
	return results
}
