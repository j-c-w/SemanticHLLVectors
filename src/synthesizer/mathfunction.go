package synthesizer

import "strconv"

var binaryOperators = []string {"+", "/", "-", "*"};
var unaryOperators = []string {"thresh"};

type Node interface {
	ToString() string
	Evaluate(int) int
	IsSimplifiable() bool
	IsConst() bool
}

type BinaryNode struct {
	op string
	child1 Node
	child2 Node
}

type UnaryNode struct {
	op string
	child Node
}

type IntChild struct {
	value int
}

type VariableChild struct {
	name string
}

func (c IntChild) Evaluate(x int) int {
	return c.value
}

func (b BinaryNode) Evaluate(x int) int {
	switch b.op {
	case "+": return b.child1.Evaluate(x) + b.child2.Evaluate(x)
	case "*": return b.child1.Evaluate(x) * b.child2.Evaluate(x)
	case "/":  {
		c2Value := b.child2.Evaluate(x)
		if c2Value != 0 {
			return b.child1.Evaluate(x) / c2Value
		} else {
			// Uhm. Clearly a bit of a hack.
			return b.child1.Evaluate(x)
		}
	}
	case "-": return b.child1.Evaluate(x) - b.child2.Evaluate(x)
	default: panic("Unknwon operator")
	}
}

func (c IntChild) IsSimplifiable() bool {
	return false
}

func (v VariableChild) IsSimplifiable() bool {
	return false
}

func (u UnaryNode) IsSimplifiable() bool {
	return u.child.IsConst()
}

func (b BinaryNode) IsSimplifiable() bool {
	return b.child1.IsConst() && b.child2.IsConst()
}

func (c IntChild) IsConst() bool {
	return true
}

func (v VariableChild) IsConst() bool {
	return false
}

func (u UnaryNode) IsConst() bool {
	return true
}

func (b BinaryNode) IsConst() bool {
	return false
}

func (u UnaryNode) Evaluate(x int) int {
	switch u.op {
	case "thresh": {
		child := u.child.Evaluate(x)
		if child > 0 {
			return child
		} else {
			return 0
		}
	}
	default: panic("Unknown operator")
	}
}

func (v VariableChild) Evaluate(x int) int {
	return x
}

func (b BinaryNode) ToString() string {
	return "(" + b.child1.ToString() + b.op + b.child2.ToString() + ")"
}

func (n IntChild) ToString() string {
	return strconv.Itoa(n.value)
}

func (v VariableChild) ToString() string {
	return v.name
}

func (u UnaryNode) ToString() string {
	return "(" + u.op + " " + u.child.ToString() + ")"
}
