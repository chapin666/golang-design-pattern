package state

import "fmt"

type State interface {
	WriteProgram(w Work)
}

type MorningState struct{}

func (this *MorningState) WriteProgram(w Work) {
	if w.Hour() < 12 {
		fmt.Println("当前时间：", w.Hour(), "，精力不错，努力工作")
	} else {
		w.SetState(new(NoonState))
		w.WriteProgram()
	}
}

type NoonState struct{}

func (this *NoonState) WriteProgram(w Work) {
	if w.Hour() < 17 {
		fmt.Println("当前时间：", w.Hour(), "，精力不足，犯困")
	} else {
		w.SetState(new(EveningState))
		w.WriteProgram()
	}
}

type EveningState struct{}

func (this *EveningState) WriteProgram(w Work) {
	if w.Finish() {
		fmt.Println("下班时间到")
	} else {
		fmt.Println("继续加班，啊啊啊啊....")
	}
}

//work
type Work struct {
	state  State
	hour   int
	finish bool
}

func (this *Work) NewWork() {
	this.state = new(MorningState)
	this.finish = false
}
func (this *Work) Hour() int {
	return this.hour
}
func (this *Work) SetHour(hour int) {
	this.hour = hour
}
func (this *Work) Finish() bool {
	return this.finish
}
func (this *Work) SetFinish(finish bool) {
	this.finish = finish
}
func (this *Work) SetState(state State) {
	this.state = state
}
func (this *Work) WriteProgram() {
	//fmt.Println(this.Hour())
	this.state.WriteProgram(*this)
}
