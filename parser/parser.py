class SymbolGroup:
	symbolList = []

class Production:
	start = ''
	groupList = []

class Parser:
	fin = None
	fout = None
	productionList = []

	def __init__(self, fileName):
		fileOutName = fileName + '.go'
		fin = open(fileName, 'r')
		fout = open(fileOutName, 'w')

	def
