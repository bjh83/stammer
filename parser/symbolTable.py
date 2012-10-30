import copy

class symbolTable:
	statesTable = {}

	def addTransition(self, symbol, production):
		if symbol in statesTable:
			productionList = statesTable[symbol]
			productionList.append(production)
			statesTable[symbol] = productionList
		else:
			statesTable[symbol] = [production]

class SymbolGroup:
	symbolList = []
	
	def addSymbol(self, symbol):
		symbolList.append(symbol)

class SimpleProduction:
	start = ''
	group = None

	def __init__(self, start):
		self.start = start
		self.group = SymbolGroup()

	def __init__(self, start, group):
		self.start = start
		self.group = group

	def addSymbol(self, symbol):
		group.addSymbol(symbol)

class StateProduction(SimpleProduction):
	state = 0

	def __init__(self, simpleProduction):
		self.start = simpleProduction.start
		self.group = simpleProduction.group

	def increment(self):
		productionCopy = copy.deepcopy(self)
		productionCopy.mark += 1
		return productionCopy

class ComplexProduction(SimpleProduciton):
	groupList = []

	def endGroup(self):
		if len(group) == 0:
			self.group.addSymbol('Epsilon')
		groupList.append(self.group)
		self.group = SymbolGroup()

	def end(self):
		if len(self.group) == 0:
			self.group.addSymbol('Epsilon')
		groupList.append(self.group)
		self.group = None

	def generateSimpleProductions(self):
		productionList = []
		for group in groupList:
			productionList.append(SimpleProduction(start, group))
		return productionList

class ProductionList:
	productionList = {}
	simpleProductionList
	states = {}

	def add(self, production):
		self.productionList[production.start] = production

	def firstClosure(self, productionName):
		firstList = []
		if productionName in productionList:
			production = self.productionList[productionName]
			symbolList = production.getSymbols(0)
			for symbol in symbolList:
				firstList.append(self.first(symbol))
			return set(firstList)
		else:
			return [productionName] #We found a terminal! BTW, this also takes care of epsilon!

	def follow(self, productionName, place):
		firstList = []
		if productionName in self.productionList:
			production = self.productionList[productionName]
			for group in production.groupList:
				symbolList = group.symbolList
				if place < len(symbolList):
					self.states[symbolList[place]]
