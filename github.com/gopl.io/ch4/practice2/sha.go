package main

import(
	"fmt"
	"flag"
	"crypto/sha256"
	"crypto/sha512"
	"strings"
)

var sha3 = flag.Bool("3", false, "print sha384" )
var sha5 = flag.Bool("5", false, "print sha512")

func main(){
	flag.Parse()

	arg := strings.Join(flag.Args(), " ")

	if *sha3{
		fmt.Printf("%x\n%s",sha512.Sum384([]byte(arg)),arg)
	}else if *sha5{
		fmt.Printf("%x\n%s",sha512.Sum512([]byte(arg)),arg)
	}else{
		fmt.Printf("%x\n",sha256.Sum256([]byte(arg)))
	}
	fmt.Println(flag.Args())
}