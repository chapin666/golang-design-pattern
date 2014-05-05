package pattern

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	fmt.Print("加法测试:")
	var factory OperatorFactory
	oper := factory.create(ADD)
	oper.SetNumA(12)
	oper.SetNumB(13)
	fmt.Println("12+13=", oper.Result())
}

func TestSub(t *testing.T) {
	fmt.Print("减法测试:")
	var factory OperatorFactory
	oper := factory.create(SUB)
	oper.SetNumA(12)
	oper.SetNumB(100)
	fmt.Println("12-100=", oper.Result())
}

func TestMul(t *testing.T) {
	fmt.Print("乘法测试:")
	var factory OperatorFactory
	oper := factory.create(MUL)
	oper.SetNumA(1001)
	oper.SetNumB(100)
	fmt.Println("1001x100=", oper.Result())
}

func TestDiv(t *testing.T) {
	fmt.Print("除法测试:")
	var factory OperatorFactory
	oper := factory.create(DIV)
	oper.SetNumA(54)
	oper.SetNumB(12)
	fmt.Println("54 / 12 = ", oper.Result())

	//oper.SetNumA(54)
	//oper.SetNumB(0)
	//fmt.Println("54 / 0 = ", oper.Result())
}
