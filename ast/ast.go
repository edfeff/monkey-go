package ast

import (
	"bytes"
	"monkey/token"
)

type Node interface {
	TokenLiteral() string //字面量
	String() string       //Debug展示值
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

func (p *Program) String() string {
	var out bytes.Buffer
	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

//LetStatement let <标识符> = <表达式> ;
type LetStatement struct {
	Token token.Token

	Name  *Identifier //变量标识符
	Value Expression  //产生值的表达式
}

func (l *LetStatement) String() string {
	var out bytes.Buffer
	out.WriteString(l.TokenLiteral() + " ")
	out.WriteString(l.Name.String())
	out.WriteString(" = ")
	if l.Value != nil {
		out.WriteString(l.Value.String())
	}
	out.WriteString(";")
	return out.String()
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

func (i *Identifier) String() string {
	return i.Value
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

func (l *ReturnStatement) String() string {
	var out bytes.Buffer
	out.WriteString(l.TokenLiteral() + " ")
	if l.ReturnValue != nil {
		out.WriteString(l.ReturnValue.String())
	}
	out.WriteString(";")
	return out.String()
}

func (l *ReturnStatement) statementNode() {
}
func (l *ReturnStatement) TokenLiteral() string {
	return l.Token.Literal
}

//ExpressionStatement 表达式
type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (l *ExpressionStatement) String() string {
	if l.Expression != nil {
		return l.Expression.String()
	}
	return ""
}

func (l *ExpressionStatement) statementNode() {
}
func (l *ExpressionStatement) TokenLiteral() string {
	return l.Token.Literal
}

type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (i *IntegerLiteral) TokenLiteral() string {
	return i.Token.Literal
}

func (i *IntegerLiteral) String() string {
	return i.Token.Literal
}

func (i *IntegerLiteral) expressionNode() {
}

//PrefixExpression 前缀表达式，由前缀token+表达式组成
type PrefixExpression struct {
	Token    token.Token
	Operator string     //前缀操作符号
	Right    Expression //紧随的表达式
}

func (p *PrefixExpression) TokenLiteral() string {
	return p.Token.Literal
}

func (p *PrefixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(p.Operator)
	out.WriteString(p.Right.String())
	out.WriteString(")")
	return out.String()
}

func (p *PrefixExpression) expressionNode() {
}
