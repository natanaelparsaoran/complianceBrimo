package main

import (
	"fmt"

	"github.com/addonrizky/complianceBrimo/module"
	"github.com/addonrizky/complianceBrimo/rule"
)

func main() {
	isSpaceExist := rule.IsSpaceExist("fewoigiow")
	isNumeric := rule.IsNumeric("8888")
	isAlphaOnly := rule.IsAlphaOnly("iijioio ")
	isAlphaNumeric := rule.IsAlphaNumeric("fwjeiofjweoi")
	wkwkwk := module.ComplyUsername("wegweg534612")
	ckckck := module.ComplyPassword("adonit2504", "ramadhan2568", "251331", "EDC")
	oke := module.GetProductTypeByAccnum("020601087063504")
	
	fmt.Println(isSpaceExist)
	fmt.Println(isNumeric)
	fmt.Println(isAlphaOnly)
	fmt.Println(isAlphaNumeric)
	fmt.Println(wkwkwk)
	fmt.Println(ckckck)
	fmt.Println(oke)
	check := module.ComplyOpenAccFlexyRequest(
		"020601087063504", 
		"BF", 
		"IR",
		"1000000",
		"239",
		"190421",
		"250489", 
		`{"allowed_saving_code":["50","56"],"borndate_min":"17","month_min":"9","month_max":"240","amount_min":"500000","amount_max":"1000000000","first_transfer_date_min":"1","first_transfer_date_max":"30","allowed_sc_code":["BF"],"allowed_currency":["IDR"]}`, 
		"dev",
	)

	fmt.Println(check)
}
