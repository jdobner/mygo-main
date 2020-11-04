package main

import (
	"github.com/jdobner/mygo/doit"
	sub "github.com/jdobner/mygo/doit/subsub"
	//"mygo/doit"
	"fmt"
)

func main() {
	fmt.Println(doit.Doit(), sub.Version())
}
