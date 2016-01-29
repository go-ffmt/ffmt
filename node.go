package ffmt

import (
	"bytes"
	"strings"
)

type head struct {
	node
	max int
}

const colSym = ": "

func spac(depth int) string {
	b := []byte{}
	if depth > 0 {
		for i := 0; i != depth; i++ {
			b = append(b, ' ')
		}
	}
	return string(b)
}

func spacing(depth int) string {
	return "\n" + spac(depth)
}
func (n *head) String() string {
	buf := bytes.Buffer{}
	n.strings(0, &buf)
	s := buf.String()
	if len(s) >= 2 {
		return s[2:]
	}
	return ""
}

type node struct {
	parent *node
	child  *node
	next   *node
	value  string
	colon  int
}

func (n *node) tablePos() {
	ms := []int{}
	sum := 0
	b := n
	max := 0
	for next := b; next != nil; next = next.next {
		if next.colon > 0 || next.child != nil {
			return
		}
		ll := len(next.value)
		ms = append(ms, ll)
		sum += ll
		if ll > max {
			max = ll
		}
	}

	if max < 10 {
		n.merge(9, ms)
	} else if max < 18 {
		n.merge(5, ms)
	} else if max < 30 {
		n.merge(3, ms)
	}
}

func (n *node) merge(m int, ms []int) {
	l := len(ms)
	col := 0
	for i := 0; i != m; i++ {
		z := m - i
		if l > z && l%z == 0 {
			col = z
			break
		}
	}
	if col > 1 {
		n.mergeNextSize(col, ms)
	}
}

func (n *node) mergeNextSize(s int, ms []int) {
	lmax := make([]int, s)
	for j := 0; j != s; j++ {
		for i := 0; i*s < len(ms); i++ {
			b := i*s + j
			if ms[b] > lmax[j] {
				lmax[j] = ms[b]
			}
		}
	}
	for i := 1; i < len(lmax); i++ {
		lmax[i] += lmax[i-1]
	}
	for next := n; next != nil; next = next.next {
		for i := 0; i < s-1 && next.next != nil; i++ {
			next.mergeNext(lmax[i])
		}
	}
}

func (n *node) mergeNext(max int) {
	n.value += spac(max - len(n.value))
	n.value += n.next.value
	n.next = n.next.next
}

func (n *node) colonPos() {
	b := n
	for b != nil {
		m := 0
		for next := b; next != nil; next = next.next {
			if next.colon > m {
				m = next.colon
			}
			if next.child != nil {
				break
			}
		}
		for next := b; next != nil; next = next.next {
			if next.colon > 0 && m-next.colon > 0 {
				next.value = strings.Replace(next.value, colSym, colSym+spac(m-next.colon), 1)
			}
			b = next.next
			if next.child != nil {
				break
			}
		}
	}
	return
}

func (n *node) strings(d int, buf *bytes.Buffer) {
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
	ss := strings.Split(a, "\n")
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
				e.child.tablePos()
				e = e.toNext()
			}
		}
		depth = d
		if d > 0 {
			d--
		}
		e.value = b[d:]
		e.colon = strings.Index(e.value, colSym)
		if max := d + len(e.value); max > o.max {
			o.max = max
		}
	}
	return o
}
