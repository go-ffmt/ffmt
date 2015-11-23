package ffmt

import "strings"

type node struct {
	parent *node
	Value  string
	Child  *node
	Next   *node
}

func (n *node) String() string {
	return n.strings(0)
}

func (n *node) strings(d int) (ret string) {
	ret += spacing(d)
	ret += n.Value

	if n.Child != nil {
		ret += n.Child.strings(d + 1)
	}
	if next := n.Next; next != nil {
		ret += next.strings(d)
	}
	return
}

func (n *node) toParent() (e *node) {
	return n.parent
}

func (n *node) toChild() (e *node) {
	if n.Child == nil {
		n.Child = &node{
			parent: n,
		}
	}

	return n.Child
}

func (n *node) toNext() (e *node) {
	if n.Next == nil {
		n.Next = &node{
			parent: n.parent,
		}
	}
	return n.Next
}

func getDepth(a string) (i int) {
	for ; i != len(a); i++ {
		if a[i] == byte(' ') {
			continue
		} else if a[i] == byte(',') {
			i++
		}
		break
	}
	return
}

func stringToNode(a string) (o *node) {
	ss := strings.Split(Fmt(a), "\n")
	depth := 0
	o = &node{}
	e := o
	for i := 0; i != len(ss); i++ {
		b := ss[i]

		d := getDepth(b)
		if d == depth {
			e = e.toNext()
		} else if d > depth {
			e = e.toChild()
		} else {
			e = e.toParent()
			if e != nil {
				e = e.toNext()
			}
		}
		depth = d
		e.Value = b
	}
	return o.Next
}
