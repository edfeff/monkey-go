package parser

import (
	"fmt"
	"monkey/ast"
	"monkey/lexer"
	"monkey/token"
)

type Parser struct {
	l         *lexer.Lexer //词法分析器
	curToken  token.Token  //当前的token
	peekToken token.Token  //下一个即将读取的token
	errors    []string     //解析中出现的错误
}

//New 创建一个解析器
func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}
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
		return nil
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
