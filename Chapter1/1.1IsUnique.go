//Determines if a string has all unique characters
package main

import "fmt"

func isUnique(s string) bool{
	c := make(map[string]int)
	for _,i:= range(s) {
		if c[string(i)] == 0 {
			c[string(i)] = 1
		} else{
			return false
		}
	}
	return true
}

func main() {
	var str string
	fmt.Println("Enter a string:")
	fmt.Scan(&str)
	fmt.Println(isUnique(str))
}