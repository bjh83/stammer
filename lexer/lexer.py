#tools for generating the syntax.go file from a syntax file
class Lexer:
	fin = None
	fout = None
	declareSec = ''
	functionCodeList = []
	funcList = []
	regexList = []
	typeList = []
	funcIndex = 0
	regexIndex = 0

	def __init__(self, fileName):
		fileOutName = fileName + '.go'
		self.fin = open(fileName, 'r')
		self.fout = open(fileOutName, 'w')

	def strip(self, string):
		out = ''
		strip = True
		index = 0
		while index < len(string):
			letter = string[index]
			index += 1
			if letter != ' ' and letter != '\t' and letter != '\n':
				strip = False
			if not strip and letter != '\n':
				out += letter
		return out

	def split(self, string):
		out = []
		current = ''
		index = 0
		inQuotes = False
		while index < len(string):
			while index < len(string):
				if string[index] == '\"':
					inQuotes = not inQuotes
				if string[index] == ' ' and not inQuotes:
					index += 1
					break
				current += string[index]
				index += 1
			out.append(current)
			current = ''
		return out

	def lex(self):
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
				self.matching(words)
		self.makeFile()

	def declare(self, words):
		self.declareSec += '\n\tregex.Declare(' + words[0] + ', ' + words[2] + ')'

	def matching(self, words):
		referenceId = self.regexHash(words[0])
		self.typeList.append(words[1])
		self.functionCodeList.append('func ' + self.nameGen() + '(input string) (type, valOrId) {\n'
				+ '\tif ' + referenceId + '.match(input) {\n'
				+ '\t\treturn ' + words[1] + ', ' + words[2] + '\n'
				+ '\t}\n'
				+ '\treturn NULL, -1\n'
				+ '}\n')

	def regexHash(self, string):
		referenceId = 'regex' + str(self.regexIndex)
		self.regexIndex += 1
		self.declareSec += '\n\t' + referenceId + ' = regex.Compile(' + string + ')'
		self.regexList.append(referenceId)
		return referenceId

	def nameGen(self):
		name = 'func' + str(self.funcIndex)
		self.funcIndex += 1
		self.funcList.append(name)
		return name

	def makeFile(self):
		self.fout.write('package lexer\n\n'
				+ 'import(\n'
				+ '\t\"../regex\"\n'
				+ ')\n\n')
		self.fout.write('const(\n')
		for tipe in self.typeList:
			self.fout.write('\t' + tipe + ' = iota' + '\n')
		self.fout.write(')\n\n')
		for regex in self.regexList:
			self.fout.write('var ' + regex + ' regex.Regex\n')
		self.fout.write('var funcArray []func(string)(int, int) = []func(string)(int, int) {')
		first = True
		for func in self.funcList:
			if first:
				self.fout.write(func)
				first = False
			self.fout.write(', ' + func)
		self.fout.write('}\n')
		self.fout.write('\nfunc setUp() {\n'
				+ self.declareSec + '\n'
				+ '}\n\n')
		for func in self.functionCodeList:
			self.fout.write(func + '\n')

