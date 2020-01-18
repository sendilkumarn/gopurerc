package test

import (
	"fmt"
	"gopurerc/pkg/concurrent"
)

func retain(o *concurrent.Object) {
	if o.Name != "" {
		concurrent.Increment(o)
	}
}

func release(o *concurrent.Object) {
	fmt.Println("releasing by decrementing")
	if o.Name != "" {
		concurrent.Decrement(o)
	}
}

func collect() {
	concurrent.CollectCycles()
	concurrent.CollectCycles()
}

func check() {
	fmt.Println("count is ", concurrent.Count)
	if concurrent.Count != 0 {
		panic("error ")
	}
}
