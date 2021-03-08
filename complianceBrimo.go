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
	ckckck := module.ComplyPassword("adonit2504", "ramadhan2568", "123456ewheHfh", "OOAO")
	fmt.Println(isSpaceExist)
	fmt.Println(isNumeric)
	fmt.Println(isAlphaOnly)
	fmt.Println(isAlphaNumeric)
	fmt.Println(wkwkwk)
	fmt.Println(ckckck)
}
