package module

import (
	"strings"

	"github.com/addonrizky/complianceBrimo/model"
	"github.com/addonrizky/complianceBrimo/rule"
)

func ComplyPassword(username string, usernameAlias string, password string, channel string) model.Validation {
	var validationObject model.Validation
	isUsernameAliasExist := false
	if usernameAlias != "" {
		isUsernameAliasExist = true
	}

	isPinPassword := false
	if channel == "EDC" || channel == "ATM" {
		isPinPassword = true
	}

	//password length must min 6 , max 6 if pin true
	if isPinPassword == true {
		if len(password) != 6 {
			validationObject = model.Validation{
				Code: "V6",
				Desc: "Panjang minimal & Maksimal pin adalah 6",
			}
			return validationObject
		}

		//password must contain only numeric
		if !rule.IsNumeric(password) {
			validationObject = model.Validation{
				Code: "VH",
				Desc: "pin harus angka",
			}
			return validationObject
		}

		switch password {
			case
			"123456",
			"654321",
			"112233",
			"111111",
			"222222",
			"333333",
			"444444",
			"555555",
			"666666",
			"777777",
			"888888",
			"999999",
			"000000":
			validationObject = model.Validation{
				Code: "VI",
				Desc: "kombinasi pin terlalu mudah",
			}
			return validationObject
		}

		
		
	}

	if isPinPassword == false {
		//password length must min 8
		if len(password) < 8 {
			validationObject = model.Validation{
				Code: "V6",
				Desc: "Panjang minimal password adalah 8",
			}
			return validationObject
		}

		//password length must max 12
		if len(password) > 12 {
			validationObject = model.Validation{
				Code: "V7",
				Desc: "Panjang maximal password adalah 12",
			}
			return validationObject
		}

		//password must not contain space
		if rule.IsSpaceExist(password) {
			validationObject = model.Validation{
				Code: "V8",
				Desc: "password tidak diperbolehkan mengandung karakter spasi",
			}
			return validationObject
		}

		//password must not contain only numeric
		if rule.IsNumeric(password) {
			validationObject = model.Validation{
				Code: "V9",
				Desc: "password tidak diperbolehkan hanya angka saja, harus mengandung minimal 1 karakter huruf kecil, 1 angka, dan 1 huruf besar.",
			}
			return validationObject
		}

		//password must not contain only alpha
		if rule.IsAlphaOnly(password) {
			validationObject = model.Validation{
				Code: "VA",
				Desc: "password tidak diperbolehkan hanya alfabet saja, harus mengandung minimal 1 karakter huruf kecil, 1 angka, dan 1 huruf besar.",
			}
			return validationObject
		}

		//password must contain alpha and numeric
		if !rule.IsAlphaNumeric(password) {
			validationObject = model.Validation{
				Code: "VB",
				Desc: "password harus mengandung alfa dan numeric",
			}
			return validationObject
		}

		//password must contain uppercase char
		if !rule.IsUppercaseLetterExist(password) {
			validationObject = model.Validation{
				Code: "VC",
				Desc: "password harus mengandung minimal 1 huruf besar",
			}
			return validationObject
		}

		//password must contain uppercase char
		if !rule.IsLowercaseLetterExist(password) {
			validationObject = model.Validation{
				Code: "VD",
				Desc: "password harus mengandung minimal 1 huruf kecil",
			}
			return validationObject
		}

		//first 4 char of password must not equal with first 4 char of username
		if strings.ToUpper(username[0:4]) == strings.ToUpper(password[0:4]) {
			validationObject = model.Validation{
				Code: "VF",
				Desc: "password tidak boleh mengandung unsur username",
			}
			return validationObject
		}

		//first 4 char of password must not equal with first 4 char of username alias
		if isUsernameAliasExist {
			if strings.ToUpper(usernameAlias[0:4]) == strings.ToUpper(password[0:4]) {
				validationObject = model.Validation{
					Code: "VG",
					Desc: "password tidak boleh mengandung unsur username",
				}
				return validationObject
			}
		}
	}

	validationObject = model.Validation{
		Code: "00",
		Desc: "format password valid",
	}
	return validationObject
}
