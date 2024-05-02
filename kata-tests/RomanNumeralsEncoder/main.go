package main

import "fmt"

func main() {
	fmt.Println()
}

func Encoder(nb int) string {
	romanLibrary := make(map[int]string)
	          romanLibrary[1]		 =	"I"
	          romanLibrary[5]		 =	"V"	          
			  romanLibrary[10]	 	=	"X"
	          romanLibrary[50]		 =	"L"
	          romanLibrary[100]		 =	"C"
	          romanLibrary[500]		 =	"D"
	          romanLibrary[1000]	 =	"M"

}