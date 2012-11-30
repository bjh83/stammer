#sample syntax file

class syntax:
	import_statement = '''
		import(
			strconv
		)'''
	regex_variables = [('digit', r'[0-9]'), ('letter', r'[a-zA-Z]')]
	tokens = [
		('constant', r'{digit}+(\.{digit}+)?|\.{digit}+', 'float32', '''
		{
			ret, _ := strconv.ParseFloat($$, 32)
			return float32(ret)
		}'''),
		('identifier', r'{letter}+', 'string', '''
		{
			return $$
		}''')
		('', r'[ /t/n]', '', '') #Throw away whitespace
		('plus', r'\+', '', '')
		('minus', r'-', '', '')
		('times', r'\*', '', '')
		('divide', r'/', '', '')
		]
