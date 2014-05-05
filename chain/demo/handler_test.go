package demo

import (
	"testing"
)

func TestHandler(t *testing.T) {
	zs := NewCommManager("经理张三")
	ls := NewMajorManager("总监李四")
	zs.SetManager(&ls)

	request := Request{}
	request.setReqContent("请假")
	request.SetReqNumber(2)
	request.SetReqType("请假")

	zs.RequestApplication(request)
}
