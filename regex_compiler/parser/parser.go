package parser

import()

func Parse(regex string) (bool, *Start) {
	start := &Start{}
	count := 0
	success := start.Parse(regex, &count)
	return success, start
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
	default:
		quant.Type = Epsilon
	}
	return true
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
		ident.Char = regex[*count]
		(*count)++
		return true
	} else {
		return false //invalid character
	}
}

