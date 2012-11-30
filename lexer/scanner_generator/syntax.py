#sample syntax file
import generator

def render():
	regex_variables = [('digit', r'[0-9]'), ('letter', r'[a-zA-Z]')]
	tokens = [
		('constant', r'{digit}+(\.{digit}+)?|\.{digit}+', '''
		{
			ret, _ := strconv.ParseFloat($$, 32)
			return float32(ret)
		}'''),
		('identifier', r'{letter}+', '''
		{
			return $$
		}'''),
		('', r'[ /t/n]', ''), #Throw away whitespace
		('plus', r'\+', ''),
		('minus', r'-', ''),
		('times', r'\*', ''),
		('divide', r'/', ''),
		]
	generator.generator(regex_variables, tokens).build()

