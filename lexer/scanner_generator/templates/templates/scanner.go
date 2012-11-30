package scanner

import(
	"stammer/regex"
)

{% for var in name_vars -%}
var {{var}}_regex regex.Regex
{%- endfor %}

var regex_functions []func(string)(Token) = []func(string)(Token) {{"{"}}{{name_vars|join("_func, ")}}{{"_func }"}}

func startup() {
	{% for var in regex_declaration_vars -%}
	regex.Declare({{var[0]}}, {{var[1]}})
	{%- endfor %}
	{% for var in regex_compile_vars -%}
	{{var[0]}}_regex = regex.Compile({{var[1]}})
	{%- endfor %}
}

{% for var in function_vars -%}
func {{var[0]}}_func(input string) Token {
	if {{var[0]}}.Match(input) {
		{{var[1]}}
		return Token{{"{"}}{{var[0]}}, {{"output}"}}
	}
	return Token{{"{NULL}"}}
}
