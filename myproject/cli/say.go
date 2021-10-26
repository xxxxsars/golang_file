package main

import (
	"example.com/rico/myproject/greeting"
	"fmt"
	"golang.org/x/example/stringutil"
)

func main() {
	greeting.Say("Hello")
	fmt.Println(stringutil.Reverse("!selpmaxe oG ,olleH"))
	b := [2]int{2,4}
	fmt.Println(b)
}