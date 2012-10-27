class Lexer:
	fileName = ''
	fileOutName = ''
	fin
	fout
	declareSec = ''
	functionCodeList = []
	regexList = []
	funcIndex = 0
	regexIndex = 0

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
		self.functionCodeList.append('func ' + self.nameGen() + '(input string) (type, valOrId) {\n'
				+ '\tif ' + referenceId + '.match(input) {\n'
				+ '\t\treturn ' + words[1] + ', ' + words[2] + '\n'
				+ '\t}\n'
				+ '\treturn NULL, -1\n'
				+ '}\n')

	def regexHash(string):
		referenceId = 'regex' + self.regexIndex
		self.regexIndex += 1
		self.declareSec += '\n\t' + referenceId + ' = regex.Compile(' + string + ')'
		regexList.append(referenceId)
		return referenceId

	def nameGen():
		name = 'func' + func + self.funcIndex
		self.funcIndex += 1
		funcList.append(name)
		return name

	def makeFile():
		self.regexList.reverse()
		self.fout.write('package lexer\n\n'
				+ 'import(\n
				\t\"../regex\"\n
				)\n\n')
		for regex in self.regexList:
			self.fout.write('var ' + regex + ' regex.Regex\n')
		self.fout.write('var funcArray []func(string)(int, int) = []func(string)(int, int) {')
		for func in self.funcList:
			self.fout.write(func + '
		self.fout.write('\nfunc setUp() {\n'
				+ self.declareSec + '\n'
				+ '}\n\n')
		for func in self.functionCodeList:
			self.fout.write(func + '\n')

