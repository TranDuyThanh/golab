package main

import (
	"github.com/TranDuyThanh/golab"
)

func main() {

	defer golab.CatchPanic()

	panic("hihi")
}
