package test

import (
	"gopurerc/pkg/concurrent"
	"testing"
)

func TestSimple(*testing.T) {
	o := concurrent.New("1")
	object := concurrent.New("2")
	o.Add(&object)

	retain(&o)
	someFunction(&o)
	release(&o)
	collect()
	check()
}
