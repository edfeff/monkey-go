package ast

import (
	"bytes"
	"monkey/token"
	"strings"
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

//InfixExpression 中缀表达式 <左表达式> <操作符> <右表达式>
type InfixExpression struct {
	Token    token.Token
	Left     Expression
	Operator string
	Right    Expression
}

func (i *InfixExpression) TokenLiteral() string {
	return i.Token.Literal
}

func (i *InfixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(i.Left.String())
	out.WriteString(" " + i.Operator + " ")
	out.WriteString(i.Right.String())
	out.WriteString(")")
	return out.String()
}

func (i *InfixExpression) expressionNode() {
}

type Boolean struct {
	Token token.Token
	Value bool
}

func (b *Boolean) TokenLiteral() string {
	return b.Token.Literal
}

func (b *Boolean) String() string {
	return b.Token.Literal
}

func (b *Boolean) expressionNode() {
}

// if ( <条件> ) <结果> else <可替代的结果>

//IfExpression if表达式
type IfExpression struct {
	Token       token.Token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (i *IfExpression) TokenLiteral() string {
	return i.Token.Literal
}

func (i *IfExpression) String() string {
	var out bytes.Buffer
	out.WriteString("if")
	out.WriteString(i.Condition.String())
	out.WriteString(" ")
	out.WriteString(i.Consequence.String())

	if i.Alternative != nil {
		out.WriteString("else ")
		out.WriteString(i.Alternative.String())
	}
	return out.String()
}

func (i *IfExpression) expressionNode() {
}

type BlockStatement struct {
	Token      token.Token
	Statements []Statement
}

func (b *BlockStatement) TokenLiteral() string {
	return b.Token.Literal
}

func (b *BlockStatement) String() string {
	var out bytes.Buffer
	for _, s := range b.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

func (b *BlockStatement) statementNode() {
}

// fn <参数列表> <块语句>
//(<参数1>, <参数2>, <参数3> ...)

type FunctionLiteral struct {
	Token      token.Token
	Parameters []*Identifier
	Body       *BlockStatement
}

func (f *FunctionLiteral) TokenLiteral() string {
	return f.Token.Literal
}

func (f *FunctionLiteral) String() string {
	var out bytes.Buffer
	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}
	out.WriteString(f.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ","))
	out.WriteString(")")
	out.WriteString(f.Body.String())
	return out.String()
}

func (f *FunctionLiteral) expressionNode() {
}

type CallExpression struct {
	Token     token.Token // The '(' token
	Function  Expression  // Identifier or FunctionLiteral
	Arguments []Expression
}

func (ce *CallExpression) expressionNode()      {}
func (ce *CallExpression) TokenLiteral() string { return ce.Token.Literal }
func (ce *CallExpression) String() string {
	var out bytes.Buffer

	args := []string{}
	for _, a := range ce.Arguments {
		args = append(args, a.String())
	}

	out.WriteString(ce.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")

	return out.String()
}
