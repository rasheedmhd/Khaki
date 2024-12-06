package lexer

type Lexer struct {
	input        string
	position     int // current char 
	readPosition int // next char, used mostly for peeking 
	char         byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	return l
}
