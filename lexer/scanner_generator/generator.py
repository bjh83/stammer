from jinja2 import Environment, PackageLoader

def escape(string):
	group = ''
	out = ''
	for char in string:
		if len(group) == 0:
			group += char
			continue
		out += group[:-1]
		group = group[-1] + char
		if group[0] != '\\':
			continue
		if group == r'\a' or group == r'\b' or group == r'\f' or group == r'\n' or \
				group == r'\r' or group == r'\t' or group == r'\v' or group == r'\x' or \
				group == r'\\':
			continue
		group = '\\' + group
	return out + group

class generator:
	regex_variables = []
	tokens = []
	template = None
	fileOut = ''

	def __init__(self, regex_variables, tokens, fileOut='scanner.go'):
		env = Environment(loader=PackageLoader('scanner_generator', 'templates'))
		self.template = env.get_template('scanner.go')
		self.tokens = tokens
		self.regex_variables = regex_variables
		self.fileOut = fileOut

	def get_names(self):
		names = []
		for token in self.tokens:
			name.append(token[0])
		return names

	def escape_regexs(self):
		new_list = []
		for group in self.tokens:
			new_group = (group[0], escape(group[1]), group[2], group[3])
			new_list.append(new_group)
		self.tokens = new_list

	def build(self):
		self.escape_regexs()
		self.template.stream(name_vars=self.tokens, declaration_vars=self.regex_variables).dump(self.fileOut)

