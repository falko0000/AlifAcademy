package main

import (
	"fmt"
	"flag"
	"strings"
)

func main()  {
	flag.Parse()
	massage := `Добро пожаловать, {name}! Ваш код доступа: {code}.`
	name := flag.Arg(0)
	code := flag.Arg(1)
	massage = strings.ReplaceAll(massage, "{name}", name)
	massage = strings.ReplaceAll(massage, "{code}", code)
	fmt.Println(massage)
}