package main

import "fmt"

func main() {
	// 使用make()函数构造切片
	/*
		如果需要动态的创建一个切片，我们就需要使用内置的make()函数，格式如下：
			make([]T, size, cap)
		其中：
			T:切片的元素类型
			size:切片中元素的数量
			cap:切片的容量
	*/
	s := make([]int, 5, 10)
	fmt.Printf("s:%v  len(s):%v  cap(s):%v \n", s, len(s), cap(s)) // s:[0 0 0 0 0]  len(s):5  cap(s):10

	// 切片的本质
	/*
		切片的本质就是对底层数组的封装，它包含了三个信息：底层数组的指针、切片的长度（len）和切片的容量（cap）。
		切片就是一个框，框住了一块连续的内存
		切片属于引用类型，真正的数据都是保存在底层数组里的
	*/

	// 切片不能直接比较
	/*
		切片之间是不能比较的，我们不能使用==操作符来判断两个切片是否含有全部相等元素。
		切片唯一合法的比较操作是和nil比较。 一个nil值的切片并没有底层数组，一个nil值的切片的长度和容量都是0。
		但是我们不能说一个长度和容量都是0的切片一定是nil，例如下面的示例：
	*/
	var s1 []int         //len(s1)=0;cap(s1)=0;s1==nil
	s2 := []int{}        //len(s2)=0;cap(s2)=0;s2!=nil
	s3 := make([]int, 0) //len(s3)=0;cap(s3)=0;s3!=nil
	fmt.Printf("len(s1):%v  cap(s1):%v  s1==nil? %v \n", len(s1), cap(s1), s1 == nil)
	fmt.Printf("len(s2):%v  cap(s2):%v  s2==nil? %v \n", len(s1), cap(s1), s2 == nil)
	fmt.Printf("len(s3):%v  cap(s3):%v  s3==nil? %v \n", len(s3), cap(s3), s3 == nil)

	// 判断切片是否为空
	// 要检查切片是否为空，请始终使用len(s) == 0来判断，而不应该使用s == nil来判断。

	// 切片的赋值拷贝
	// 下面的代码中演示了拷贝前后两个变量共享底层数组，对一个切片的修改会影响另一个切片的内容，这点需要特别注意。
	s4 := make([]int, 3)
	s5 := s4            // s4和s5都指向同一个底层数组
	fmt.Println(s4, s5) // [0 0 0] [0 0 0]
	s4[0] = 123
	fmt.Println(s4, s5) // [123 0 0] [123 0 0]

	// 切片遍历
	// 切片的遍历方式和数组是一致的，支持索引遍历和for range遍历
	s6 := []int{1, 3, 5}
	for i := 0; i < len(s6); i++ {
		fmt.Println(s6[i])
	}

	for _, v := range s6 {
		fmt.Println(v)
	}

}
