package parser

import()

const(
	Pipe = iota
	Star
	Plus
	Ques
	Epsilon
	LParen
	RParen
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
	Char rune
	Empty bool
}

