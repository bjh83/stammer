//parsetree.go

type Type int

const(
	Star = iota
	Plus
	Ques
	Pipe
	Char
)

type Pipe struct {
	Left Juxta
	Right Pipe_
}

func (p Pipe) Evaluate() list *OpList {
	list = p.Left.Evaluate()
	if p.Right != nil {
		list.Append(p.Right.Evaluate())
	}
}

type Pipe_ struct {
	Left Juxta
	Right Pipe_
}

func (p Pipe_) Evaluate() list *OpList {

type Juxta struct {
	Left Quanta
	Right Juxta_
}

type Juxta_ struct {
	Left Quanta
	Right Juxta_
}

type Quanta struct {
	Left, Right Prim
}

type Prim struct {
	Val Evaluater
	Op Type
}

type Evaluater interface {
	Evaluate() OpList
}

type Char struct {
	Val rune
}

