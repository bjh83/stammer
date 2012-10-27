import lexer

print 'generating syntax.go'
generator = lexer.Lexer('syntax')
generator.lex()
print 'generation complete'

