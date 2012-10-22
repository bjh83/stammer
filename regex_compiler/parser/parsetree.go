package parser

import()

const(Buffer = 255)

const(
	Pipe = iota + Buffer
	Star = iota + Buffer
	Plus = iota + Buffer
	Ques = iota + Buffer
	Epsilon = iota + Buffer
	LParen = iota + Buffer
	RParen = iota + Buffer
)

type Start struct {
	Left *Juxt
	Right *Start_
	Empty bool
}

type Start_ struct {
	Left *Juxt
	Right *Start_
	Empty bool
}

type Juxt struct {
	Left *Quant
	Right *Juxt_
	Empty bool
}

type Juxt_ struct {
	Left *Quant
	Right *Juxt_
	Empty bool
}

type Quant struct {
	Left *Ident
	Type int
	Empty bool
}

type Ident struct {
	Left *Start
	Char uint8
	Empty bool
}

