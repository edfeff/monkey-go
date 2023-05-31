package parser

import (
	"fmt"
	"monkey/ast"
	"monkey/lexer"
	"monkey/token"
	"strconv"
)

const (
	_ int = iota //优先级常量定义 数值越大优先级越高
	LOWEST
	EQUALS      //==
	LESSGREATER // > <
	SUM         //+
	PRODUCT     //*
	PREFIX      //-X !X
	CALL        //myFunction(X)
)

type Parser struct {
	l      *lexer.Lexer //词法分析器
	errors []string     //解析中出现的错误

	curToken  token.Token //当前的token
	peekToken token.Token //下一个即将读取的token

	prefixParseFns map[token.TokenType]prefixParseFn
	inParseFns     map[token.TokenType]inParseFn
}

type (
	prefixParseFn func() ast.Expression
	inParseFn     func(ast.Expression) ast.Expression
)

//New 创建一个解析器
func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}
	//初始化
	// 注册前缀解析函数
	p.prefixParseFns = make(map[token.TokenType]prefixParseFn)
	p.registerPrefix(token.IDENT, p.parseIdentifier)   //处理标识符
	p.registerPrefix(token.INT, p.parseIntegerLiteral) //处理整数
	p.registerPrefix(token.BANG, p.parsePrefixExpression)
	p.registerPrefix(token.MINUS, p.parsePrefixExpression)

	// 注册中缀解析函数
	//p.inParseFns = make(map[token.TokenType]inParseFn)

	//读取出两个token，用于初始化cur和peek
	p.nextToken()
	p.nextToken()
	return p
}
func (p *Parser) Errors() []string {
	return p.errors
}

//ParseProgram  解析程序
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}
	//消费所有的token，直到结束EOF
	for !p.curTokenIs(token.EOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}
	return program
}

//parseStatement 解析语句
func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	//let
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return p.parseExpressionStatement()
	}
}

//parseLetStatement 解析let语句
func (p *Parser) parseLetStatement() *ast.LetStatement {
	//let 标识符 = 表达式
	stmt := &ast.LetStatement{Token: p.curToken} //存储当前let的对应的token

	// let后必须是标识符
	if !p.exceptPeek(token.IDENT) {
		return nil
	}
	//存储标识符
	stmt.Name = &ast.Identifier{
		Token: p.curToken,
		Value: p.curToken.Literal,
	}
	//标识符后必须是赋值符号
	if !p.exceptPeek(token.ASSIGN) {
		return nil
	}
	//解析表达式并存储
	//todo 跳过表达式处理

	//读取直到语句结束分号
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.curToken}
	p.nextToken()

	//todo 跳过表达式
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

// parseExpressionStatement 解析表达式陈故居
func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.curToken}
	stmt.Expression = p.parseExpression(LOWEST)
	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

//解析辅助方法

//检查当前token类型
func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

//检查下一个token类型
func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

//nextToken 读取下一个token
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

//断言下一个token类型，类型正确时会自动取出下一个token
func (p *Parser) exceptPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("excepted nex token to be %s, got %s instead", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

//registerPrefix 注册前缀解析函数
func (p *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

//registerInfix 注册中缀解析函数
func (p *Parser) registerInfix(tokenType token.TokenType, fn inParseFn) {
	p.inParseFns[tokenType] = fn
}

//parseExpression 按照优先级解析语句
func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefix := p.prefixParseFns[p.curToken.Type]
	if prefix == nil {
		p.noPrefixParseFnError(p.curToken.Type)
		return nil
	}
	leftExp := prefix()
	return leftExp
}

//parseIdentifier 解析标识符
func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
}

//parseIntegerLiteral 解析整形字面量
func (p *Parser) parseIntegerLiteral() ast.Expression {
	lit := &ast.IntegerLiteral{Token: p.curToken}
	value, err := strconv.ParseInt(p.curToken.Literal, 0, 64)
	if err != nil {
		msg := fmt.Sprintf("could not parse %s as integer", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}
	lit.Value = value
	return lit
}

//noPrefixParseFnError 处理不能解析的token错误
func (p *Parser) noPrefixParseFnError(t token.TokenType) {
	msg := fmt.Sprintf("no prefix parse function for %s found", t)
	p.errors = append(p.errors, msg)
}

//parsePrefixExpression 解析前缀的表达式
func (p *Parser) parsePrefixExpression() ast.Expression {
	//创建一个前缀表达式
	expression := &ast.PrefixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
	}
	p.nextToken()
	//递归解析出右侧的表达式
	expression.Right = p.parseExpression(PREFIX)
	return expression
}
