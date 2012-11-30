#sample syntax file
import generator

def render():
	regex_variables = [('digit', r'[0-9]'), ('letter', r'[a-zA-Z]')]
	tokens = [
		('constant', r'{digit}+(\.{digit}+)?|\.{digit}+', 
		'''output, _ := strconv.ParseFloat(input, 32)''', ''),
		('identifier', r'{letter}+', '''output = input''', ''),
		('whitespace', r'[ \t\n]', '', 'IGNORE'), #Throw away whitespace
		('plus', r'\+', '', ''),
		('minus', r'-', '', ''),
		('times', r'\*', '', ''),
		('divide', r'/', '', ''),
		]
	generator.generator(regex_variables, tokens).build()

render()
