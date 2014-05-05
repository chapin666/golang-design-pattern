package pattern

type Oper int

const (
	ADD Oper = iota
	SUB
	MUL
	DIV
)

type Operator interface {
	Result() float64
	SetNumA(a float64)
	SetNumB(b float64)
}

// BaseOperator
type BaseOperator struct {
	numberA float64
	numberB float64
}

func (this *BaseOperator) SetNumA(numberA float64) {
	this.numberA = numberA
}
func (this *BaseOperator) SetNumB(numberB float64) {
	this.numberB = numberB
}

// AddOperator
type AddOperator struct {
	BaseOperator
}

func (this *AddOperator) Result() float64 {
	return this.numberA + this.numberB
}

// SubOperator
type SubOperator struct {
	BaseOperator
}

func (this *SubOperator) Result() float64 {
	return this.numberA - this.numberB
}

// MulOperator
type MulOperator struct {
	BaseOperator
}

func (this *MulOperator) Result() float64 {
	return this.numberA * this.numberB
}

// DivOperator
type DivOperator struct {
	BaseOperator
}

func (this *DivOperator) Result() float64 {
	if this.numberB == 0 {
		panic("被除数不能为0")
	}

	return this.numberA / this.numberB
}

//OperatorFactory
type OperatorFactory struct{}

func (this *OperatorFactory) create(oper Oper) (operator Operator) {
	switch oper {
	case ADD:
		operator = new(AddOperator)
	case SUB:
		operator = new(SubOperator)
	case MUL:
		operator = new(MulOperator)
	case DIV:
		operator = new(DivOperator)
	default:
		panic("运算符号错误")
	}
	return
}
