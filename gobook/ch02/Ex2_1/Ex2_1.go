package main

import (
	"fmt"
	"jotranen/golang/gobook/ch02/Ex2_1/lib"
)

func main() {
	fmt.Printf("Brrrr! %v\n", tempconv.AbsoluteZeroC)
	fmt.Printf("Brrrr! %v\n", tempconv.CToK(tempconv.AbsoluteZeroC))

}
