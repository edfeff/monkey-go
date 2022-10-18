package lexer

import "monkey/token"

type Lexer struct {
	input    string
	position int //所输入字符串中的当前位置（指向当前字符）
	// line     int
	readPosition int  //所输入字符串中的当前读取位置（指向当前字符之后的一个字符
	ch           byte //当前正在查看的字符 仅支持ASCII
}

// New 创建词法解析器
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar() //尝试读取一个
	return l
}

// NextToken 读取下一个Token
func (l *Lexer) NextToken() token.Token {
	var t token.Token
	switch l.ch {
	case '=':
		t = newToken(token.ASSIGN, l.ch)
	case ';':
		t = newToken(token.SEMICOLON, l.ch)
	case '(':
		t = newToken(token.LPAREN, l.ch)
	case ')':
		t = newToken(token.RPAREN, l.ch)
	case ',':
		t = newToken(token.COMMA, l.ch)
	case '{':
		t = newToken(token.LBRACE, l.ch)
	case '+':
		t = newToken(token.PLUS, l.ch)
	case '}':
		t = newToken(token.RBRACE, l.ch)
	case 0:
		t.Literal = ""
		t.Type = token.EOF
	}
	l.readChar()
	return t
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{tokenType, string(ch)}
}

// readChar 读取字符到ch中
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition //position 是当前字符
	l.readPosition += 1         //readPosition 是下一个字符
}
