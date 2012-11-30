from jinja2 import Environment, PackageLoader

class generator:
	regex_variables = []
	tokens = []
	template = None
	fileOut = ''

	def __init__(self, regex_variables, tokens, fileOut='scanner.go'):
		env = Environment(loader=PackageLoader('templates', 'templates'))
		self.template = env.get_template('scanner.go')
		self.tokens = tokens
		self.regex_variables = regex_variables
		self.fileOut = fileOut

	def get_names(self):
		names = []
		for token in self.tokens:
			name.append(token[0])
		return names

	def build(self):
		self.template.stream(name_vars=self.tokens, declaration_vars=self.regex_variables).dump(self.fileOut)

