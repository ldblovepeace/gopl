package main

import (
	"fmt"
	// "github.com/gopl.io/ch2/tempconv"
	"github.com/gopl.io/ch2/popcount"
	// "os"
)

func main(){
	// var c tempconv.Celsius = 100
	// c := tempconv.Celsius(100)
	// f := tempconv.CToF(c)

	// fmt.Println(tempconv.String(c))
	// fmt.Printf("%g\n",f)
	// fmt.Println(tempconv.CToT(tempconv.BoilingC))

	fmt.Println(popcount.PopCount(1023))
}