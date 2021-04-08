package main

import (
	"fmt"
	"math"

	"github.com/shopspring/decimal"
)

func main() {
	// 因为float32的有效bit位只有23个，其它 的bit位用于指数和符号；当整数大于23bit能表达的范围时，float32的表示将出现误差
	var f float32 = 16777216
	fmt.Println(f == f+1) // "true"!

	const Avogadro = 6.02214129e23 // 阿伏伽德罗常数
	const Planck = 6.62606957e-34  // 普朗克常数

	fmt.Printf("Avogadro = %g, Planck = %e \n", Avogadro, Planck)

	for x := 0; x < 8; x++ {
		fmt.Printf("x = %d e^x = %8.3f \n", x, math.Exp(float64(x)))
	}

	fmt.Printf("%f \n", math.Pi)   // 默认保留6位小数点
	fmt.Printf("%.2f \n", math.Pi) // 保留2位小数点

	// 科学计数法表示浮点数
	var f1 float32 = 3.14e2 // 表示f1等于3.14*10的2次方
	fmt.Printf("%v--%T \n", f1, f1)
	var f2 float32 = 3.14e-2 // 表示f2等于3.14*10除以的2次方
	fmt.Printf("%v--%T \n", f2, f2)

	// 精度丢失
	var f3 float64 = 1129.6
	fmt.Println(f3 * 100) // 112959.99999999999

	m1 := 3.4
	m2 := 2.5
	fmt.Println(m1 - m2) // 0.8999999999999999
	/* 使用第三方包解决精度丢失问题 decimal */
	// 加法Add 减法Sub 乘法Mul 除法Div
	fmt.Println(decimal.NewFromFloat(m1).Sub(decimal.NewFromFloat(m2)))

	// decimal.DivisionPrecision = 2 // 保留两位小数，如有更多位，则进行四舍五入保留两位小数

	// 判断浮点数是否有效，不能是NAN或INF
	var val1 = 2.1
	if math.IsInf(val1, 0) || math.IsNaN(val1) {
		panic(fmt.Sprintf("Cannot create a Decimal from %v", val1)) // panic 抛出异常
	}

	fmt.Println(decimal.NewFromFloat(m1).IntPart()) // IntPart 返回小数的整数部分
	e, error := decimal.NewFromFloat(m1).Float64()
	if !error {
		fmt.Println(e)
	} else {
		fmt.Println(0)
	}

}
