package opcode

import()

const(
	Start = iota
	Char
	Jump
	Split
	Match
)

type OpCode int

type Instruct struct {
	Opperation Opcode
	Line1, Line2 int
}

type OpList struct {
	Head, Tail *OpNode
	Length int
}

type OpNode struct {
	Instruction Instruct
	Next *OpNode
}

func New() *OpList {
	return OpList{}
}

func (list *OpList) add(opperation OpCode, line1, line2 int) *Instruct {
	node := OpNode{Instruction: Instruct{opperation, line1, line2}}
	if list.Length == 0 {
		list.Head = &node
	} else {
		list.Tail.Next = &node
	}
	list.Tail = &node
	list.Length++
	return &node.Instruction //return a reference so that if we do not know an 
	// address we can figure it out later
}

func (list *OpList) AddSplit(line1, line2 int) *Instruct {
	return list.add(Split, line1, line2)
}

func (list *OpList) AddJump(line int) *Instruct {
	return list.add(Jump, line, -1)
}

func (list *OpList) AddChar(char rune) {
	//No reason for this to return a reference
	list.add(Char, int(char), -1)
}

func (list *OpList) Start() {
	list.add(Start, -1, -1)
}

func (list *OpList) Finish() {
	list.add(Match, -1, -1)
}

func (list *OpList) Append(toAppend *OpList) {
	toAdd := list.Length
	if toAdd == 0 {
		list.Head = toAppend.Head
		list.Tail = toAppend.Tail
		list.Length = toAppend.Length
		return
	}
	for node := toAppend.Head; node != nil; node = node.Next {
		if node.Line1 != -1 {
			node.Line1 += toAdd
		}
		if node.Line2 != -1 {
			node.Line2 += toAdd
		}
	}
	list.Length += toAppend.Length
	list.Tail.Next = toAppend.Head
	list.Tail = toAppend.Tail
}

