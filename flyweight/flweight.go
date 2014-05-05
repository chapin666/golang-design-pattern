package flyweight

import "fmt"

type Websit interface {
	Use(user User)
}

type ConcreteWebsite struct {
	Websit
	name string
}

func NewConcreteWebsit(name string) ConcreteWebsite {
	return ConcreteWebsite{name: name}
}

func (this *ConcreteWebsite) Use(user User) {
	fmt.Println("网站分类：", this.name, "用户名称：", user.Name())
}
