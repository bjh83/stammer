package parser

import()

func (start *Start) Generate() *OpList {
	oplist := start.Left.Generate()
	if !start.Right.Empty {
		//this means there should be a split here
		save := oplist.AddSplit(oplist.Length - 1, -1) //we do not know the second address yet
		oplist.Append(start.Right.Generate())
		save.Line2 = oplist.Length - 1
	}
	return oplist
}

func (start *Start_) Generate() *OpList {
	oplist := start.Left.Generate()
	if !start.Right.Empty {
		save := oplist.AddSplit(oplist.Length - 1, -1)
		oplist.Append(start.Right.Generate())
		save.Line2 = oplist.Length - 1
	}
	return oplist
}

func (juxt *Juxt) Generate() *OpList {
	oplist := juxt.Left.Generate()
	if !juxt.Right.Empty {
		oplist.Append(juxt.Right.Generate())
	}
	return oplist
}

func (juxt *Juxt_) Generate() *OpList {
	oplist := juxt.Left.Generate()
	if !juxt.Right.Empty {
		oplist.Append(juxt.Right.Generate())
	}
	return oplist
}

func (quant *Quant) Generate() *OpList {
	oplist := New()
	switch quant.Type {
	case Star:
		save := oplist.AddSplit(1, -1)
		oplist.Append(quant.Left.Generate())
		oplist.AddJump(0)
		save.Line2 = oplist.Length
		break
	case Plus:
		oplist.Append(quant.Left.Generate())
		splitAddress := oplist.Length - 1
		save := oplist.AddSplit(splitAddress + 1, -1)
		oplist.Append(quant.Left.Generate())
		oplist.AddJump(splitAddress)
		save.Line2 = oplist.Length
		break
	case Ques:
		save := oplist.AddSplit(1, -1)
		oplist.Append(quant.Left.Generate())
		save.Line2 = oplist.Length
		break
	case Epsilon:
		oplist = quant.Left.Generate()
		break
	}
	return oplist
}

func (ident *Ident) Generate() *OpList {
	if ident.Left != nil {
		return ident.Left.Generate()
	}
	oplist := New()
	oplist.AddChar(ident.Char)
	return oplist
}

