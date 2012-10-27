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

class Parser:
	fin = None
	fout = None
	productionList = []

	def __init__(self, fileName):
		fileOutName = fileName + '.go'
		fin = open(fileName, 'r')
		fout = open(fileOutName, 'w')

	def process(self):
