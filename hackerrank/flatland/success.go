package main

import ("fmt"
		"sort"
		"math"
)


func main(){
	d:=flatlandSpaceStations(95, []int32{68, 81, 46, 54, 30, 11, 19, 23, 22, 12, 38, 91, 48, 75, 26, 86, 29, 83, 62})
	fmt.Println(d)
}




func flatlandSpaceStations(n int32, c []int32) int32 {
	distances:=[]int32{}
	sort.Slice(c, func(i, j int) bool { return c[i] < c[j] })
	var ind int32 
	var ans float64
	for int(ind) <len(c) {
		if ind ==0 && c[ind] ==0 {
			ans = 0
		} else if ind == 0 {
			ans = float64(c[ind])
		}else {
			ans = math.Ceil((float64(c[ind])- float64(c[ind-1]) -1.0)/2.0)
		}
		distances = append(distances, int32(ans))
		ind+=1
	}	
	if c[ind-1] != n-1{
		ans = float64(n-1-c[ind-1])
		distances = append(distances, int32(ans))
	}
	return maxx(distances)
}


func maxx (arr []int32)int32{
	ans :=arr[0]
	for _, v :=range arr{
		if v > ans{
			ans = v
		}
	}
	return ans
}



