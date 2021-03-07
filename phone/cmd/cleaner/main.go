package main

import (
	"fmt"
	"flag"
	"strings"
)

func main()  {
	flag.Parse()
	number := flag.Arg(0)
	number = strings.ReplaceAll(number, " ", "")
	number = strings.ReplaceAll(number, "(", "")
	number = strings.ReplaceAll(number, ")", "")
	number = strings.ReplaceAll(number, "-", "")
	fmt.Println(number)
}