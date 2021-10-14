package main

import (
	"fmt"
	"github.com/WahidinAji/helper-say-hello"
)

func main() {
	fmt.Println("hello")
	say_hello := helper.SayHello()
	fmt.Println(say_hello)
}

/*
remember this one, if the package name (name of function or variable) using Capital in the first words, then it can access on other package.
if the package name using lowercase, then it can't be accessed by other package
 */