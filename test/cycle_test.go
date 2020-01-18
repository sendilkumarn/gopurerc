package test

import (
	"gopurerc/pkg/concurrent"
	"testing"
)

// Cycle

func TestCycle(*testing.T) {
	createCycle()
	collect()
	check()
}

func createCycle() {
	o := concurrent.New("self")
	o.Add(&o)

	retain(&o)
	release(&o)
}

// Cycle 1

func TestCycle1(*testing.T) {
	createCycle1()
	collect()
	check()

}

func createCycle1() {
	o := concurrent.New("self")
	o.Add(&o)

	doRoutine(&o)
}

func doRoutine(o *concurrent.Object) {
	retain(o)
	someFunction(o)
	release(o)
}

// Cycle 2

func TestCycle2(*testing.T) {
	createCycle2()
	collect()
	check()
}

func createCycle2() {
	o := concurrent.New("array")
	t := concurrent.New("element")
	t.Add(&o)
	o.Add(&t)
	doRoutine(&o)
}

// Cycle 3

func TestCycle3(*testing.T) {
	createCycle3()
	collect()
	check()

}

func createCycle3() {
	o := concurrent.New("level1")
	l2 := concurrent.New("level2")
	o.Add(&l2)
	l3 := concurrent.New("level3")
	l4 := concurrent.New("level4")
	o.Add(&l3)
	o.Add(&l4)
	o.Add(&o)
	doRoutine(&o)
}

// Cycle 4

func TestCycle4(*testing.T) {
	createCycle4()
	collect()
	check()

}

func createCycle4() {
	outer := concurrent.New("outer")
	o := concurrent.New("level1")
	l2 := concurrent.New("level2")
	o.Add(&l2)
	l3 := concurrent.New("level3")
	l4 := concurrent.New("level4")
	o.Add(&l3)
	o.Add(&l4)
	o.Add(&o)
	outer.Add(&o)
	doRoutine(&outer)
}

// Cycle 5

func TestCycle5(*testing.T) {
	createCycle5()
	collect()
	check()

}

func createCycle5() {
	s := concurrent.New("level0")
	t := concurrent.New("level1")
	u := concurrent.New("level2")
	v := concurrent.New("level3")

	v.Add(&t)
	v.Add(&u)
	u.Add(&v)
	u.Add(&u)
	u.Add(&s)
	t.Add(&u)
	t.Add(&v)
	s.Add(&t)

	doRoutine(&s)
}

// Cycle 6

func TestCycle6(*testing.T) {
	createCycle6()
	collect()
	check()

}

func createCycle6() {
	o := concurrent.New("outer")
	s := concurrent.New("level1")
	l2 := concurrent.New("level2")
	s.Add(&l2)
	s.Add(&s)
	inner := concurrent.New("inner")
	s.Add(&inner)
	o.Add(&s)

	doRoutine(&o)
}

// Cycle 7

func TestCycle7(*testing.T) {
	createCycle7()
	collect()
	check()

}

func createCycle7() {
	var s1, s2, t1, t2 concurrent.Object

	cycle1 := concurrent.New("a:outer")
	s1 = concurrent.New("a:level1")
	t1 = concurrent.New("a:level2")
	t1.Add(&s1)
	inner1 := concurrent.New("a:inner")
	t1.Add(&inner1)
	s1.Add(&t1)
	cycle1.Add(&s1)

	cycle2 := concurrent.New("b:outer")
	s2 = concurrent.New("b:level1")
	t2 = concurrent.New("b:level2")
	t2.Add(&s2)
	inner2 := concurrent.New("b:inner")
	t2.Add(&inner2)
	s2.Add(&t2)
	cycle2.Add(&s2)

	cycle3 := concurrent.New("c:level1")
	l2 := concurrent.New("c:level2")
	l3 := concurrent.New("c:level3")
	l4 := concurrent.New("c:level4")

	l4.Add(&cycle3)
	l3.Add(&l4)
	l2.Add(&l3)
	cycle3.Add(&l2)

	cycle1.Add(&cycle2)
	t1.Add(&cycle2)
	cycle1.Add(&cycle3)
	cycle2.Add(&cycle1)

	t2.Add(&cycle1)
	t2.Add(&cycle3)

	retain(&cycle1)

	concurrent.CheckAlive(&cycle1, map[interface{}]bool{})
	concurrent.CheckAlive(&cycle2, map[interface{}]bool{})
	concurrent.CheckAlive(&cycle3, map[interface{}]bool{})
	concurrent.CheckAlive(&t1, map[interface{}]bool{})
	concurrent.CheckAlive(&t2, map[interface{}]bool{})

	collect()
	collect()
	collect()

	cycle1.Add(&cycle2)
	t1.Add(&cycle2)
	cycle1.Add(&cycle3)
	cycle2.Add(&cycle1)
	t2.Add(&cycle1)
	t2.Add(&cycle3)
	collect()

	collect()
	cycle3.Add(&t1)
	collect()
	cycle3.Add(&t2)
	collect()

	collect()
	collect()
	collect()

	concurrent.CheckAlive(&cycle1, map[interface{}]bool{})
	concurrent.CheckAlive(&cycle2, map[interface{}]bool{})
	concurrent.CheckAlive(&cycle3, map[interface{}]bool{})
	concurrent.CheckAlive(&t1, map[interface{}]bool{})
	concurrent.CheckAlive(&t2, map[interface{}]bool{})

	release(&cycle1)
}
