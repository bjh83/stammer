package lexer

import(
	"../regex"
	"fmt"
)

func setup() {
	var any string
	for var i uint8 = 0; i < 128; i++ {
		any += string(i)
	}
	regex.Declare("letter", "[a-zA-Z]")
	regex.Declare("name_char", "({letter}|_)")
	regex.Declare("digit", "[0-9]")
	regex.Declare("white_space", "[\n\t ]")
	regex.Declare("any", "[" + any + "]")
	reg_intConst = regex.Compile("{digit}+")
	reg_floatConst = regex.Compile("{digit}+.{digit}*|{digit}*.{digit}+")
	reg_stringConst = regex.Compile("{any}+")
	reg_variableName = regex.Compile("{name_char}({digit}|{name_char})*")
	reg_int = regex.Compile("int")
	reg_float = regex.Compile("float")
	reg_string = regex.Compile("string")
	reg_var = regex.Compile("var")
	reg_class = regex.Compile("class")
	reg_if = regex.Compile("if")
	reg_for = regex.Compile("for")
	reg_func = regex.Compile("func")
	reg_EQ = regex.Compile("=")
	reg_EQEQ = regex.Compile("==")
	reg_NOT = regex.Compile("!")
	reg_NOTEQ = regex.Compile("!=")
