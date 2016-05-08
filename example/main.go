package main

import (
	"fmt"
	"github.com/TranDuyThanh/golab"
)

func main() {

	defer golab.CatchPanic(func() {
		fmt.Println("--------------panic--------------")
	})

	panic("hihi")
}
