package main 

import (
	"fmt"
)

func main(){
	ans:= fairRations([]int32{1,2})
	fmt.Println(ans)
}



func fairRations(B []int32) int32 {
	var count int32
	for index, val := range B {
		if index == len(B) -1 {
			if !isEven(val){
				fmt.Println("NO")
			}
			break
		}
		if !isEven(val){
			B[index]+=1
			B[index+1]+=1
			count +=2
		}
	}
	fmt.Println(count)
}



func isEven(val int32)bool{
	return val%2 ==0
}