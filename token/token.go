package token

//TokenType 类型
type TokenType string

type Token struct {
	//类型
	Type TokenType
	//字面量
	Literal string
}

//token类型枚举
const (

	//特殊类型

	ILLEGAL = "ILLEGAL" //未知符号
	EOF     = "EOF"     //文件结尾

	//标识符 字面量

	IDENT  = "IDENT" //标识符
	INT    = "INT"   //int类型
	STRING = "STRING"

	//运算符

	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	LT       = "<"
	GT       = ">"

	EQ     = "=="
	NOT_EQ = "!="

	//分隔符

	COMMA     = ","
	SEMICOLON = ";"
	COLON     = ":"

	LPAREN   = "("
	RPAREN   = ")"
	LBRACE   = "{"
	RBRACE   = "}"
	LBRACKET = "["
	RBRACKET = "]"

	//关键字

	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"

	FOR = "FOR"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
	"for":    FOR,
}

//LookupIdent 判定是否是关键字还是标识符
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
