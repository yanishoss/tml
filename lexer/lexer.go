package lexer

const (
	EOF       = "EOF"
	HASH      = "#"
	AT        = "@"
	CROSS     = "x"
	TEXT      = "TEXT"
	IDENT     = "IDENT"
	LABEL     = "LABEL"
	NUMBER    = "NUMBER"
	DELIMITER = "DELIMITER"
)

type Token struct {
	Type    string
	Literal string

	// It is [Start; End[
	Start int

	// End is not included
	End int
}

type Lexer interface {
	NextToken() Token
	PeekToken() Token
}

type lexer struct {
	input   string
	curPos  int
	peekPos int
	tokens  []Token
	tokPos  int
}

func New(input string) Lexer {
	l := &lexer{input, 0, 0, make([]Token, 0), 0}

	l.init()

	l.lex()

	return l
}

func (l *lexer) init() {
	if len(l.input) <= 1 {
		l.peekPos = l.curPos
	} else {
		l.peekPos = l.curPos + 1
	}
}

func (l *lexer) char() string {
	return l.input[l.curPos:l.curPos+1]
}

func (l *lexer) peekChar() string {
	return l.input[l.peekPos:l.peekPos+1]
}

func (l *lexer) readChar() {
	if l.curPos+1 > len(l.input) {
		return
	}

	l.curPos++

	if l.peekPos+1 >= len(l.input) {
		l.peekPos = l.curPos
	} else {
		l.peekPos++
	}
}

func (l *lexer) eatWhitespaces() {
	for !l.eof() {
		ch := l.char()

		if ch != " " && ch != "\t" {
			break
		}

		l.readChar()
	}
}

func (l *lexer) eof() bool {
	return l.curPos >= len(l.input)
}

func (l *lexer) push(t Token) {
	l.tokens = append(l.tokens, t)
}

func (l *lexer) parseText() Token {
	literal := ""
	start := l.curPos

	for !l.eof() {
		ch := l.char()

		if ch == "\n" && (l.peekChar() == "\n" || l.peekChar() == "#") {
			literal += ch
			l.readChar()
			break
		}

		literal += ch

		l.readChar()
	}

	return Token{TEXT, literal, start, l.curPos}
}

func (l *lexer) parseIdent() Token {
	literal := ""
	start := l.curPos

	for !l.eof() {
		ch := l.char()

		if ch == "\n" {
			break
		}

		literal += ch

		l.readChar()
	}

	return Token{IDENT, literal, start, l.curPos}
}

func (l *lexer) parseLabel() Token {
	literal := ""
	start := l.curPos


	for !l.eof() {
		ch := l.char()

		if ch == " " {
			break
		}

		literal += ch

		l.readChar()
	}

	return Token{LABEL, literal, start, l.curPos}
}

func (l *lexer) parseNumber() Token {
	literal := ""
	start := l.curPos

	for !l.eof() {
		ch := l.char()

		if !isDigit(ch) && ch != "." {
			break
		}

		literal += ch

		l.readChar()
	}

	return Token{NUMBER, literal, start, l.curPos}
}

func (l *lexer) lex() {
	for !l.eof() {
		l.eatWhitespaces()

		ch := l.char()

		start := l.curPos

		switch ch {
		case "#":
			l.readChar()
			l.push(Token{HASH, ch, start, l.curPos})
			break
		case "x":
			l.readChar()
			l.push(Token{CROSS, ch, start, l.curPos})
			break
		case "@":
			l.readChar()
			l.push(Token{AT, ch, start, l.curPos})
			break
		case "\n":
			l.readChar()
			l.push(Token{DELIMITER, ch, start, l.curPos})
			break
		default:
			// Parse <exercise name>
			if len(l.tokens) > 0 && l.tokens[len(l.tokens)-1].Type == HASH {
				l.push(l.parseIdent())
				break
			}

			// Parse <unit>
			if len(l.tokens) > 0 && l.tokens[len(l.tokens)-1].Type == NUMBER {
				l.push(l.parseLabel())
				break
			}

			// Parse <number>
			if isDigit(ch) {
				l.push(l.parseNumber())
				break
			}

			// Parse <comments>
			l.push(l.parseText())
		}
	}

	l.push(Token{EOF, "", l.curPos, l.curPos})
}

func (l *lexer) NextToken() Token {
	if l.tokPos >= len(l.tokens) {
		return l.tokens[len(l.tokens)-1]
	}

	t := l.tokens[l.tokPos]

	l.tokPos++

	return t
}

func (l *lexer) PeekToken() Token {
	if l.tokPos >= len(l.tokens) {
		return l.tokens[len(l.tokens)-1]
	}

	t := l.tokens[l.tokPos]

	return t
}

func isDigit(c string) bool {
	return c >= "0" && c <= "9"
}
