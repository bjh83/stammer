# Contains code for computing a symbolTable given an arbitrary grammar

#TODO: figure out what I am doing with epsilons and the whole bit...
class CompoundProduction:
	name = ''
	groupList = []
	currentGroup = []

	def __init__(self, name):
		self.name = name

	def addSymbol(self, symbol):
		self.currentGroup.append(symbol)

	def endGroup(self):
		groupList.append(currentGroup)
		currentGroup = []

	#Generates a list of simple productions
	def split(self):
		productions = []
		for group in self.groupList:
			productions.append(Production(self.name, group))
		return productions

class Production:
	name = ''
	group = []

	def __init__(self, name):
		self.name = name

	def __init__(self, name, group):
		self.name = name
		self.group = group

	def addSymbol(self, symbol):
		self.group.append(symbol)

# This should allow me to have all the associative info I need
# without having to have tons of different productions
class Pair:
	production = None
	dot = 0
	successor = None
	successorSymbol = None

	def __init__(self, production, dot):
		self.production = production
		self.dot = dot

	def afterDot(self):
		if self.dot < len(self.production.group):
			return self.production.group[self.dot]
		else:
			return None

	def beforeDot(self):
		if self.dot > 0:
			return self.production.group[self.dot]
		else:
			return None

	def advance(self):
		return Pair(self.production, self.dot + 1)

	def successor(self):
		return (successorSymbol, successor)

class ProductionList:
	startSymbol = ''
	terminals = []
	productionList = []

	def __init__(self, startSymbol, terminals):
		self.startSymbol = startSymbol
		self.terminals = terminals

	def generate(self):
		generator = TableGenerator()
		generator.startSymbol = self.startSymbol
		index = 0
		while index < len(self.terminals):
			generator.terminals[self.terminals[index]] = index
			index += 1
		for production in self.productionList:
			generator.productionMap[production.name] = production.split
		return generator

# I want a set that supports list operations
def setify(alist):
	returnList = []
	for element in set(alist):
		returnList.append(element)
	return returnList

class TableGenerator:
	startSymbol = ''
	primeStartSymbol = ''
	symbolList = []
	terminals = {} #Map (symbol:index)
	productionMap = {} #Map (productionName:list of productions

	def closure(self, configList):
		returnList = []
		returnList.extend(configList)
		for pair in returnList:
			productionName = pair.afterDot()
			if productionName is not None and productionName is in self.productionMap: #It is not a terminal
				for production in self.productionMap[productionName]:
					returnList.append(Pair(production, 0))
		return setify(returnList)

	#TODO: refactor so that each pair remembers its successor
	def successor(self, configList, symbol):
		newConfig = []
		toUpdate = []
		for pair in configList:
			productionName = pair.afterDot()
			if productionName is not None and productionMap == symbol:
				newConfig.append(pair.advance())
				toUpdate.append(pair)
		newConfig = self.closure(newConfig)
		for pair in toUpdate:
			pair.successor = newConfig
			pair.successorSymbol = symbol
		return newConfig

	def buildConfigurationSets(self):
		family = [self.closure([Pair(Production(self.startSymbol, [self.primeStrartSymbol]), 0)])]
		for configList in family:
			for symbol in self.symbolList:
				family.append(self.successor(configList, symbol))
		return family

