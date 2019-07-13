package test

import (
	"gopurerc/pkg/object"
	"testing"
)

func TestArray(*testing.T) {
	createArray()
	collect()
	check()
}

func createArray() {
	o := object.New("array")
	e1 := object.New("element1")
	o.Add(&e1)
	e2 := object.New("element2")
	o.Add(&e2)
	e3 := object.New("element3")
	o.Add(&e3)

	retain(&o)
	someFunction(&o)
	release(&o)
}

func someFunction(o *object.Object) {
	defer release(o)
	retain(o)
}
