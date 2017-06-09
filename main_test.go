package main_test

import (
	"fmt"
	"testing"
)

var ay float64 = 45.4343
var qy float64 = 45.4343
var ry float64 = 45.4343

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ry = ay + qy
		ay = qy - ry
	}
}

func BenchmarkProduct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ry = ay * qy
		ay = ry / qy
	}
}

func BenchmarkSumProduct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ry = ay*qy - ry + ay + qy
		ay = (2.0*ry - qy) / (qy + 1.0)
	}
}

type Sign int

func (sign Sign) String() (s string) {
	switch sign {
	case plus:
		s += fmt.Sprintf("+")
	case minus:
		s += fmt.Sprintf("-")
	case multiply:
		s += fmt.Sprintf("*")
	case none:
		s += fmt.Sprintf("{NONE}")
	}
	return s
}

const (
	plus Sign = iota
	minus
	multiply
	none
)

type node struct {
	value string
	left  *node
	right *node
	sign  Sign
}

func newNode(v string) (n *node) {
	if len(v) == 0 {
		panic("Value cannot be zero lenght")
	}
	n = new(node)
	n.sign = none
	n.value = v
	return
}

func op(n1, n2 *node, sign Sign) *node {
	switch sign {
	case plus:
		return Plus(n1, n2)
	case minus:
		return Minus(n1, n2)
	case multiply:
		return Multiply(n1, n2)
	}
	panic("Haven`t function")
}

func Multiply(n1, n2 *node) *node {
	sum := new(node)
	(*sum).sign = multiply
	(*sum).left = n1
	(*sum).right = n2
	return sum
}

func Minus(n1, n2 *node) *node {
	sum := new(node)
	(*sum).sign = minus
	(*sum).left = n1
	(*sum).right = n2
	return sum
}

func Plus(n1, n2 *node) *node {
	sum := new(node)
	(*sum).sign = plus
	(*sum).left = n1
	(*sum).right = n2
	return sum
}

func (n *node) amountMultiplications() (amount int) {
	formula := fmt.Sprintf("%v", n)
	sign := []byte(fmt.Sprintf("%v", multiply))
	for i := range formula {
		if formula[i] == sign[0] {
			amount++
		}
	}
	return
}

func (n *node) simplification() (result *node) {
	defer func() {
		fmt.Printf("result = %v\n", result)
	}()

	// input: empty node
	// for example in left and right for
	// value node
	if n == nil {
		return nil
	}

	// input: value
	if n.sign == none {
		return n
	}

	// input: plus or minus nodes
	// no need anything except simplification
	if n.sign == plus || n.sign == minus {
		a := n.left.simplification()
		b := n.right.simplification()
		return op(a, b, n.sign)
	}
	// Now, root node is multiplication

	// input: left and right is value
	//     *
	//    / \
	// left right
	if len(n.left.value) != 0 && len(n.right.value) != 0 {
		return n
	}
	// Now, some of or all (left , right) is not value

	// input: left and right is not value
	if len(n.left.value) == 0 && len(n.right.value) == 0 {
		// left or/and right is (+,-,*)
		if n.left.sign == multiply && n.right.sign == multiply {
			//     *
			//    / \
			//   *   *
			// left right
			l := n.left.simplification()
			r := n.right.simplification()
			if l.sign == multiply && r.sign == multiply {
				return op(l, r, multiply)
			}
			return op(l, r, multiply).simplification()
		}
		if n.left.sign != multiply && n.right.sign != multiply {
			//      *
			//     / \
			//   +,- +,-
			//  / \  / \
			// a   b c  d
			// (a +/- b) * (c +/- d)
			// ac +/- ad +/- bc +/- bd
			// =========
			//     V1
			//    =============
			//         V2
			//        ================
			//                 V3
			a := n.left.left
			b := n.left.right
			c := n.right.left
			d := n.right.right
			a.simplification()
			b.simplification()
			c.simplification()
			d.simplification()
			ac := op(a,c,multiply)
			ad := op(a,d,multiply)
			bc := op(b,c,multiply)
			bd := op(b,d,multiply)
			return op(
				,
				,

			).simplification()
		}
	}
	/*
		// input: left is value
		if n.left.value != "" && n.right.sign != none {
			//        *
			//      /   \
			//     V1    +
			//          / \
			//         c   d
			// V1 *( c -d )
			// V1 * c  -  V1 *d
			V1 := n.left
			c := n.right.left
			d := n.right.right
			c.simplification()
			d.simplification()
			result := op(op(V1, c, multiply), op(V1, d, multiply), n.right.sign)

			fmt.Println("D1    --> ", result)
			fmt.Println("D1.V1 --> ", c)
			fmt.Println("D1.V2 --> ", d)
			return result.simplification()
		}

		a := n.left.left
		b := n.left.right
		c := n.right.left
		d := n.right.right
		fmt.Println("a =", a)
		fmt.Println("b =", b)
		fmt.Println("c =", c)
		fmt.Println("d =", d)

		if a == nil && b == nil && c == nil && d == nil {
			return n
		}
		if n.left == nil {
			fmt.Println("***********************")
			return n
		}
		if n.right == nil {
			fmt.Println("***********************")
			return n
		}
		if a == nil {
			fmt.Println("***********************")
			fmt.Println("left = ", n.left)
			fmt.Println("ri   = ", n.right)
			fmt.Println("a =", a)
			fmt.Println("b =", b)
			fmt.Println("c =", c)
			fmt.Println("d =", d)
			fmt.Println("***********************")
			return n
		}
		if (n.left.sign == plus || n.left.sign == minus) && (n.right.sign == plus || n.right.sign == minus) {
			switch {
			case n.left.sign == plus && n.right.sign == plus:
				// (a + b) * (c + d)
				// (a*c + a*d) + (b*c + b*d)
				//      V1     +      V2
				V1 := op(op(a, c, multiply), op(a, d, multiply), plus)
				V2 := op(op(b, c, multiply), op(b, d, multiply), plus)
				result := op(V1, V2, plus)
				//V1.simplification()
				//V2.simplification()
				fmt.Println("A1    --> ", result)
				fmt.Println("A1.V1 --> ", V1)
				fmt.Println("A1.V2 --> ", V2)
				return result.simplification()
			case n.left.sign == minus && n.right.sign == plus:
				panic("")
			case n.left.sign == plus && n.right.sign == minus:
				panic("")
			case n.left.sign == minus && n.right.sign == minus:
				panic("")
			}
		}
		if n.left.sign == plus || n.left.sign == minus {
			//      *
			//    /   \
			//   -     *
			//  / \    c
			//  a  b
			// (a - b) * c
			// a*c - b*c
			// V1  - V2
			c = n.right
			V1 := op(a, c, multiply)
			V2 := op(b, c, multiply)
			c.simplification()
			return op(V1, V2, n.left.sign)
		}
		if n.right.sign == plus || n.right.sign == minus {
			//          *
			//        /   \
			//       *     +
			//      a     /  \
			//           c    d
			// a * (c-d)
			// a*c - a*d
			// V1  - V2
			a = n.left
			V1 := op(a, c, multiply)
			V2 := op(a, d, multiply)
			a.simplification()
			V1.simplification()
			V2.simplification()
			return op(V1, V2, n.right.sign).simplification()
		}*/
	panic(fmt.Sprintf("ATTENTION : %v", n))
}

func (n *node) haveInternal() bool {
	if n.sign == none {
		return false
	}
	if n.sign == multiply && (n.left.sign == plus || n.left.sign == minus) {
		return true
	}
	if n.sign == multiply && (n.right.sign == plus || n.right.sign == minus) {
		return true
	}
	return n.left.haveInternal() || n.right.haveInternal()
}

func (n *node) String() (s string) {
	if n.sign == none {
		s += fmt.Sprintf("%v", n.value)
	} else {
		if n.sign == multiply && (n.left.sign == plus || n.left.sign == minus) {
			s += fmt.Sprintf("(%v)", n.left)
		} else {
			s += fmt.Sprintf("%v", n.left)
		}
		s += fmt.Sprintf(" %v ", n.sign)
		if n.sign == multiply && (n.right.sign == plus || n.right.sign == minus) {
			s += fmt.Sprintf("(%v)", n.right)
		} else {
			s += fmt.Sprintf("%v", n.right)
		}
	}
	return s
}

func TestNodes(t *testing.T) {
	a := newNode("a")
	b := newNode("b")
	c := newNode("c")
	d := newNode("d")
	e := newNode("e")
	f := newNode("f")
	g := newNode("g")
	h := newNode("h")

	s := op(op(op(a, b, plus), c, multiply), op(op(d, e, plus), op(f, op(g, h, plus), minus), multiply), plus)

	fmt.Println("s      = ", s)
	fmt.Println("simple = ", s.simplification())

	/*
		c11 := (a.multiply(&e)).plus(b.multiply(&h))
		c12 := (a.multiply(&f)).plus(b.multiply(&g))
		c21 := (c.multiply(&e)).plus(d.multiply(&h))
		c22 := (c.multiply(&f)).plus(d.multiply(&g))
		fmt.Println("c11 = ", c11)
		fmt.Println("c12 = ", c12)
		fmt.Println("c21 = ", c21)
		fmt.Println("c22 = ", c22)

		s := (a.plus(&b).plus(&c)).multiply(d.plus(&f).plus(&g))
		fmt.Println("s = ", s)
		fmt.Println("ss =", s.haveInternal())
		fmt.Println("simply1 = ", s.simplification(0))
		//fmt.Println("simply2 = ", s.simplification(0))
		fmt.Printf("\n\n\n\n ---\n")

		z := c21.plus(c22)
		fmt.Println("z = ", z)
		fmt.Println("zz= ", z.haveInternal())

		y := (a.multiply(&b).plus(&c)).multiply(&d)
		fmt.Println("y = ", y)
		fmt.Println("yy= ", y.haveInternal())

		p := a.multiply(&b).multiply(c.plus(&d)).plus(e.multiply(&h))
		fmt.Println("p = ", p)
		fmt.Println("pp =", p.haveInternal())
		fmt.Println("p*= ", p.amountMultiplications())
	*/
}
