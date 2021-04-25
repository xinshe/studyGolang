package main

import "fmt"

// append() 为切片追加元素
/*
Go语言的内建函数append()可以为切片动态添加元素。 可以一次添加一个元素，可以添加多个元素，也可以添加另一个切片中的元素（后面加…）。

*/

func main() {
	s := []string{"上海", "北京", "广州"}
	fmt.Printf("s:%v  len(s):%v  cap(s):%v \n", s, len(s), cap(s)) // s:[上海 北京 广州]  len(s):3  cap(s):3

	// 添加一个元素
	s = append(s, "深圳")
	fmt.Printf("s:%v  len(s):%v  cap(s):%v \n", s, len(s), cap(s)) // s:[上海 北京 广州 深圳]  len(s):4  cap(s):6

	// 添加多个元素
	s = append(s, "长沙", "武汉", "郑州")
	fmt.Printf("s:%v  len(s):%v  cap(s):%v \n", s, len(s), cap(s)) // s:[上海 北京 广州 深圳 长沙 武汉 郑州]  len(s):7  cap(s):12

	// 添加另一个切片中的元素（后面加…）
	s1 := []string{"成都", "重庆", "西安"}
	s = append(s, s1...)
	fmt.Printf("s:%v  len(s):%v  cap(s):%v \n", s, len(s), cap(s)) // s:[上海 北京 广州 深圳 长沙 武汉 郑州 成都 重庆 西安]  len(s):10  cap(s):12

	// append()追加元素，原来的底层数组放不下的时候，Go底层会把底层数组换一个
	// 调用append()函数必须用变量接收返回值
	/*
		每个切片会指向一个底层数组，这个数组的容量够用就添加新增元素。当底层数组不能容纳新增的元素时，
		切片就会自动按照一定的策略进行“扩容”，此时该切片指向的底层数组就会更换。“扩容”操作往往发生在append()函数调用时，
		所以我们通常都需要用原变量接收append函数的返回值。
	*/

	// 通过var声明的零值切片可以在append()函数直接使用，无需初始化。
	var s2 []int
	s2 = append(s2, 2, 5, 8)

	// 切片的扩容策略
	// 可以通过查看$GOROOT/src/runtime/slice.go源码，其中扩容相关代码如下：
	/*
		newcap := old.cap
		doublecap := newcap + newcap
		if cap > doublecap {
			newcap = cap
		} else {
			if old.len < 1024 {
				newcap = doublecap
			} else {
				// Check 0 < newcap to detect overflow
				// and prevent an infinite loop.
				for 0 < newcap && newcap < cap {
					newcap += newcap / 4
				}
				// Set newcap to the requested cap when
				// the newcap calculation overflowed.
				if newcap <= 0 {
					newcap = cap
				}
			}
		}
	*/
	/*
		从上面的代码可以看出以下内容：
			首先判断，如果新申请容量（cap）大于2倍的旧容量（old.cap），最终容量（newcap）就是新申请的容量（cap）。
			否则判断，如果旧切片的长度小于1024，则最终容量(newcap)就是旧容量(old.cap)的两倍，即（newcap=doublecap），
			否则判断，如果旧切片长度大于等于1024，则最终容量（newcap）从旧容量（old.cap）开始循环增加原来的1/4，
		即（newcap=old.cap,for {newcap += newcap/4}）直到最终容量（newcap）大于等于新申请的容量(cap)，即（newcap >= cap）
			如果最终容量（cap）计算值溢出，则最终容量（cap）就是新申请容量（cap）。
		需要注意的是，切片扩容还会根据切片中元素的类型不同而做不同的处理，比如int和string类型的处理方式就不一样。
	*/

}
