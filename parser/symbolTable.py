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
	mark = 0
	toReduce = False
	stateGoto = 0

	def __init__(self, simpleProduction):
		self.start = simpleProduction.start
		self.group = simpleProduction.group

	def increment(self):
		productionCopy = copy.deepcopy(self)
		productionCopy.mark += 1
		return productionCopy

	def getMarkedSymbol(self):
		if self.mark < len(self.group.symbolList):
			return self.group.symbolList[mark]
		else:
			self.toReduce = True
			return ''

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
	states = {}

	def add(self, production):
		self.productionList[production.start] = production

	def closure(self, productionName):
		configuratingSet = [productionName]
		if productionName in self.productionList:
			productions = self.productionList[productionName]
			for production in productions:
				symbol = production.getMarkedSymbol()
				if symbol != '':
					configuratingSet.extend(self.closure(symbol))
		return set(configuratingSet)

	def breakDown(self):
		for production in self.productionList:
			self.productionList[production] = StateProduction(
					self.productionList[production].generateSimpleProductions())

	def incrementProductionList(self):
		for productionName in self.productionList:
			newProductions = []
			for production in self.productionList[productionName]:
				newProductions.append(production.increment())
			self.productionList[productionName] = newProductions

	def successor(self, configurationSet, symbol):
