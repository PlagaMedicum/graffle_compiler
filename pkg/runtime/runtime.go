package runtime

import (
	"fmt"
	"log"
	"strconv"
)

type Builtin interface {
	Number() (Number, bool)
	String() (String, bool)
	Bool() (Bool, bool)
}

type Number struct {
	float64
}

func NewNumber(val float64) Number {
	return Number{val}
}

func NewString(val string) String {
	return String{val}
}

func NewBool(val bool) Bool {
	return Bool{val}
}

func (n Number) Number() (Number, bool) {
	return n, true
}

func (n Number) String() (String, bool) {
	return String{fmt.Sprintf("%f", n.float64)}, true
}

func (n Number) Bool() (Bool, bool) {
	if n.float64 == 0 {
		return Bool{false}, true
	}
	return Bool{true}, true
}

type String struct {
	string
}

func (s String) Number() (Number, bool) {
	f, err := strconv.ParseFloat(s.string, 64)
	if err != nil {
		return Number{0}, false
	}
	return Number{f}, true
}

func (s String) String() (String, bool) {
	return s, true
}

func (s String) Bool() (Bool, bool) {
	if s.string == "" {
		return Bool{false}, true
	}
	return Bool{true}, true
}

type Bool struct {
	bool
}

func (b Bool) Number() (Number, bool) {
	if b.bool {
		return Number{1}, true
	}
	return Number{0}, false
}

func (b Bool) String() (String, bool) {
	if b.bool {
		return String{"True"}, true
	}
	return String{"False"}, false
}

func (b Bool) Bool() (Bool, bool) {
	return b, true
}

func Add(l, r Builtin) Builtin {
	ln, lb := l.Number()
	rn, rb := r.Number()
	if lb && rb {
		ln.float64 += rn.float64
		return ln
	}
	log.Fatalf("Error! Wrong types passed in Addition! left: %t, right: %t", l, r)
	return Number{}
}

func Subtract(l, r Builtin) Builtin {
	ln, lb := l.Number()
	rn, rb := r.Number()
	if lb && rb {
		ln.float64 -= rn.float64
		return ln
	}
	log.Fatalf("Error! Wrong types passed in Subtraction! left: %t, right: %t", l, r)
	return Number{}
}

func Multiply(r, l Builtin) Builtin {
	ln, lb := l.Number()
	rn, rb := r.Number()
	if lb && rb {
		ln.float64 *= rn.float64
		return ln
	}
	log.Fatalf("Error! Wrong types passed in Multiplication! left: %t, right: %t", l, r)
	return Number{}
}

func Divide(r, l Builtin) Builtin {
	ln, lb := l.Number()
	rn, rb := r.Number()
	if lb && rb {
		ln.float64 /= rn.float64
		return ln
	}
	log.Fatalf("Error! Wrong types passed in Dividing! left: %t, right: %t", l, r)
	return Number{}
}

func Print(i interface{}) {
	if bi, b := i.(Builtin); b {
		str, t := bi.String()
		if t {
			fmt.Println(str.string)
			return
		}
	}

	fmt.Println(i)
}
