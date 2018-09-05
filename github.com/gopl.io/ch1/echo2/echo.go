package main

import(
	"fmt"
	"os"
)

func main()  {
	// s, sep := "", ""
	// for _, arg:=range os.Args[1:]{
	// 	s += sep + arg
	// 	sep = ""
	// }

	// fmt.Println(s)
	// fmt.Println(os.Args[0]);
	for i:=1; i<len(os.Args); i++ {
		fmt.Printf("%d\t%s\n",i,os.Args[i])
	}
}