import re
import 

def lex(fileName):
	f = open(fileName, 'r')
	regex = re.compile('%D')
	declare = False
	for line in f:
		if regex.match(line):
			declare = True
			regex = re.compile('\"\w\" D= \".+\"')
		if declare:
			match = re.search('\"\w\" D= \".+\"', line)
			match = re.split(' D= ', match)
