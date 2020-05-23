package runtime

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type BuiltinType interface {
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
	return NewString(fmt.Sprintf("%f", n.float64)), true
}

func (n Number) Bool() (Bool, bool) {
	if n.float64 == 0 {
		return NewBool(false), true
	}
	return NewBool(true), true
}

type String struct {
	string
}

func (s String) Number() (Number, bool) {
	str := strings.ReplaceAll(s.string, ",", ".")
	str = strings.ReplaceAll(str, " ", "")
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return NewNumber(0), false
	}
	return NewNumber(f), true
}

func (s String) String() (String, bool) {
	return s, true
}

func (s String) Bool() (Bool, bool) {
	str := strings.ToLower(s.string)
	str = strings.ReplaceAll(str, " ", "")
	if str == "" || str == "false" || str == "nottrue" || str == "0" {
		return NewBool(false), true
	}
	return NewBool(true), true
}

type Bool struct {
	bool
}

func (b Bool) Number() (Number, bool) {
	if b.bool {
		return NewNumber(1), true
	}
	return NewNumber(0), true
}

func (b Bool) String() (String, bool) {
	if b.bool {
		return NewString("True"), true
	}
	return NewString("False"), true
}

func (b Bool) Bool() (Bool, bool) {
	return b, true
}

type Stringer interface {
	String() (String, bool)
}

type Vertice struct {
	Label string
	Val   BuiltinType
}

func (v *Vertice) String() (String, bool) {
	str, _ := v.Val.String()
	if v.Label == "" {
		return NewString(fmt.Sprintf("(%s)", str)), true
	}
	return NewString(fmt.Sprintf("(%s)@[%s]", str, v.Label)), true
}

type Edge struct {
	Label string
	Weight Number
	IsDirected bool
	V1 Vertice
	V2 Vertice
}

func (e * Edge) String() (String, bool) {
	v1, _ := e.V1.String()
	str := fmt.Sprintf("%s -", v1)

	if e.Weight.float64 != 0 {
		w, _ := e.Weight.String()
		str = fmt.Sprintf("%s[%s]-", str, w)
	}

	if e.IsDirected {
		str += ">"
	}

	v2, _ := e.V2.String()
	str = fmt.Sprintf("%s %s", str, v2)

	if e.Label != "" {
		str = fmt.Sprintf("{%s}@[%s]", str, e.Label)
	}

	return NewString(str), true
}

type Graph struct {
	Label string
	V []Vertice
	E []Edge
}

func (g *Graph) GetSingleVertices() []Vertice {
	connv := map[Vertice]struct{}{}
	for _, e := range g.E {
		connv[e.V1] = struct{}{}
		connv[e.V2] = struct{}{}
	}

	var sv []Vertice
	for _, v := range g.V {
		if _, b := connv[v]; !b {
			sv = append(sv, v)
		}
	}
	return sv
}

func (g *Graph) String() (String, bool) {
	var str string
	for _, v := range g.GetSingleVertices() {
		vstr, _ := v.String()
		str = fmt.Sprintf("%s\n\t%s;", str, vstr)
	}

	for _, e := range g.E {
		estr, _ := e.String()
		str = fmt.Sprintf("%s\n\t%s;", str, estr)
	}

	if g.Label != "" {
		str = fmt.Sprintf("{%s\n}@[%s]", str, g.Label)
	}

	return NewString(str), true
}

func Add(l, r BuiltinType) BuiltinType {
	ln, lb := l.Number()
	rn, rb := r.Number()
	if lb && rb {
		ln.float64 += rn.float64
		return ln
	}
	log.Fatalf("Error! Wrong types passed in Addition! left: %t, right: %t", l, r)
	return NewNumber(0)
}

func Subtract(l, r BuiltinType) BuiltinType {
	ln, lb := l.Number()
	rn, rb := r.Number()
	if lb && rb {
		ln.float64 -= rn.float64
		return ln
	}
	log.Fatalf("Error! Wrong types passed in Subtraction! left: %t, right: %t", l, r)
	return NewNumber(0)
}

func Multiply(r, l BuiltinType) BuiltinType {
	ln, lb := l.Number()
	rn, rb := r.Number()
	if lb && rb {
		ln.float64 *= rn.float64
		return ln
	}
	log.Fatalf("Error! Wrong types passed in Multiplication! left: %t, right: %t", l, r)
	return NewNumber(0)
}

func Divide(r, l BuiltinType) BuiltinType {
	ln, lb := l.Number()
	rn, rb := r.Number()
	if lb && rb {
		ln.float64 /= rn.float64
		return ln
	}
	log.Fatalf("Error! Wrong types passed in Dividing! left: %t, right: %t", l, r)
	return NewNumber(0)
}

func Print(i interface{}) {
	if bi, b := i.(Stringer); b {
		str, t := bi.String()
		if t {
			fmt.Println(str.string)
			return
		}
	}

	fmt.Println(i)
}
