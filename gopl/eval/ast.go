package eval

import (
	"fmt"
	"math"
)

// Expr 任意一种表达式
type Expr interface {
	Eval(env Env) float64
}

// Var 表达式中的变量
type Var string

// 表达式中常量
type literal float64

//一元操作
type unary struct {
	op rune // + or -
	x Expr
}

//二元操作
type binary struct {
	op rune  // + - * /
	x, y Expr
}

type call struct {
	fn string
	args []Expr
}

type Env map[Var]float64

func (v Var) Eval(env Env) float64 {
	return env[v]
}

func (v literal) Eval(_ Env) float64 {
	return float64(v)
}

func (u unary) Eval(env Env) float64 {
	switch u.op {
	case '+':
		return +u.x.Eval(env)
	case '-':
		return -u.x.Eval(env)
	}
	panic(fmt.Sprintf("invalid unary.op:%v", u.op))
}

func (b binary) Eval(env Env) float64 {
	switch b.op {
	case '+':
		return b.x.Eval(env) + b.y.Eval(env)
	case '-':
		return b.x.Eval(env) - b.y.Eval(env)
	case '*':
		return b.x.Eval(env) * b.y.Eval(env)
	case '/':
		return b.x.Eval(env) / b.y.Eval(env)
	}

	panic(fmt.Sprintf("invalid binary op:%v", b.op))
}

func (c call) Eval(env Env) float64 {
	switch c.fn {
	case "pow":
		return math.Pow(c.args[0].Eval(env), c.args[1].Eval(env))
	case "sin":
		return math.Sin(c.args[0].Eval(env))
	case "sqrt":
		return math.Sqrt(c.args[0].Eval(env))
	}
	panic(fmt.Sprintf("invalid call op:%v", c.fn))
}


