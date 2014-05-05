package main

import (
	"fmt"
	"pattern/observer/obs"
	"time"
)

type MClass struct {
	dispatcher tbs.Dispatcher
}

func main() {
	mc := &MClass{}
	mc.Start()
}

func (this *MClass) Start() {
	//获取分派器单例
	dispatcher := tbs.SharedDispatcher()

	//添加监听1
	var fun1 tbs.EventCallback = this.onTest
	dispatcher.AddEventListener("test", &fun1)

	//再添加监听2
	var fun2 tbs.EventCallback = this.onTest2
	dispatcher.AddEventListener("test", &fun2)

	//随便弄个事件携带的参数，我把参数定义为一个map
	params := make(map[string]interface{})
	params["id"] = 1000
	//创建一个事件对象
	event := tbs.CreateEvent("test", params)
	//把事件分派出去
	dispatcher.DispatchEvent(event)

	//移除监听1
	dispatcher.RemoveEventListener("test", &fun1)

	//再把事件分派出去一次
	dispatcher.DispatchEvent(event)

	//因为主线程不会等子线程而直接关闭进程，这样会看不到效果，所以我在这里加了阻塞式延时
	time.Sleep(time.Second * 1)
}

//回调出得到的就是一个event对象了
func (this *MClass) onTest(event *tbs.Event) {
	fmt.Println("onTest", event.Params["id"])
}

func (this *MClass) onTest2(event *tbs.Event) {
	fmt.Println("onTest2", event.Params["id"])
}
