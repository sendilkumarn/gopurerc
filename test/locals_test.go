package test

import (
	"gopurerc/pkg/concurrent"
	"testing"
)

func customRetain(o *concurrent.Object) {

	//b := retain(&o)

}

func TestLocals(*testing.T) {
	o := concurrent.New("1")

	customRetain(&o)
	someFunction(&o)
	release(&o)
	collect()
	check()
}
