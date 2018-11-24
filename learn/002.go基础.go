package main // 代码包声明语句

// 代码包导入语句
import (
	"fmt" // 导入代码包fmt。
)

// main函数
func main() { // 代码块由“{”和“}”包裹。

	// 变量声明和赋值语句，由关键字var、变量名num、变量类型uint64、特殊标记=，以及值10组成。
	var num uint64 = 65535

	// 短变量声明语句，由变量名size、特殊标记:=，以及值（需要你来填写）组成。
	size := (8)

	// 打印函数调用语句。在这里用于描述一个uint64类型的变量所需占用的比特数。
	// 这里用到了字符串的格式化函数。
	fmt.Printf("类型为 uint64 的整数 %d 需占用的存储空间为 %d 个字节。\n", num, size)
	fmt.Println("___________fun main2___________")
	main2()
	integerTrs()
	floatFun()
	complexFun()
	runeFun()
	stringFun()
}

//2.进制转换
func integerTrs() {
	// 声明一个整数类型变量并赋值
	var num1 int = -0x1000

	// 这里用到了字符串格式化函数。其中，%X用于以16进制显示整数类型值，%d用于以10进制显示整数类型值。
	fmt.Printf("16进制数 %X 表示的是 %d。\n", num1, -4096)
}

//3.申明
func main2() {
	var (
		num1 int
		num2 int
		num3 int
	)
	num1, num2, num3 = 1, 2, 3
	// 打印函数调用语句。用于打印上述三个变量的值。
	fmt.Println(num1, num2, num3)
}

//4.浮点数
func floatFun() {
	var fnum = 3.14e10
	fmt.Printf("浮点数是:%f, e= %E \n", fnum, fnum)
}

//5.复数
func complexFun() {
	var num3 = 3.7E+1 + 5.98E-2i
	// 这里用到了字符串格式化函数。其中，%E用于以带指数部分的表示法显示浮点数类型值，%f用于以通常的方法显示浮点数类型值。
	fmt.Printf("浮点数 %E 表示的是 %f。\n", num3, (num3))
}

//6.rune类型
func runeFun() {
	// 声明一个rune类型变量并赋值
	var char1 rune = '赞'
	fmt.Printf("字符 '%c' 的Unicode代码点是 %s。\n", char1, ("U+8D5E"))
}

//7.string
func stringFun() {
	var str1 string = "\\\""
	fmt.Printf("用解释型字符串表示法表示的 %q 所代表的是 %s。\n", str1, `\"`)
}
