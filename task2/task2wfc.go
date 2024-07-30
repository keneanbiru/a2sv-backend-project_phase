/*Write a Go function that takes a string as input and returns
a dictionary containing the frequency of each word in the string.
 Treat words in a case-insensitive manner and ignore punctuation
 marks.*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main(){
reader := bufio.NewReader(os.Stdin)
input,_ := reader.ReadString('\n')
input = strings.TrimSpace(input)
var arr []string
for _, char := range input {
	if unicode.IsPunct(char) {
		continue
	} else {
		arr = append(arr,string(char))
	}
}
// fmt.Println(arr)
s := strings.Join(arr,"")
// fmt. Println(s)
word := strings.Fields(s)
mapp := make(map[string]int)
for  _,ele := range word{
	_,ex := mapp[ele]
	if ex {
		mapp[ele] +=1

	}else {
		mapp[ele] = 1}

}
for key,val:= range mapp{
	fmt.Println(key + " : " + strconv.Itoa(val))
}
}
	
