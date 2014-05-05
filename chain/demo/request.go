package demo

type Request struct {
	reqType    string
	reqContent string
	reqNumber  int
}

func (this *Request) ReqType() string {
	return this.reqType
}
func (this *Request) SetReqType(reqType string) {
	this.reqType = reqType
}
func (this *Request) ReqContent() string {
	return this.reqContent
}
func (this *Request) setReqContent(reqContent string) {
	this.reqContent = reqContent
}

func (this *Request) ReqNumber() int {
	return this.reqNumber
}
func (this *Request) SetReqNumber(number int) {
	this.reqNumber = number
}
