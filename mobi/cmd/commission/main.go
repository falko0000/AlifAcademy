package main

import (
	"fmt"
	"mobi/pkg/commission"
)

func main() {
	result := commission.Calculate(9_999_99)
	fmt.Println(result)
}
