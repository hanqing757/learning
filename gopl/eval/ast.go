package eval

import (
	"fmt"
	"math"
	"strings"
)

// Expr 任意一种表达式
type Expr interface {
	Eval(env Env) float64
	Check(env Env) error
	String() string
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
func (v Var) Check(env Env) error  {
	if _, ok := env[v]; !ok {
		return fmt.Errorf("undefined variable:%v", v)
	}
	return nil
}
func (v Var) String() string {
	return string(v)
}

func (v literal) Eval(_ Env) float64 {
	return float64(v)
}

func (v literal) Check(_ Env) error {
	return nil
}
func (v literal) String() string {
	return fmt.Sprintf("%g",v)
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
func (u unary) String() string {
	return fmt.Sprintf("%c%s",u.op, u.x)
}

func (u unary) Check(env Env) error {
	if u.op != '+' &&  u.op != '-' {
		return fmt.Errorf("undefined op:%v", u.op)
	}
	return u.x.Check(env)
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

func (b binary) Check(env Env) error {
	if !strings.ContainsRune("+-*/",b.op) {
		return fmt.Errorf("undefined op:%v", b.op)
	}
	if err := b.x.Check(env); err != nil{
		return err
	}

	if err := b.y.Check(env); err != nil{
		return err
	}else {
		if b.op == '/' && b.y.Eval(env) == 0 {
			return fmt.Errorf("divided by zero")
		}
	}
	return nil
}

func (b binary) String() string {
	by := b.y.String()
	return fmt.Sprintf("%s %c %s", b.x, b.op, by)
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

var fnArgsMap = map[string]int{"pow":2, "sin":1, "sqrt":1}

func (c call) Check(env Env) error  {
	if v, ok := fnArgsMap[c.fn]; !ok {
		return fmt.Errorf("undefined fn:%v", c.fn)
	}else if len(c.args) != v {
		return fmt.Errorf("func args cnt invalid")
	}

	for _, ar := range c.args {
		if err := ar.Check(env); err != nil{
			return err
		}
	}
	return nil
}

func (c call) String() string {
	var argsArr []string
	for _, ar := range c.args {
		argsArr = append(argsArr, ar.String())
	}
	return fmt.Sprintf("%s(%s)", c.fn, strings.Join(argsArr, ", "))
}
