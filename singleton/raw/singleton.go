// singleton pattrn
package pattern

import (
	"fmt"
)

// 写接口的目的：
// 1.为用户在不修改原代码的情况下进行扩展
// 2.客户可以从一个众所周知的访问点访问它
type Singleton interface {
	doSometing()
}

// 定义结构体并实现Singleton接口
type singleton struct {
	something string
}

func (this *singleton) doSometing() {
	fmt.Println("you want to ", this.something)
}

// 产生单例
var oneSingleton Singleton

func NewSingleton(something string) Singleton {
	if oneSingleton == nil {
		oneSingleton = &singleton{something: something}
	}
	return oneSingleton
}
