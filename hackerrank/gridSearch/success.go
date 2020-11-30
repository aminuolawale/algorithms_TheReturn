package main

import (
	"fmt"
)


func main(){

	g:=[]string{"123456",
		"567890",
		"234567",
		"194729"}
	p:= []string{"34", "78"}
	fmt.Println(gridSearch(g,p))
}


func gridSearch(G []string, P []string) string {
    ans:="NO"
    for i, line := range G {
		bb:= line[:len(line)-len(P[0])+1] 
        for j, _ := range bb{
			// fmt.Println(line[j:j+len(P[0])],P[0])
            if line[j:j+len(P[0])]==P[0]{
				ans = "YES"
				fmt.Println("got here")
                for q, m := range P{
                    fmt.Println(m, G[i+q][j:j+len(P[0])])
                    if m != G[i+q][j:j+len(P[0])] {
						ans =  "NO"
                    }
				}
				if ans == "YES" {
					return ans
				}
            }
        }
    }
    return ans
}