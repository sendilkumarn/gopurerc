package test

import (
	"fmt"
	"gopurerc/pkg/concurrent"
	"gopurerc/pkg/object"
)

func retain(o *object.Object) {
	if o.Name != "" {
		concurrent.Increment(o)
	}
}

func release(o *object.Object) {
	if o.Name != "" {
		concurrent.Decrement(o)
	}
}

func collect() {
	concurrent.CollectCycles()
	concurrent.CollectCycles()
}

func check() {
	fmt.Println("count is ", object.Count)
	if object.Count != 0 {
		panic("error ")
	}
}
