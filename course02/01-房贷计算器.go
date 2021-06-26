package main

import (
	"fmt"
	"math"
)

/*
	输入金额、年化利息（0.05代表5%）、还款年数
	输出月供，还款总额
*/
func main() {
	var money, pYear, year float64
	fmt.Print("客官请输入您要借贷的金额、年化利率、还款年数: ")
	fmt.Scanf("%f %f %f", &money, &pYear, &year)
	fmt.Println(money, pYear, year)

	p := pYear / 12

	/*套用公式进行计算*/
	monthPay := (p * math.Pow(1.0+p, pYear*12.0) * money) / (math.Pow(1.0+p, pYear*12.0) - 1.0)
	fmt.Println(monthPay)
}
