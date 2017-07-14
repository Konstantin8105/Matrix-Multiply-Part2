package main

/*
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

/*
type Sign int

func (sign Sign) String() (s string) {
	switch sign {
	case plus:
		s += fmt.Sprintf(" + ")
	case minus:
		s += fmt.Sprintf(" - ")
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
	// value parameters
	value string
	//valueSign Sign
	// node parameters
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
	//n.valueSign = plus
	n.value = v
	return
}

func op(n1, n2 *node, sign Sign) (result *node) {
	result = new(node)
	switch sign {
	case plus:
		(*result).sign = plus
	case minus:
		(*result).sign = minus
	case multiply:
		(*result).sign = multiply
	}
	(*result).left = n1
	(*result).right = n2
	return
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

func (n *node) compressValues() (result *node) {
	if n == nil {
		return nil
	}
	switch n.sign {
	// input: value
	case none:
		panic("Error")
	case plus:
		//fmt.Printf(">+ %v\n", n)
		result = op(n.left.compressValues(), n.right.compressValues(), plus)
	case minus:
		//fmt.Printf(">- %v\n", n)
		result = op(n.left.compressValues(), n.right.compressValues(), minus)
	case multiply:
		//fmt.Printf(">* %v\n", n)
		var name string
		if ([]byte(n.left.value))[0] < ([]byte(n.right.value))[0] {
			name = fmt.Sprintf("%v%v", n.left.value, n.right.value)
		} else {
			name = fmt.Sprintf("%v%v", n.right.value, n.left.value)
		}
		result = newNode(name)
	}
	return
}

func (n *node) compressEquation() (result *node) {
	defer func() {
		// check again
	}()
	if n == nil {
		return nil
	}
	// 2 values
	if len(n.left.value) != 0 && len(n.right.value) != 0 {
		var name string
		if ([]byte(n.left.value))[0] < ([]byte(n.right.value))[0] {
			name = fmt.Sprintf("%v%v%v", n.left.value, n.sign, n.right.value)
		} else {
			name = fmt.Sprintf("%v%v%v", n.right.value, n.sign, n.left.value)
		}
		result = newNode(name)
	}
	switch n.sign {
	// input: value
	case none:
		panic("Error")
	case plus:
		a := n.left.compressEquation()
		b := n.right.compressEquation()
		result = op(a, b, plus)
	//panic("")
	case minus:
		if n.right.sign == minus {
			n.sign = plus
			if n.right.left != nil && n.right.right != nil {
				n.right.left, n.right.right = n.right.right, n.right.left
			}
		}
		a := n.left.compressEquation()
		b := n.right.compressEquation()
		result = op(a, b, plus)
		//panic("")
	case multiply:
		a := n.left.compressEquation()
		b := n.right.compressEquation()
		result = op(a, b, multiply)
		//panic("")
	}
	return
}

func (n *node) simplification() (result *node) {
	defer func() {
		if result.haveInternal() {
			result = result.simplification()
		}
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
		v := op(a, b, n.sign)
		return v
	}
	// Now, root node is multiplication

	// input: left and right is value
	//     *
	//    / \
	// left right
	if len(n.left.value) != 0 && len(n.right.value) != 0 {
		return op(n.left, n.right, multiply)
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
			return op(n.left, n.right, multiply)
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
			ac := op(a, c, multiply)
			ad := op(a, d, multiply)
			bc := op(b, c, multiply)
			bd := op(b, d, multiply)
			V1 := op(ac, ad, n.right.sign)
			V2 := op(V1, bc, n.left.sign)
			var V3 *node
			if n.left.sign == n.right.sign {
				// - * - = +
				// + * + = +
				V3 = op(V2, bd, plus)
			} else {
				// - * + = -
				// + * - = -
				V3 = op(V2, bd, minus)
			}
			return V3
		}
	}
	// Now, left or right is value

	// left  = value
	// right = +,-,*
	if len(n.left.value) != 0 && len(n.right.value) == 0 {
		if n.right.sign == multiply {
			//       *
			//      / \
			//     v   *
			//        / \
			//       a   b
			//    v *a * b
			//   a * b * v
			//   =====
			//     V1
			//     =======
			//        V2
			a := n.right.left
			b := n.right.right
			v := n.left
			a.simplification()
			b.simplification()
			V1 := op(a, b, multiply)
			V2 := op(V1, v, multiply)
			return V2
		}
		//       *
		//     /   \
		//    v    +,-
		//        /   \
		//       a     b
		//  v*a +/- v,b
		//  V1       V2
		v := n.left
		a := n.right.left
		b := n.right.right
		V1 := op(v, a, multiply)
		V2 := op(v, b, multiply)
		a.simplification()
		b.simplification()
		return op(V1, V2, n.right.sign)
	}

	// left  = +,-,*
	// right = value
	if len(n.left.value) == 0 && len(n.right.value) != 0 {
		if n.left.sign == multiply {
			//       *
			//      / \
			//      *  v
			//    /  \
			//   a    b
			//   b * v * a
			//   =====
			//     V1
			//     =======
			//        V2
			v := n.right
			a := n.left.left
			b := n.left.right
			a.simplification()
			b.simplification()
			V1 := op(b, v, multiply)
			V2 := op(V1, a, multiply)
			return V2
		}
		//       *
		//      / \
		//    +,-  v
		//    /  \
		//   a    b
		//  v*a +/- v,b
		//  V1       V2
		v := n.right
		a := n.left.left
		b := n.left.right
		V1 := op(v, a, multiply)
		V2 := op(v, b, multiply)
		a.simplification()
		b.simplification()
		return op(V1, V2, n.left.sign)
	}
	fmt.Printf("n     := %#v\n", n)
	fmt.Printf("left  := %#v\n", n.left)
	fmt.Printf("right := %#v\n", n.right)

	fmt.Printf("n     := %v\n", n)
	fmt.Printf("left  := %v\n", n.left)
	fmt.Printf("sign  := %v\n", n.sign)
	fmt.Printf("right := %v\n", n.right)
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
		s += fmt.Sprintf("(%v%v%v)", n.left, n.sign, n.right)
	}
	return s
}

func TestFormula1(t *testing.T) {
	a := newNode("a")
	b := newNode("b")
	c := newNode("c")

	p1 := op(a, op(b, c, minus), multiply)
	p1 = p1.simplification()
	if fmt.Sprintf("%v", p1) != "((a*b) - (a*c))" {
		t.Errorf("Not correct")
	}
}

func TestFormula2(t *testing.T) {
	a := newNode("a")
	b := newNode("b")
	c := newNode("c")
	d := newNode("d")
	e := newNode("e")

	p1 := op(op(b, c, minus), op(a, op(d, e, minus), plus), multiply)
	p1 = p1.simplification()
	if fmt.Sprintf("%v", p1) != "((((b*a) + ((b*d) - (b*e))) - (c*a)) - ((c*d) - (c*e)))" {
		t.Errorf("Not correct")
	}

}

func TestNodes(t *testing.T) {
	a := newNode("a")
	b := newNode("b")
	c := newNode("c")
	d := newNode("d")
	e := newNode("e")

	w1 := op(op(b, c, minus), op(a, op(d, e, minus), plus), multiply)

	fmt.Printf("I.    = %v\n", w1)
	w2 := w1.simplification()
	fmt.Printf("II.   = %v\n", w1)
	fmt.Printf("III.  = %v\n", w2)
	fmt.Printf("IV.   >>>%v\n", w2.haveInternal())
	w3 := w2.compressValues()
	fmt.Printf("V.    = %v\n", w2)
	fmt.Printf("VI.   = %v\n", w3)
	w4 := w3.compressEquation()
	fmt.Printf("VII.  = %v\n", w3)
	fmt.Printf("VIII. = %v\n", w4)
}
*/
