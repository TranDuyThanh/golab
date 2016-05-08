package golab

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"runtime"

	"github.com/maruel/panicparse/stack"
)

func CatchPanic() {
	if r := recover(); r != nil {
		err, ok := r.(error)
		if !ok {
			err = fmt.Errorf("pkg: %v", r)
		}
		trace := make([]byte, 8192)
		runtime.Stack(trace, true)
		fmt.Println(err.Error())

		in := bytes.NewBufferString(string(trace))
		goroutines, err := stack.ParseDump(in, os.Stdout)
		if err != nil {
			return
		}
		p := &stack.Palette{}
		buckets := stack.SortBuckets(stack.Bucketize(goroutines, 8192))
		srcLen, pkgLen := stack.CalcLengths(buckets, false)

		for _, bucket := range buckets {
			io.WriteString(os.Stdout, p.BucketHeader(&bucket, false, len(buckets) > 1))
			io.WriteString(os.Stdout, p.StackLines(&bucket.Signature, srcLen, pkgLen, false))
		}
	}
}
