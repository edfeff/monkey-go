package lexer

import (
	"monkey/token"
)

//Lexer 词法解析器
type Lexer struct {
	input        string //代码
	position     int    //当前位置(当前字符的位置）
	readPosition int    //当前的读取位置(词法解析需要预读）词法分析器除了查看当前字符，还需要进一步“查看”字符串，即查看字符串中的下一个字符
	ch           byte   //正在查看的字符
}

//New 创建词法解析器
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

//NextToken 解析代码，输出解析出的token，包括类型和字面量
func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	//跳过空白符
	l.skipWhitespace()
	switch l.ch {
	case '=':
		if '=' == l.peekChar() {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.EQ, Literal: literal}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '!':
		if '=' == l.peekChar() {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.NOT_EQ, Literal: literal}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	case '"':
		tok.Type = token.STRING
		tok.Literal = l.readString()
	case '[':
		tok = newToken(token.LBRACKET, l.ch)
	case ']':
		tok = newToken(token.RBRACKET, l.ch)
	case ':':
		tok = newToken(token.COLON, l.ch)
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}
	l.readChar()
	return tok
}

//isDigit  是否是[0-9]
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

//isLetter  是否是[a-zA-Z_]
func isLetter(ch byte) bool {
	return ('a' <= ch && ch <= 'z') ||
		('A' <= ch && ch <= 'Z') ||
		(ch == '_')
}

//newToken 创建指定类型和字面量的token
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

//readChar  读取下一个字符并设置到Lexer中，遇到末尾返回0
//该词法分析器仅支持ASCII字符，不能支持所有的Unicode字符。
//这么做也是为了让事情保持简单，让我们能够专注于解释器的基础部分。
//如果要完全支持Unicode和UTF-8，就要将l.ch的类型从byte改为rune，同时还要修改读取下一个字符的方式
func (l *Lexer) readChar() {
	//1 是否读取结束
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		//2 正常读取
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

//peekChar 瞅一眼马上要读取的字符
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

//readIdentifier 读取连续的letter字符
func (l *Lexer) readIdentifier() string {
	startPosition := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[startPosition:l.position]
}

//skipWhitespace 跳过空白字符
func (l *Lexer) skipWhitespace() {
	for isWhitespace(l.ch) {
		l.readChar()
	}
}

//readNumber 读取数字，仅支持整数
func (l *Lexer) readNumber() string {
	startPosition := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[startPosition:l.position]
}

func (l *Lexer) readString() string {
	position := l.position + 1
	for {
		l.readChar()
		if l.ch == '"' || l.ch == 0 {
			break
		}
	}
	return l.input[position:l.position]
}

//isWhitespace 是否是空白符
func isWhitespace(ch byte) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}
