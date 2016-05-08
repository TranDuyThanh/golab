package main

import (
	"fmt"
	"github.com/TranDuyThanh/golab"
)

func main() {

	defer golab.CatchPanicAndExec(func() {
		fmt.Println("--------------panic--------------")
	})

	panic("hihi")
}
