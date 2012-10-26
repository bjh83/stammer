class Lexer:
	fileName = ''
	fileOutName = ''
	fin
	fout
	declareSec = ''
	funcList = []

	def _init_(self, fileName):
		self.fileName = fileName
		self.fileOutName = fileName + '.go'
		self.fin = open(fileName, 'r')
		self.fout = open(fileOutName, 'w')

	def strip(string):
		out = ''
		strip = True
		for letter in string:
			if letter != ' ' or letter != '\t' or letter != '\n':
				strip = False
			if not strip and letter != '\n':
				out += letter
		return out

	def split(string):
		out = []
		current = ''
		index = 0
		while index < len(string):
			while index < len(string):
				if string[index] == ' ':
					index += 1
					break
				current += string[index]
				index += 1
			out.append(current)
			current = ''
		return out

	def lex():
		declaration = False
		matching = False
		for line in self.fin:
			line = self.strip(line)
			words = self.split(line)
			if words[0] == '%D':
				declaration = not declaration
				continue
			if words[0] == '%%':
				matching = not matching
				continue
			if declaration:
				self.declare(words)
			if matching:
				self.match(words)

	def declare(words):
		self.declareSec += '\n\tregex.Declare(' + words[0] + ', ' + words[2] + ')'

	def matching(words):
		referecneId = self.regexHash(words[0])
		self.funcList.append('func ' + self.nameGen() + '(input string) (type, valOrId) {\n'
				+ '\tif ' + referenceId + '.match(input) {\n'
				+ '\t\treturn ' + words[1] + ', ' + words[2] + '\n'
				+ '\t}\n'
				+ '\treturn NULL, -1\n'
				+ '}\n')

	def regexHash(string):
		return referenceId

