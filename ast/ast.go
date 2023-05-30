package ast

import "monkey/token"

type Node interface {
	TokenLiteral() string //字面量
}

//Statement 语句（不产生值）
type Statement interface {
	Node
	statementNode()
}

//Expression 表达式（产生值）
type Expression interface {
	Node
	expressionNode()
}

//Program 程序节点（是所有AST的根节点) 由一组语句组成
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	//输入程序的第一个语句字面量
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

//LetStatement let <标识符> = <表达式> ;
type LetStatement struct {
	Token token.Token

	Name  *Identifier //变量标识符
	Value Expression  //产生值的表达式
}

func (l *LetStatement) statementNode() {
}
func (l *LetStatement) TokenLiteral() string {
	return l.Token.Literal
}

//Identifier 标识符
type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode() {

}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

//ReturnStatement return <表达式>;
type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (l *ReturnStatement) statementNode() {
}
func (l *ReturnStatement) TokenLiteral() string {
	return l.Token.Literal
}
