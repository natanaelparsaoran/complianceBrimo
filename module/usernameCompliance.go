package module

import (
	"github.com/addonrizky/complianceBrimo/model"
	"github.com/addonrizky/complianceBrimo/rule"
)

func ComplyUsername(username string) model.Validation {
	var validationObject model.Validation

	//username length must min10
	if len(username) < 12 {
		validationObject = model.Validation{
			Code: "V0",
			Desc: "Panjang minimal username adalah 12",
		}
		return validationObject
	}

	//username length must max23
	if len(username) > 23 {
		validationObject = model.Validation{
			Code: "V1",
			Desc: "Panjang maximal username adalah 23",
		}
		return validationObject
	}

	//username must not contain space
	if rule.IsSpaceExist(username) {
		validationObject = model.Validation{
			Code: "V2",
			Desc: "username tidak diperbolehkan mengandung karakter spasi",
		}
		return validationObject
	}

	//username must not contain only numeric
	if rule.IsNumeric(username) {
		validationObject = model.Validation{
			Code: "V3",
			Desc: "username tidak diperbolehkan hanya angka saja",
		}
		return validationObject
	}

	//username must not contain only alpha
	if rule.IsAlphaOnly(username) {
		validationObject = model.Validation{
			Code: "V4",
			Desc: "username tidak diperbolehkan hanya alfabet saja",
		}
		return validationObject
	}

	//username must contain alpha and numeric
	if !rule.IsAlphaNumeric(username) {
		validationObject = model.Validation{
			Code: "V5",
			Desc: "username harus/hanya mengandung alfa dan numeric",
		}
		return validationObject
	}

	validationObject = model.Validation{
		Code: "00",
		Desc: "format username valid",
	}
	return validationObject
}
