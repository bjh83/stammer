class SymbolGroup:
	symbolList = []
	
	def addSymbol(self, symbol):
		symbolList.append(symbol)

class Production:
	start = ''
	groupList = []
	toAppend = None

	def __init__(self, start):
		self.start = start
		toAppend = SymbolGroup()

	def addSymbol(self, symbol):
		toAppend.addSymbol(symbol)

	def endGroup(self):
		groupList.append(toAppend)
		toAppend = SymbolGroup()

	def end(self):
		groupList.append(toAppend)
		toAppend = None

class Parser:
	fin = None
	fout = None
	productionList = []

	def __init__(self, fileName):
		fileOutName = fileName + '.go'
		fin = open(fileName, 'r')
		fout = open(fileOutName, 'w')

	def process(self):
		word = ''
		currentProduction = None
		for line in fin:
			for letter in line:
				if letter == ':':
					currentProduction = Production(word)
					word = ''
				elif letter == '|':
					currentProduction.endGroup()
				elif letter == ';':
					currentProduction.end()
					productionList.append(currentProduction)
					currentProduction = None
				elif letter == ' ':
					currentProduction.append(word)
					word = ''
				else:
					word += letter
