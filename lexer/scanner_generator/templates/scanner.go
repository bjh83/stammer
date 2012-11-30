package scanner

import(
	"stammer/regex"
)

{% for var in name_vars -%}
var {{var[0]}}_regex regex.Regex
{% endfor -%}

var regex_functions []func(string)(Token) = []func(string)(Token){{"{"}}
	{% for var in name_vars -%}
	{{var[0]}}_func,
	{% endfor -%}
}

func startup() {
	{% for var in declaration_vars -%}
	regex.Declare("{{var[0]}}", "{{var[1]}}")
	{% endfor -%}
	{% for var in name_vars -%}
	{{var[0]}}_regex = regex.Compile("{{var[1]}}")
	{% endfor -%}
}

{% for var in name_vars -%}
func {{var[0]}}_func(input string) Token {
	if {{var[0]}}.Match(input) {
	{% if var[2] -%}
		return Token{{"{"}}"{{var[0]}}", {{"output}"}}
	{% elif var[3] -%}
		return Token{{"{"}}"{{var[3]}}"{{"}"}}
	{% else -%}
		return Token{{"{"}}"{{var[0]}}"{{"}"}}
	{% endif -%}
	}
	return Token{{"{NULL}"}}
}

{% endfor %}
