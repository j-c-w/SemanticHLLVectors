package synthesizer

import "strconv"

var operators = []string {"+", "/", "-", "*"};

type Node interface {
	ToString() string
	Evaluate(int) int
}

type BinaryNode struct {
	op string
	child1 Node
	child2 Node
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
