package test

import (
	object2 "gopurerc/pkg/object"
	"testing"
)

func TestSimple(*testing.T) {
	o := object2.New("1")
	object := object2.New("2")
	o.Add(&object)

	retain(&o)
	someFunction(&o)
	release(&o)
	collect()
	check()
}
