package main

import "fmt"

func main() {
	/* 定义局部变量 */
	var a int = 100
	var b int = 200
	var c int = 100
	var d int = 200

	fmt.Printf("交换前 a 的值为 : %d\n", a)
	fmt.Printf("交换前 b 的值为 : %d\n", b)

	/* 通过调用函数来交换值 */
	valueSwap(a, b)

	fmt.Printf("value_swap 交换后 a 的值 : %d\n", a)
	fmt.Printf("value_swap 交换后 b 的值 : %d\n", b)

	fmt.Printf("交换前 c 的值为 : %d\n", c)
	fmt.Printf("交换前 d 的值为 : %d\n", d)

	refSwap(&c, &d)
	fmt.Printf("pref_swap 交换后 c 的值为 : %d\n", c)
	fmt.Printf("pref_swap 交换后 d 的值为 : %d\n", d)
}

/* 定义相互交换值的函数 */
func valueSwap(x, y int) int {
	var temp int

	temp = x /* 保存 x 的值 */
	x = y    /* 将 y 值赋给 x */
	y = temp /* 将 temp 值赋给 y*/

	return temp
}

/* 定义交换值函数*/
func refSwap(x *int, y *int) {
	var temp int
	temp = *x /* 保持 x 地址上的值 */
	*x = *y   /* 将 y 值赋给 x */
	*y = temp /* 将 temp 值赋给 y */
}
