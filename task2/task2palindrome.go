/*Write a Go function that takes a string as input and checks 
whether it is a palindrome or not. A palindrome is a word, phrase,
 number, or other sequence of characters that reads the same forward 
 and backward (ignoring spaces, punctuation, and capitalization).*/

 package main
 import (
	"bufio"
	"fmt"
	"os"
	// "strconv"
	"strings"
	"unicode"
)

func main(){
reader := bufio.NewReader(os.Stdin)
input,_ := reader.ReadString('\n')
input = strings.TrimSpace(input)
input  = strings.ToLower(input)
var arr []string
for _,ch := range input{
	if unicode.IsLetter(ch){
		s:= string(ch)
			arr = append(arr,s)
	}
}
// fmt.Println((arr))
l:=0
r:=len(arr) -1
for ;l<=r; {
	if arr[l]!=arr[r]{
		fmt.Println(" The string is not palindrome")
		os.Exit(0)
	}
	l+=1
	r-=1
}
fmt.Println(" The string is palindrome")
}
