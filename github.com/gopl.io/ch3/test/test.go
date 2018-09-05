package main

import(
	"github.com/gopl.io/ch3/comma"
	"github.com/gopl.io/ch3/comma2"
	"fmt"
)

func main(){
	i := "123456978"
	fmt.Println(comma.Comma(i))
	fmt.Println(comma2.Comma(i))
}