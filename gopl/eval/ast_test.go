package eval

import (
	"fmt"
	"math"
	"testing"
)

func TestEval(t *testing.T) {
	tests := []struct {
		expr string
		env Env
		want string
	}{
		{"sqrt(A / pi)", Env{"A": 87616, "pi": math.Pi}, "167"},
		{"pow(x, 3) + pow(y, 3)", Env{"x": 12, "y": 1}, "1729"},
		{"pow(x, 3) + pow(y, 3)", Env{"x": 9, "y": 10}, "1729"},
		{"5 / 9 * (F - 32)", Env{"F": -40}, "-40"},
		{"5 / 9 * (F - 32)", Env{"F": 32}, "0"},
		{"5 / 9 * (F - 32)", Env{"F": 212}, "100"},
	}
	for _, test := range tests {
		expr, err := Parse(test.expr)
		if err != nil{
			t.Error(err)
			continue
		}
		got := fmt.Sprintf("%.6g", expr.Eval(test.env))
		if got != test.want {
			t.Errorf("expr:%v, env:%v, got:%v, want:%v", test.expr, test.env, got, test.want)
		}
	}
}

func TestCheck(t *testing.T) {
	tests := []struct{
		expr string
		env Env
	}{
		{"x ^ 2", Env{"x":1}},
		{"2 / (x - 1)", Env{"x":1}},
		{"!x",Env{"x":1}},
		{"log(10)",Env{}},
		{"sqrt(1, 2)", Env{}},
		{"x / 3", Env{"y":1}},
	}

	for _, test := range tests {
		expr, err := Parse(test.expr)
		if err != nil{
			t.Log(err)
			continue
		}
		err = expr.Check(test.env)
		if err == nil{
			t.Error("err should not be nil")
		}else {
			t.Log(err)
		}
	}
}

func TestString(t *testing.T) {
	tests := []struct {
		expr string
		env Env
		want string
	}{
		{"sqrt(A / pi)", Env{"A": 87616, "pi": math.Pi}, "167"},
		{"pow(x, 3) + pow(y, 3)", Env{"x": 12, "y": 1}, "1729"},
		{"pow(x, 3) + pow(y, 3)", Env{"x": 9, "y": 10}, "1729"},
		//{"5 / 9 * (F - 32)", Env{"F": -40}, "-40"},
		//{"5 / 9 * (F - 32)", Env{"F": 32}, "0"},
		//{"5 / 9 * (F - 32)", Env{"F": 212}, "100"},
	}
	for _, test := range tests {
		expr, err := Parse(test.expr)
		if err != nil{
			t.Error(err)
			continue
		}
		got := expr.String()
		if got != test.expr {
			t.Errorf("expr:%v, after string:%v", test.expr, got)
		}
	}
}