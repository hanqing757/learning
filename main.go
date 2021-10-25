package main

import (
	"fmt"
	"learning/gopl"
)

func main()  {
	//d := gopl.TopSort()
	b := gopl.HasRing()
	fmt.Println(b)
}