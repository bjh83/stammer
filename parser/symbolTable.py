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
	startSymbol = ''
	terminals = []
	productionList = {}
	symbolList = []
	saveStack = []
	configList = []

	def __init__(self, startSymbol, terminals):
		self.terminals = terminals
		self.startSymbol = startSymbol

	def add(self, production):
		self.productionList[production.start] = production

	def save(self):
		self.saveStack.append(self.productionList)
		self.productionList = copy.deepcopy(self.productionList)

	def unsave(self):
		self.productionList = self.saveStack.pop()

	def findSymbols(self):
		self.symbolList.extend(terminals)
		for productionName in self.productionList:
			self.symbolList.add(productionName)

	#TODO: This does not handle closures with respect to marks properly
	def closure(self, productionName):
		configuratingSet = [productionName]
		if productionName in self.productionList:
			productions = self.productionList[productionName]
			for production in productions:
				symbol = production.getMarkedSymbol()
				if symbol != '':
					configuratingSet.extend(self.closure(symbol))
		return configuratingSet

	def breakDown(self):
		self.save()
		for production in self.productionList:
			self.productionList[production] = StateProduction(
					self.productionList[production].generateSimpleProductions())

	def increment(self):
		for productionName in self.productionList:
			newProductions = []
			for production in self.productionList[productionName]:
				newProductions.append(production.increment())
			self.productionList[productionName] = newProductions

	def successor(self, configurationSet, symbol):
		elementsWithSymbol = []
		for productionName in self.productionList:
			for production in self.productionList[productionName]:
				if production.getMarkedSymbol() == symbol:
					elementsWithSymbol.append(production)
		incElementsWithSymbol = []
		for production in elementsWithSymbol:
			incElementWithSymbol.append(production.increment())
		elementsWithSymbol = []
		self.save() #XXX:must unsave()!!!
		self.increment()
		for production in incElementsWithSymbol:
			elementsWithSymbol.extend(self.closure(production.start))
		self.unsave()
		return elementsWithSymbol

	def buildConfiguratingSets(self):
		startPrime = self.startSymbol + '\''
		self.productionList[startPrime] = StateProduction(SimpleProduction(startPrime, self.startSymbol))
		self.startSymbol = startPrime
		configStack = [self.closure(startPrime)]
		configList = configStack
		while len(configStack) > 0:
			currentConfig = configStack.pop()
			for symbol in self.symbolList:
				toAppend = self.successor(currentConfig, symbol)
				configList.append(toAppend)
				configStack.append(toAppend)
		return configList

	#TODO: So I basically need to define a better way to keep track of states
	# since the current method is messy, inflexible, and does not provide a
	# way to keep association with its successor
	def constructTable(self):
		configList = self.buildConfigurationSets()
		symbolTable = SymbolTable(len(self.symbolList), len(configList))
		configIndex = 0
		while configIndex < len(configList):
			configState = configList[configIndex]
			for production in configState:
				if production.mark > 0:
					symbol = production.group[production.mark - 1]
					#TODO: Need to figure out what the reduction is!!!
					symbolTable.setRow(configIndex, 'reduce')
				if production.start == self.start and production.mark == len(production.group):
					symbolTable.add(symbols.index('$'), configIndex, 'accept')
				if #TODO: Figure out how to match states
				if #TODO: Same as above
			configIndex += 1
		return symbolTable

class SymbolTable:
	table = []
	size = 0
	xMult = 0

	def __init__(self, numOfSymbols, numOfStates):
		self.size = numOfSymbols
		self.xMult = numOfSymbols
		index = 0
		while index < self.size:
			table.append(0)
			index++

	def add(self, x, y, element):
		self.table[x + y * self.xMult] = element

	def setRow(self, y, element):
		x = 0
		while x < self.xMult:
			self.add(x, y, element)
			x += 1

	def get(self, x, y):
		return self.table[x + y * self.xMult]

