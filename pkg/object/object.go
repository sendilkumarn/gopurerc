package object

import (
	"fmt"
	"gopurerc/pkg/common"
	"gopurerc/pkg/concurrent"
)

type Object struct {
	Name      string
	Rc        int
	Crc       int
	Color     string
	Buffered  bool
	Children  []*Object
	Freed     bool
	isAcyclic bool
}

var Count int

func New(name string) Object {
	o := Object{
		Name:     name,
		Rc:       0,
		Crc:      0,
		Color:    common.Color.BLACK,
		Buffered: false,
		Children: []*Object{},
		Freed:    false,
	}
	Count++
	fmt.Println("Creating new object", o.String())
	return o
}

func (o *Object) String() string {
	return fmt.Sprintf("name: %s, Rc: %d, Crc: %d, Color: %s, Children: %d, Count: %d", o.Name, o.Rc, o.Crc, o.Color, len(o.Children), Count)
}

func (o *Object) PrintChild() {
	for _, child := range o.Children {
		fmt.Println(child.String())
	}

}


func IsAcyclic(o *Object) bool {
	return cyclesTo(o, make(map[interface{}]bool))
}

func cyclesTo(other *Object, except map[interface{}]bool) bool {
	if except[other] == true {
		return false
	}

	except[other] = true
	for _, child := range other.Children {
		if child.Name == other.Name || cyclesTo(child, except) {
			return true
		}
	}

	return false
}

func (o *Object) Add(s *Object) {
	concurrent.Increment(s)
	o.Children = append(o.Children, s)
}

func (o *Object) remove(s Object) {
	index := -1
	for i, child := range o.Children {
		if child.Name == s.Name {
			index = i
			break
		}
	}
	if index != -1 {
		o.Children = append(o.Children[:index], o.Children[index+1:]...)
		concurrent.Decrement(&s)
	}
}

func CheckAlive(o *Object, except map[interface{}]bool) bool {
	if except[o] == true {
		return false
	}
	except[o] = true
	if o.Freed || Count == 0 {
		panic("should be alive")
	}

	for _, c := range o.Children {
		CheckAlive(c, except)
	}
	return true
}

func CheckDead(o *Object, except map[interface{}]bool) bool {
	if except[o] == true {
		return false
	}
	except[o] = true
	if !o.Freed {
		panic("should be dead")
	}

	for _, c := range o.Children {
		CheckDead(c, except)
	}
	return true
}
