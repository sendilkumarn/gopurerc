package test

import (
	"gopurerc/pkg/concurrent"
	"testing"
)

func TestArray(*testing.T) {
	createArray()
	collect()
	check()
}

func createArray() {
	o := concurrent.New("array")
	e1 := concurrent.New("element1")
	o.Add(&e1)
	e2 := concurrent.New("element2")
	o.Add(&e2)
	e3 := concurrent.New("element3")
	o.Add(&e3)

	retain(&o)
	someFunction(&o)
	release(&o)
}

func someFunction(o *concurrent.Object) {
	defer release(o)
	retain(o)
}
