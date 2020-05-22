package runtime

import (
	"fmt"
	"strconv"
)

type builtin interface {
	Number() (Number, bool)
	String() (String, bool)
	Bool() (Bool, bool)
}

type Number struct {
	float64
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

func (n *Number) Add(i builtin) {
	if in, b := i.Number(); b {
		n.float64 += in.float64
	}
}

func (n *Number) Subtract(i builtin) {
	if in, b := i.Number(); b {
		n.float64 -= in.float64
	}
}

func (n *Number) Multiply(i builtin) {
	if in, b := i.Number(); b {
		n.float64 *= in.float64
	}
}

func (n *Number) Divide(i builtin) {
	if in, b := i.Number(); b {
		n.float64 /= in.float64
	}
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

func Print(i interface{}) {
	if bi, b := i.(builtin); b {
		str, t := bi.String()
		if t {
			fmt.Println(str.string)
			return
		}
	}

	fmt.Println(i)
}
