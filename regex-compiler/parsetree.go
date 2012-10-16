package parsetree

import()

const(
	Pipe = iota
	Star
	Plus
	Ques
	LParen
	RParen
)

const(
	Start = iota
	Juxt
	Quant
	Ident
	End
)

type Start struct {
	Left *Juxt
	Right *Start_
	Empty bool
}

func (start *Start) Parse(regex string, count *int) bool {
	if *count >= len(regex) {
		start.Empty = true
		return true
	}
	start.Left = &Juxt{}
	if !start.Left.Parse(regex, count) { //did not consume anything
		return false
	}
	start.Right = &Start_{}
	return start.Right.Parse(regex, count)
}

type Start_ struct {
	Left *Juxt
	Right *Start_
	Empty bool
}

func (start *Start_) Parse(regex string, count *int) bool {
	if *count >= len(regex) {
		start.Empty = true
		return true
	}
	if regex[count] == Pipe {
		(*count)++ //consume input
		start.Left = &Juxt{}
		if !start.Left.Parse(regex, count) {
			return false
		}
		start.Right = &Start_{}
		return start.Right.Parse(regex, count)
	} else {
		start.Empty = true
		return true
	}
}

type Juxt struct {
	Left *Quant
	Right *Juxt_
	Empty bool
}

func (juxt *Juxt) Parse(regex string, count *int) bool {
	if count >= len(regex) {
		juxt.Empty = true
		return true
	}
	juxt.Left = &Quant{}
	if !juxt.Left.Parse(regex, count) {
		return false
	}
	juxt.Right = &Juxt_{}
	return juxt.Right.Parse(regex, count)
}

type Juxt_ struct {
	Left *Quant
	Right *Juxt_
	Empty bool
}

func (juxt *Juxt_) Parse(regex string, count *int) bool {
	if count >= len(regex) {
		juxt.Empty = true
		return true
	}
	if regex[count] < Pipe { //that means its a regular character
		juxt.Left = &Quant{}
		if !juxt.Left.Parse(regex, count) {
			return false
		}
		juxt.Right = &Juxt_{}
		return juxt.Right.Parse(regex, count)
	} else {
		juxt.Empty = true
		return true
	}
}

type Quant struct {
	Left *Ident
	Type int
	Empty bool
}

func (quant *Quant) Parse(regex string, count *int) bool {
	if count >= len(regex) {
		quant.Empty = true
		return true
	}
	quant.Left = &Ident{}
	if !quant.Left.Parse(regex, count) {
		return false
	}
	switch(regex[*count]) {
	case Star:
		(*count)++
		quant.Type = Star
		break
	case Plus:
		(*count)++
		quant.Type = Plus
		break
	case Ques:
		(*count)++
		quant.Type = Ques
		break
	}
	return true
}

type Ident struct {
	Left *Start
	Char rune
	Empty bool
}

func (ident *Ident) Parse(regex string, count *int) bool {
	if count >= len(regex) {
		ident.Empty = true
		return true
	}
	if regex[*count] == LParen {
		(*count)++
		ident.Left = &Start{}
		if !ident.Left.Parse(regex, count) {
			return false
		}
		if regex[*count] == RParen {
			(*count)++
			return true
		} else {
			return false
		}
	} else if regex[*count] < Pipe {
		ident.Char = regex[(*count)++]
		return true
	} else {
		return false //invalid character
	}
}

