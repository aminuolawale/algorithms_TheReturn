https://www.hackerrank.com/challenges/encryption/problem

package main


import (
	"fmt"
	"math"
)


func main(){
	fmt.Println(encryption("chillout"))
}



 
func encryption(s string) string {
	s = cleanString(s)
	size := len(s)
	row := int(math.Floor(math.Sqrt(float64(size))))
	col := int(math.Ceil(math.Sqrt(float64(size))))
	for row * col < size {
		row ++
	}
	mtx := []string{}
	ind := 0
	for i:=0; i< int(row); i++{
		rs:= ""
		for j:=0; j< int(col); j++ {
			rs += string(s[ind])
			ind +=1
			if ind == len(s){
				mtx = append(mtx, rs)
				rs = ""
				break
			}
		}
		mtx = append(mtx, rs)
		rs = ""
	}
	fmt.Println(mtx)
	cs := ""
	for j:=0; j < int(col); j++{
		for i:=0; i < int(row); i++{
			cs += string(mtx[i][j])
			if i < row -1 && len(mtx[i+1]) < j+1{
				break
			}
		}
		cs += " "
	}
	return cs
}

func cleanString(s string) string {
	result := ""
	for i:=0; i<len(s); i++ {
		if s[i] != ' ' {
			result +=  string(s[i])
		}
	}
	return result
}