package main


import (
	"fmt"
	"sort"
)


func main(){
	fmt.Println(stones(73, 25, 25))
}



func stones(n int32, a int32, b int32) []int32 {

	ans :=[]int32{}
    var p int32 = n -1
	var i int32 =0
	if a == b {
		return []int32{p* a}
	}
    for p>=0{
        d := p * a + i * b
        ans = append(ans, d)
        p -=1
        i+=1
	}
	sort.Slice(ans, func(i, j int) bool { return ans[i] < ans[j] })
    return ans
}
