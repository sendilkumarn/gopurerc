package concurrent

import (
	"fmt"
	"gopurerc/pkg/common"
)

var roots []*Object

var cycleBuffer [][]*Object

func Increment(o *Object) {
	fmt.Println("incrementing", o.String())
	o.Rc = o.Rc + 1
	scanBlack(o)
}

func Decrement(o *Object) {
	fmt.Println("decrementing", o.String())
	o.Rc = o.Rc - 1
	if o.Rc == 0 {
		release(o)
	} else if !IsAcyclic(o) {
		possibleRoot(o)
	}
}

func release(o *Object) {
	fmt.Println("release", o.String())
	for _, child := range o.Children {
		Decrement(child)
	}
	o.Color = common.Color.BLACK
	if !o.Buffered {
		free(o)
	}
}

func possibleRoot(o *Object) {
	fmt.Println("possible root")
	scanBlack(o)
	o.Color = common.Color.PURPLE
	fmt.Println("ColorChange", o.Color)
	if !o.Buffered {
		o.Buffered = true
		roots = append(roots, o)
	}
}

func CollectCycles() {
	fmt.Println("collect Cycles")
	freeCycles()
	findCycles()
	sigmaPreparation()
}

func findCycles() {
	fmt.Println("find Cycles")
	markRoots()
	scanRoots()
	collectRoots()
}

func markRoots() {
	fmt.Println("markRoots")
	sn := 0
	for si := 0; si < len(roots); si++ {
		s := roots[si]
		if s.Color == common.Color.PURPLE && s.Rc > 0 {
			//fmt.Println(len(roots))
			markGray(s)
			roots[sn] = s
			sn = sn + 1
		} else {
			s.Buffered = false
			if s.Rc == 0 {
				free(s)
			}
		}
	}
	roots = roots[:sn]
}

func scanRoots() {
	fmt.Println("scanRoots")
	si := 0
	sk := len(roots)
	for si < sk {
		s := roots[si]
		scan(s)
		si = si + 1
	}
}

func collectRoots() {
	fmt.Println("collectRoots")
	si := 0
	sk := len(roots)

	for si < sk {
		s := roots[si]
		if s.Color == common.Color.WHITE {
			var currentCycle []*Object
			gc := collectWhite(s, currentCycle)
			cycleBuffer = append(cycleBuffer, gc)
		} else {
			s.Buffered = false
		}
		si++
	}
	roots = []*Object{}
}

func scanBlack(o *Object) {
	fmt.Println("Scan black")
	if o.Color != common.Color.BLACK {
		o.Color = common.Color.BLACK
		for _, child := range o.Children {
			scanBlack(child)
		}
	}
}

func markGray(o *Object) {
	fmt.Println("marking gray", o.String())
	if o.Color != common.Color.GRAY {
		o.Color = common.Color.GRAY
		fmt.Println("Color change ", o.Color)
		o.Crc = o.Rc
		for _, child := range o.Children {
			markGray(child)
			if child.Crc > 0 {
				child.Crc = child.Crc - 1
			}
		}
	}
}

func scan(o *Object) {
	fmt.Println("scanning")
	if o.Color == common.Color.GRAY {
		if o.Crc > 0 {
			scanBlack(o)
		} else {
			o.Color = common.Color.WHITE
			for _, child := range o.Children {
				scan(child)
			}
		}
	}
}

func collectWhite(o *Object, cc []*Object) []*Object {
	fmt.Println("collecting white")
	if o.Color == common.Color.WHITE {
		o.Color = common.Color.ORANGE
		o.Buffered = true
		cc = append(cc, o)
		for _, child := range o.Children {
			cc = collectWhite(child, cc)
		}
	}
	return cc
}

func free(o *Object) {
	fmt.Println("freeing count", o.String())
	Count = Count - 1
	o.Freed = true
}

func freeCycles() {
	fmt.Println("free cycles")
	last := len(cycleBuffer) - 1
	for i := last; i >= 0; i-- {
		if sigmaDeltaTest(cycleBuffer[i]) {
			freeCycle(cycleBuffer[i])
		} else {
			refurbish(cycleBuffer[i])
		}
	}

	cycleBuffer = [][]*Object{}
}

func sigmaDeltaTest(o []*Object) bool {
	externRC := 0
	for _, n := range o {
		if n.Color != common.Color.ORANGE {
			return false
		}
		externRC = externRC + n.Crc
	}
	return externRC == 0
}

func freeCycle(o []*Object) {
	fmt.Println("free cycle", o)

	for _, n := range o {
		n.Color = common.Color.RED
		fmt.Println("Color Change", n.Color)
	}

	for _, n := range o {
		for _, c := range n.Children {
			cyclicDecrement(c)
		}
	}

	for _, n := range o {
		free(n)
	}

}

func cyclicDecrement(o *Object) {
	fmt.Println("cyclic Decrement")
	if o.Color != common.Color.RED {
		if o.Color == common.Color.ORANGE {
			o.Rc = o.Rc - 1
			o.Crc = o.Crc - 1
		} else {
			Decrement(o)
		}
	}
}

func refurbish(o []*Object) {
	fmt.Println("refurbish")
	first := true
	ni := 0
	nk := len(o)

	for ni < nk {
		n := o[ni]
		if (first && n.Color == common.Color.ORANGE) || n.Color == common.Color.PURPLE {
			n.Color = common.Color.PURPLE
			roots = append(roots, n)
		} else {
			n.Color = common.Color.BLACK
			n.Buffered = false
		}
		first = false
		ni = ni + 1
	}
}

func sigmaPreparation() {
	fmt.Println("sigma preparation")
	for _, c := range cycleBuffer {
		for _, n := range c {
			n.Color = common.Color.RED
			n.Crc = n.Rc
		}

		for _, n := range c {
			for _, m := range n.Children {
				if m.Color == common.Color.RED && m.Crc > 0 {
					m.Crc = m.Crc - 1
				}
			}
		}

		for _, n := range c {
			n.Color = common.Color.ORANGE
		}
	}
}

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
	return !o.cyclesTo(o, make(map[*Object]bool))
}

func (o *Object) cyclesTo(other *Object, except map[*Object]bool) bool {
	if except[o] == true {
		return false
	}

	except[o] = true
	for _, child := range o.Children {
		if child == other || child.cyclesTo(other, except) {
			return true
		}
	}

	return false
}

func (o *Object) Add(s *Object) {
	Increment(s)
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
		Decrement(&s)
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
