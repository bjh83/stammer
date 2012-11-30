package lexer

import(
	"../regex"
)

const(
	NULL = iota
	INT_V = iota
	FLOAT_V = iota
	STRING_V = iota
	VAR_V = iota
	INT = iota
	FLOAT = iota
	STRING = iota
	VAR = iota
	CLASS = iota
	IF = iota
	FOR = iota
	FUNC = iota
	EQ = iota
	EQEQ = iota
	NOT = iota
	NOTEQ = iota
)

var regex0 regex.Regex
var regex1 regex.Regex
var regex2 regex.Regex
var regex3 regex.Regex
var regex4 regex.Regex
var regex5 regex.Regex
var regex6 regex.Regex
var regex7 regex.Regex
var regex8 regex.Regex
var regex9 regex.Regex
var regex10 regex.Regex
var regex11 regex.Regex
var regex12 regex.Regex
var regex13 regex.Regex
var regex14 regex.Regex
var regex15 regex.Regex
var funcArray []func(string)(int, int) = []func(string)(int, int) {func0, func0, func1, func2, func3, func4, func5, func6, func7, func8, func9, func10, func11, func12, func13, func14, func15}

func setUp() {

	regex.Declare("letter", "[a-zA-Z]")
	regex.Declare("name_char", "({letter}|_)")
	regex.Declare("digit", "[0-9]")
	regex.Declare("white_space", "[\n\t ]")
	regex.Declare("any", "[ -\\-]")
	regex0 = regex.Compile("{digit}+")
	regex1 = regex.Compile("{digit}+.{digit}*|{digit}*.{digit}+")
	regex2 = regex.Compile("{any}+")
	regex3 = regex.Compile("{name_char}({digit}|{name_char})*")
	regex4 = regex.Compile("int")
	regex5 = regex.Compile("float")
	regex6 = regex.Compile("string")
	regex7 = regex.Compile("var")
	regex8 = regex.Compile("class")
	regex9 = regex.Compile("if")
	regex10 = regex.Compile("for")
	regex11 = regex.Compile("func")
	regex12 = regex.Compile("=")
	regex13 = regex.Compile("==")
	regex14 = regex.Compile("!")
	regex15 = regex.Compile("!=")
}

func func0(input string) (type, valOrId) {
	if regex0.match(input) {
		return INT_V, VAL
	}
	return NULL, -1
}

func func1(input string) (type, valOrId) {
	if regex1.match(input) {
		return FLOAT_V, VAL
	}
	return NULL, -1
}

func func2(input string) (type, valOrId) {
	if regex2.match(input) {
		return STRING_V, VAL
	}
	return NULL, -1
}

func func3(input string) (type, valOrId) {
	if regex3.match(input) {
		return VAR_V, ID
	}
	return NULL, -1
}

func func4(input string) (type, valOrId) {
	if regex4.match(input) {
		return INT, NIL
	}
	return NULL, -1
}

func func5(input string) (type, valOrId) {
	if regex5.match(input) {
		return FLOAT, NIL
	}
	return NULL, -1
}

func func6(input string) (type, valOrId) {
	if regex6.match(input) {
		return STRING, NIL
	}
	return NULL, -1
}

func func7(input string) (type, valOrId) {
	if regex7.match(input) {
		return VAR, NIL
	}
	return NULL, -1
}

func func8(input string) (type, valOrId) {
	if regex8.match(input) {
		return CLASS, NIL
	}
	return NULL, -1
}

func func9(input string) (type, valOrId) {
	if regex9.match(input) {
		return IF, NIL
	}
	return NULL, -1
}

func func10(input string) (type, valOrId) {
	if regex10.match(input) {
		return FOR, NIL
	}
	return NULL, -1
}

func func11(input string) (type, valOrId) {
	if regex11.match(input) {
		return FUNC, NIL
	}
	return NULL, -1
}

func func12(input string) (type, valOrId) {
	if regex12.match(input) {
		return EQ, NIL
	}
	return NULL, -1
}

func func13(input string) (type, valOrId) {
	if regex13.match(input) {
		return EQEQ, NIL
	}
	return NULL, -1
}

func func14(input string) (type, valOrId) {
	if regex14.match(input) {
		return NOT, NIL
	}
	return NULL, -1
}

func func15(input string) (type, valOrId) {
	if regex15.match(input) {
		return NOTEQ, NIL
	}
	return NULL, -1
}

