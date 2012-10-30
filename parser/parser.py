class Parser:
	fin = None
	fout = None
	termList = []
	productionList = []

	def __init__(self, fileName, termList):
		fileOutName = fileName + '.go'
		self.fin = open(fileName, 'r')
		self.fout = open(fileOutName, 'w')
		f = open(termList)
		for line in termList:
			terminal = ''
			for letter in line:
				if letter != '\n':
					terminal += letter
			self.termList.append(terminal)

	def process(self):
		word = ''
		started = False
		currentProduction = None
		for line in self.fin:
			for letter in line:
				if started:
					if letter == ':':
						currentProduction = Production(word)
						word = ''
					elif letter == '|':
						currentProduction.endGroup()
					elif letter == ';':
						started = False
						currentProduction.end()
						productionList.append(currentProduction)
						currentProduction = None
					elif letter == ' ':
						currentProduction.append(word)
						word = ''
					else:
						word += letter
				else:
					started = True
					word += letter

	def makeFile(self):
		self.fout.write('package parser\n\n')
		self.fout.write('import(\n'
				+ '\t. \"../lexer\"\n'
				+ '\t\"container/list\"\n'
				+ ')\n\n')
		self.fout.write('const(\n')
		for production in self.productionList:
			self.fout.write('\t' + production + '\n')
		self.fout.write(')\n')
		for production in self.productionList:
			self.fout.write('\ntype ' + production.start + ' struct {\n'
					+ '\tLeft, Right ParseNode\n'
					+ '}\n\n')
			self.fout.write('func (node *' + production.start + ') first() list.List {\n'
					+ '\tterminals := list.New()\n'
					+ '\tterminals.Init()\n')
			writeEpsilon = True
			symbolSet = []
			for symbolGroup in production.groupList:
				for symbol in symbolGroup.symbolList:
					symbolSet.append(symbol)
			symbolSet = set(symbolSet)
			for symbol in symbolSet:
				if symbol in self.termList:
					writeEpsilon = False
					self.fout.write('\tterminals.PushBack(' + symbol + ')\n')
			if writeEpsilon:
				self.fout.write('\tterminals.PushBack(Epsilon)\n')
