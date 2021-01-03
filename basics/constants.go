package main

import "fmt"

/**
Use CamelCase for contatnts naming
*/

func getPrice(plan int, unitPrice float32) float32 {

	return float32(plan) * unitPrice

}

func main() {
	const PI = 3.14
	const (
		PlanFree = iota
		PlanBasic
		PlanAdvanced
		PlanPremiume
	)
	fmt.Println("Value of pi is ", PI)
	fmt.Printf("Plan free %v, plan basic %v, plan advanced %v and plane premiume %v", PlanFree, PlanBasic, PlanAdvanced, PlanPremiume)
	const UnitPrice float32 = 90.90
	const PriceBasic = PlanBasic * UnitPrice

	fmt.Printf("\nBasic plane price %v and type %T", PriceBasic, PriceBasic)

	var PriceAvanced = getPrice(PlanAdvanced, UnitPrice)

	fmt.Printf("\nAdvanved plan price : %v", PriceAvanced)

	const BrandName = "Golang"

	fmt.Printf("\nBrand name is %v ", BrandName)
}
