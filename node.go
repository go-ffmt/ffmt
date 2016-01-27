package ffmt

import (
	"bytes"
	"strings"
)

type head struct {
	node
	max int
}

func (n *head) String() string {
	buf := bytes.Buffer{}
	n.strings(0, &buf)
	return moveSpac(buf.String())
}

type node struct {
	parent *node
	child  *node
	next   *node
	value  string
	colon  int
}

func (n *node) colonPos() {
	m := 0
	for next := n; next != nil; next = next.next {
		if next.colon > m {
			m = next.colon
		}
		if next.child != nil {
			break
		}
	}
	for next := n; next != nil; next = next.next {
		if next.colon > 0 && m-next.colon > 0 {
			next.value = strings.Replace(next.value, ": ", ": "+spac(m-next.colon), 1)
		}
		if next.child != nil {
			break
		}
	}
	return
}

func (n *node) strings(d int, buf *bytes.Buffer) (ret string) {
	buf.WriteString(spacing(d))
	buf.WriteString(n.value)
	if n.child != nil {
		n.child.strings(d+1, buf)
	}
	if next := n.next; next != nil {
		next.strings(d, buf)
	}
	return
}

func (n *node) toParent() (e *node) {
	return n.parent
}

func (n *node) toChild() (e *node) {
	if n.child == nil {
		n.child = &node{
			parent: n,
			child:  nil,
			next:   nil,
		}
	}
	return n.child
}

func (n *node) toNext() (e *node) {
	if n.next == nil {
		n.next = &node{
			parent: n.parent,
			child:  nil,
			next:   nil,
		}
	}
	return n.next
}

func getDepth(a string) int {
	for i := 0; i != len(a); i++ {
		switch a[i] {
		case ' ':

		case ',':
			return i + 1
		default:
			return i
		}
	}
	return 0
}

func stringToNode(a string) (o *head) {
	ss := strings.Split(fmts(a, 0, 0), "\n")
	depth := 0
	o = &head{}
	e := &o.node
	for i := 0; i != len(ss); i++ {
		b := ss[i]
		d := getDepth(b)
		if d == depth {
			e = e.toNext()
		} else if d > depth {
			e = e.toChild()
		} else if d < depth {
			e = e.toParent()
			if e != nil {
				e.child.colonPos()
				e = e.toNext()
			}
		}
		depth = d
		e.value = b
		e.colon = strings.Index(e.value, ": ")
		if max := d + len(e.value); max > o.max {
			o.max = max
		}
	}
	return o
}
