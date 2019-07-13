package test

import (
	"gopurerc/pkg/object"
	"testing"
)

func customRetain(o *object.Object) {

	//b := retain(&o)

}

func TestLocals(*testing.T) {
	o := object.New("1")

	customRetain(&o)
	someFunction(&o)
	release(&o)
	collect()
	check()
}

