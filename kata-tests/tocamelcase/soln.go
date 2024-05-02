package main


import (
	"fmt"	
	"unicode"


func main() {
	
	arg1 := "the-stealth-warrior"
	fmt.Println(TocCamelCase(arg1)
}
/*1. Complete the method/function so that it converts dash/underscore delimited words into camel casing.
   The first word within the output should be capitalized only if the original word was capitalized (known as Upper Camel Case,
   also often referred to as Pascal case). The next words should be always capitalized.
  Examples

  "the-stealth-warrior" gets converted to "theStealthWarrior"

  "The_Stealth_Warrior" gets converted to "TheStealthWarrior"

  "The_Stealth-Warrior" gets converted to "TheStealthWarrior"*/

func ToCamelCase(str string) string {
	strC := []rune(str)
	for i, char := range strC {
		if char == '-' || char == '_' && i > 0 {
			strC[i+1] = unicode.ToTitle(strC[i+1])
		}
		strC = append(strC[:i], strC[i+1:]...)
		str = string([]rune(strC))
	}

	return str
}
