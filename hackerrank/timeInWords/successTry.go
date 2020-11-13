package main

import (
	"fmt"
)

func main(){
	fmt.Println(timeInWords(1, 1))
}




func timeInWords(h int32, m int32) string {
	if m == 0 {
		return translate(h) + " o' clock"
	}
	direction := "past "
	minuteTranslation := ""
	minutes := "minutes "
	if m == 15 {
		minuteTranslation = "quarter "
		minutes = ""
	}else if m == 30 {
		minuteTranslation = "half "
		minutes = ""
	}else if m == 45 {
		direction = "to "
		h+=1
		minuteTranslation = "quarter "
		minutes = ""
	}else if m >30 {
		direction = "to "
		h+=1
		m = 60 - m
		minuteTranslation = translate(m) + " "
	} else {
		direction = "past "
		minuteTranslation = translate(m) + " "
	}
	if m == 1 && minutes != ""{
		minutes = "minute "
	}
	translation := minuteTranslation + minutes + direction + translate(h)
	return translation
}	


func translate(num int32)string{
	table := make(map[int32]string)
	table[1] = "one"
	table[2] = "two"
	table[3] = "three"
	table[4] = "four"
	table[5] = "five"
	table[6] = "six"
	table[7] = "seven"
	table[8] = "eight"
	table[9] = "nine"
	table[10] = "ten"
	table[11] = "eleven"
	table[12] = "twelve"
	table[13] = "thirteen"
	table[14] = "fourteen"
	table[15] = "fifteen"
	table[16] = "sixteen"
	table[17] = "seventeen"
	table[18] = "eighteen"
	table[19] = "nineteen"
	table[20] = "twenty"

	translation := table[num]
	if num !=0 && translation == "" {
		intTemp1 := (num /10) * 10
		temp1 := table[intTemp1]
		intTemp2 := num - intTemp1
		temp2 := table[intTemp2]
		translation = string(temp1) + " " + string(temp2)
	}
	return translation

}