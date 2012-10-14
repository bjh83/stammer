//parsetree.go

type Type int

const(
	Star = iota
	Plus
	Ques
	Pipe
	Char
)

type ParseTree struct {
	Type ttype
	left, right ParseTree
	char rune
}

func (p ParseTree) AddLeft(ttype Type) {
	toAdd := ParseTree{ttype: ttype}
	p.left = toAdd
}

