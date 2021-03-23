package main

import (
	"first/pkg1"
	"first/pkg2"
	"fmt"
	"math"
	"strconv"
)

func main() {
	//t1() //基础语法
	//t2() //基础语法
	//t3() //基础语法
	//t4()
	//t5() //换行转移
	//t6() //
	//t7() //字符串转换，格式化浮点，格式化输出

	//简单计算和多返回
	//a, b := t8(5, 3)
	//println(a, b)

	//值传递，非引用传递
	//str := "23"
	//t9(str)
	//println(str)

	//引用传递加*号,调用方&
	//str := "1122"
	//t10(&str)
	//println(str)

	t11() //其他模块引用

}

//其他模块引用
func t11() {
	cc := pkg1.Sum1(1, 3)
	println(cc)

	//同包下面可以直接使用
	dd := Sum(16, 15)
	println(dd)

	//
	ee := pkg2.Div(20, 3)
	var ees = strconv.FormatFloat(float64(ee), 'f', 5, 32)
	println(ees)
}

//引用传递
func t10(str *string) {
	//str = "aaa"//报错
	*str = "bbb"
}

//值传递
func t9(str string) {
	str = "a"
}

//计算
func t8(a int, b int) (int, int) {
	return a + b, b
}

//强转，格式化输出
func t7() {
	//https://docs.microsoft.com/zh-cn/learn/modules/go-variables-functions-packages/2-data-types
	var integer16 int16 = 127
	var integer32 int32 = 32767
	println(int32(integer16) + integer32)

	// 'b' (-ddddp±ddd，二进制指数)
	// 'e' (-d.dddde±dd，十进制指数)
	// 'E' (-d.ddddE±dd，十进制指数)
	// 'f' (-ddd.dddd，没有指数)
	// 'g' ('e':大指数，'f':其它情况)
	// 'G' ('E':大指数，'f':其它情况)

	s := strconv.Itoa(-42)
	i, err := strconv.ParseFloat("-42.0122", 64)
	var str = strconv.FormatFloat(i, 'f', -1, 32)
	println(i, s)
	println(err)
	println(str)
}

//默认值
func t6() {
	var defaultInt int
	var defaultFloat32 float32
	var defaultFloat64 float64
	var defaultBool bool
	var defaultString string
	println(defaultInt, defaultBool, defaultFloat32, defaultFloat64, defaultString)
}
func t5() {
	//	\n：新行
	//\r：回车符
	//\t：选项卡
	//\'：单引号
	//\"：双引号
	//\\：反斜杠

	var firstName = "John"
	lastName := "Doe"
	println(firstName, lastName)

	fullName := "John Doe \t(alias \"Foo\")\n"
	println(fullName)
}

func t4() {
	println(math.MaxFloat32, math.MaxFloat64)
}

func t3() {
	var integer8 int8 = 127
	var integer16 int16 = 32767
	var integer32 int32 = 2147483647
	var integer64 int64 = 9223372036854775807
	println(integer8, integer16, integer32, integer64)

	var integer int8 = -10
	println(integer)
}

const HTTPStatusOK = 200

//定义变量
func t2() {
	var firstName string
	println("第一个：" + firstName)

	var (
		firstName1 string = "John"
		lastName2  string = "Doe"
		age3       int    = 32
	)
	println(firstName1, lastName2, age3)

	firstName4 := "xxx"
	println(firstName4)

	println(HTTPStatusOK)
}

//第一个
func t1() {
	fmt.Println("hello world")
}
