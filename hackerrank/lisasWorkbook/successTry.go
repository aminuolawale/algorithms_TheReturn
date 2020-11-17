package main


import "fmt"


func main (){
	answer:= workbook(10,5, []int32{3, 8, 15, 11, 14, 1, 9, 2, 24, 31})
	fmt.Println(answer)
}	


func workbook(n int32, k int32, arr []int32) int32 {
	var page int32
	var count int32
	for _, questions := range arr {
		pagesOccupied := questions/k
		if questions %k >0 {
			pagesOccupied ++
		} 
		var m int32 = 0
		for i:=page+1; i<= page+pagesOccupied; i++{
			var pageQuestions int32
			if questions < k{
				pageQuestions=questions
			}else {
				pageQuestions =k
			}
			questions-=k
			if  i>=m+1 && i<=m+pageQuestions{
				count ++
			}
			m+=pageQuestions
		}
		page += pagesOccupied
	}
	return count
}


// func getProblems(page int32)[]int32{

// }