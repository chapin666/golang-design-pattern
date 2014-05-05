package decorator

import (
	"fmt"
	"testing"
)

func TestDecorator(t *testing.T) {

	// 生成一部手机
	android := &Android{version: "4.4"}
	fmt.Println("没有装饰之前：")
	android.showMe()

	//给手机加上DIY,成为带DIY的手机
	diy := new(DIY)
	diy.Decorate(android)

	//给带DIY的手机再加上屏保
	protection := new(SrcProtection)
	protection.Decorate(diy)

	//显示手机信息
	fmt.Println("没有装饰之后：")
	protection.showMe()

	//生产苹果手机
	ios := &IOS{version: "5"}
	//加上屏保
	protection = new(SrcProtection)
	protection.Decorate(ios)
	fmt.Println()
	protection.showMe()
}
