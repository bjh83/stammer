package parser

import()

func (start *Start) Generate() *OpList {
	oplist := New()
	oplist.Start()
	oplist.Append(start.generate())
	oplist.Finish()
	return oplist
}

func (start *Start) generate() *OpList {
	oplist := start.Left.generate()
	if !start.Right.Empty {
		//this means there should be a split here
		save := oplist.AddSplit(oplist.Length - 1, -1) //we do not know the second address yet
		oplist.Append(start.Right.generate())
		save.Line2 = oplist.Length - 1
	}
	return oplist
}

func (start *Start_) generate() *OpList {
	oplist := start.Left.generate()
	if !start.Right.Empty {
		save := oplist.AddSplit(oplist.Length - 1, -1)
		oplist.Append(start.Right.generate())
		save.Line2 = oplist.Length - 1
	}
	return oplist
}

func (juxt *Juxt) generate() *OpList {
	oplist := juxt.Left.generate()
	if !juxt.Right.Empty {
		oplist.Append(juxt.Right.generate())
	}
	return oplist
}

func (juxt *Juxt_) generate() *OpList {
	oplist := juxt.Left.generate()
	if !juxt.Right.Empty {
		oplist.Append(juxt.Right.generate())
	}
	return oplist
}

func (quant *Quant) generate() *OpList {
	oplist := New()
	switch quant.Type {
	case Star:
		save := oplist.AddSplit(1, -1)
		oplist.Append(quant.Left.generate())
		oplist.AddJump(0)
		save.Line2 = oplist.Length
		break
	case Plus:
		oplist.Append(quant.Left.generate())
		splitAddress := oplist.Length - 1
		save := oplist.AddSplit(splitAddress + 1, -1)
		oplist.Append(quant.Left.generate())
		oplist.AddJump(splitAddress)
		save.Line2 = oplist.Length
		break
	case Ques:
		save := oplist.AddSplit(1, -1)
		oplist.Append(quant.Left.generate())
		save.Line2 = oplist.Length
		break
	case Epsilon:
		oplist = quant.Left.generate()
		break
	}
	return oplist
}

func (ident *Ident) generate() *OpList {
	if ident.Left != nil {
		return ident.Left.generate()
	}
	oplist := New()
	oplist.AddChar(ident.Char)
	return oplist
}

