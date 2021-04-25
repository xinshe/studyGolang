package main

import "fmt"

// 切片
/*
切片（Slice）是一个拥有相同类型元素的可变长度的序列。它是基于数组类型做的一层封装。它非常灵活，支持自动扩容。
切片是一个引用类型，它的内部结构包含地址、长度和容量。切片一般用于快速地操作一块数据集合。
*/

func main() {
	// 切片的声明
	var a []int
	var b []string
	fmt.Println(a, b)     // [] []
	fmt.Println(a == nil) // true
	fmt.Println(b == nil) // true

	// 切片的初始化
	var c = []int{1, 2}
	var d = []string{"Andy", "Sunny"}
	fmt.Println(c)        // [1 2]
	fmt.Println(d)        // [Andy Sunny]
	fmt.Println(c == nil) // false
	fmt.Println(d == nil) // false

	var e = []bool{true, false}
	var f = []bool{true, false}
	fmt.Println(e, f) // [true false] [true false]
	//fmt.Println(e == f)	// 切片是引用类型，不支持直接比较，只能和nil比较

	// 切片的长度和容量
	// 切片拥有自己的长度和容量，我们可以通过使用内置的len()函数求长度，使用内置的cap()函数求切片的容量。
	fmt.Println(len(c)) // 2
	fmt.Println(cap(c)) // 2

	fmt.Println(len(a)) // 0
	fmt.Println(cap(a)) // 0

	// 切片表达式
	/*
		切片表达式从字符串、数组、切片、指向数组或切片的指针构造子字符串或切片。
		它有两种变体：一种指定low和high两个索引界限值的简单的形式，另一种是除了low和high索引界限值外还指定容量的完整的形式。
	*/

	// 通过数组得到切片
	/*
		切片的底层就是一个数组，所以我们可以基于数组通过切片表达式得到切片。切片表达式中的low和high表示一个索引范围（左包含，右不包含），
		也就是下面代码中从数组arr中选出1<=索引值<4的元素组成切片s，
		得到的切片长度 len=high-low，容量等于得到的切片的底层数组的容量 cap=原数组长度-low。
	*/
	arr := [6]int{1, 2, 3, 4, 5, 6}
	s := arr[1:4]
	fmt.Printf("s:%v  len(s):%v  cap(s):%v \n", s, len(s), cap(s)) // s:[2 3 4]  len(s):3  cap(s):5
	s1 := arr[3:5]
	fmt.Printf("s1:%v  len(s1):%v  cap(s1):%v \n", s1, len(s1), cap(s1)) // s1:[4 5]  len(s1):2  cap(s1):3

	// 为了方便起见，可以省略切片表达式中的任何索引。省略了low则默认为0；省略了high则默认为切片操作数的长度:
	s2 := arr[:4]
	fmt.Printf("s2:%v  len(s2):%v  cap(s2):%v \n", s2, len(s2), cap(s2)) // s2:[1 2 3 4]  len(s2):4  cap(s2):6
	s3 := arr[2:]
	fmt.Printf("s3:%v  len(s3):%v  cap(s3):%v \n", s3, len(s3), cap(s3)) // s3:[3 4 5 6]  len(s3):4  cap(s3):4
	s4 := arr[:]
	fmt.Printf("s4:%v  len(s4):%v  cap(s4):%v \n", s4, len(s4), cap(s4)) // s4:[1 2 3 4 5 6]  len(s4):6  cap(s4):6

	// 切片再切片：通过切片得到切片
	/*
		对于数组或字符串，如果0 <= low <= high <= len(a)，则索引合法，否则就会索引越界（out of range）。
		对切片再执行切片表达式时（切片再切片），high的上限边界是切片的容量cap(a)，而不是长度。
		常量索引必须是非负的，并且可以用int类型的值表示；对于数组或常量字符串，常量索引也必须在有效范围内。
		如果low和high两个指标都是常数，它们必须满足low <= high。如果索引在运行时超出范围，就会发生运行时panic。
	*/
	ss := s[2:4]
	fmt.Printf("ss:%v  len(ss):%v  cap(ss):%v \n", ss, len(ss), cap(ss)) // ss:[4 5]  len(ss):2  cap(ss):3

	// 切片是引用类型，都指向了底层数组
	fmt.Println(s, ss) // [2 3 4] [4 5]
	arr[3] = 999
	fmt.Println(s, ss) // [2 3 999] [999 5]

	// 完整切片表达式
	/*
		对于数组，指向数组的指针，或切片a(注意不能是字符串)支持完整切片表达式：a[low : high : max]
		上面的代码会构造与简单切片表达式a[low: high]相同类型、相同长度和元素的切片。
		另外，它会将得到的结果切片的容量设置为max-low。在完整切片表达式中只有第一个索引值（low）可以省略；它默认为0。
		完整切片表达式需要满足的条件是0 <= low <= high <= max <= cap(a)，其他条件和简单切片表达式相同。
	*/
	arr1 := [5]int{1, 2, 3, 4, 5}
	s5 := arr1[1:3:5]
	fmt.Printf("s5:%v  len(s5):%v  cap(s5):%v \n", s5, len(s5), cap(s5)) // s5:[2 3]  len(s5):2  cap(s5):4
}
