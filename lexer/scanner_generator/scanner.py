class scanner:
	import_statement = ''
	regex_variables = []
	tokens = []
	fileOut = None

	def __init__(self, fileName):
		self.fileOut = open(fileName, 'w')

	def writeImports(self):
		self.fileOut.write(import_statement)

	def getRegexs(self):
		regexList = []
		nameless = 0
		for token in self.tokens:
			if token[0] == ''
				tokenList.append(('regex_' + nameless, token[1]))
			else
				tokenList.append((token[0] + '_regex', token[1]))
		return regexList

	def buildTokenFuncs(self):
		tokenFuncList = []
		for token in self.tokens:
			tokenFuncList.append(
