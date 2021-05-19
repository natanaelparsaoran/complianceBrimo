package module
import (
	"encoding/json"
	"time"
	"strconv"

	"github.com/addonrizky/complianceBrimo/library"
	"github.com/addonrizky/complianceBrimo/model"
	"github.com/addonrizky/complianceBrimo/constant"
)

func ComplyOpenAccFlexyRequest(
	accountNumber string, 
	productType string, 
	currency string,
	targetAmount string,
	month string,
	firstTransferDate string,
	bornDate string, //yyyy-mm-dd
	parameter string, 
	env string,
) model.Validation {
	var flexyParameter map[string]interface{}
	err := json.Unmarshal([]byte(parameter), &flexyParameter)

	if err != nil {
		return library.GetValidationResult(constant.RC_J0, constant.DESC_J0)
	}

	//validate source account, must be in array of parameter.allowed_saving_code [britama (50) & Britama bisnis (56)]
	if flexyParameter["allowed_saving_code"] == nil {
		return library.GetValidationResult(constant.RC_J1, constant.DESC_J1)
	}
	allowedSavingCodeInterface := flexyParameter["allowed_saving_code"].([]interface{})
	allowedSavingCode := make([]string, len(allowedSavingCodeInterface))
	for i, v := range allowedSavingCodeInterface {
		allowedSavingCode[i] = v.(string)
	}
	isSavingCodeAllowed := library.StringInSlice(accountNumber[12:14], allowedSavingCode)
	if(!isSavingCodeAllowed){
		//return "saving code not oke, it is neither britama nor britamabisnis"
		return library.GetValidationResult(constant.RC_F0, constant.DESC_F0)
	}

	//borndate->age must be >= 17yold
	if flexyParameter["borndate_min"] == nil {
		return library.GetValidationResult(constant.RC_J2, constant.DESC_J2)
	}
	allowedAge, _ := strconv.Atoi(flexyParameter["borndate_min"].(string))
	formatDate := "020106"
	formattedTime, err := time.Parse(formatDate, bornDate)
	if err != nil {
		//return "woy yg bener tanggalnyaz, format yg bener ini ...."
		return library.GetValidationResult(constant.RC_F1, constant.DESC_F1)
	}
	ageActual := library.Age(formattedTime, time.Now())
	if ageActual < allowedAge {
		//return "age notzz allowed"
		return library.GetValidationResult(constant.RC_F2, constant.DESC_F2)
	}

	//month must be >= 9month && <= 240
	if flexyParameter["month_min"] == nil {
		return library.GetValidationResult(constant.RC_J3, constant.DESC_J3)
	}
	if flexyParameter["month_max"] == nil {
		return library.GetValidationResult(constant.RC_J4, constant.DESC_J4)
	}
	allowedMonthMin, _ := strconv.Atoi(flexyParameter["month_min"].(string))
	allowedMonthMax, _ := strconv.Atoi(flexyParameter["month_max"].(string))
	monthInt, _ := strconv.Atoi(month)
	if monthInt < allowedMonthMin {
		//return "month not allowed, minimum tenor is " + strconv.Itoa(allowedMonthMin) + " month"
		return library.GetValidationResult(constant.RC_F3, constant.DESC_F3 + "" + strconv.Itoa(allowedMonthMin) + " bulan")
	}
	if monthInt > allowedMonthMax {
		//return "month not allowed, maximumt tenor is " + strconv.Itoa(allowedMonthMax) + " month"
		return library.GetValidationResult(constant.RC_F4, constant.DESC_F4 + "" + strconv.Itoa(allowedMonthMax) + " bulan")
	}

	//Target amount must be >= 500.000
	if flexyParameter["amount_min"] == nil {
		return library.GetValidationResult(constant.RC_J5, constant.DESC_J5)
	}
	if flexyParameter["amount_max"] == nil {
		return library.GetValidationResult(constant.RC_J6, constant.DESC_J6)
	}
	allowedAmountMin, _ := strconv.Atoi(flexyParameter["amount_min"].(string))
	allowedAmountMax, _ := strconv.Atoi(flexyParameter["amount_max"].(string))
	targetAmountInt, _ := strconv.Atoi(targetAmount)
	if targetAmountInt < allowedAmountMin {
		//return "target amount not allowed, minimum target is " + strconv.Itoa(allowedAmountMin)
		return library.GetValidationResult(constant.RC_F5, constant.DESC_F5)
	}
	if targetAmountInt > allowedAmountMax {
		//return "target amount not allowed, maximum target is " + strconv.Itoa(allowedAmountMax)
		return library.GetValidationResult(constant.RC_F6, constant.DESC_F6)
	}

	if env != "development" {
		//first transfer date must be > today && <= today+30
		if flexyParameter["first_transfer_date_min"] == nil {
			return library.GetValidationResult(constant.RC_J7, constant.DESC_J7)
		}
		if flexyParameter["first_transfer_date_max"] == nil {
			return library.GetValidationResult(constant.RC_J8, constant.DESC_J8)
		}
		allowedFirstTransferDateMin, _ := strconv.Atoi(flexyParameter["first_transfer_date_min"].(string))
		allowedFirstTransferDateMax, _ := strconv.Atoi(flexyParameter["first_transfer_date_max"].(string))
		formatDate = "020106"
		formattedTime, _ = time.Parse(formatDate, firstTransferDate)
		if err != nil {
			//return "woy yg bener tanggalnya, format yg bener ini ...."
			return library.GetValidationResult(constant.RC_F7, constant.DESC_F7)
		}
		rangeFirstTransferDate := library.Diffday(time.Now(), formattedTime)
		if rangeFirstTransferDate < allowedFirstTransferDateMin {
			//return "first transfer date not allowed, minimum first transfer date is H+" + strconv.Itoa(allowedFirstTransferDateMin)
			return library.GetValidationResult(constant.RC_F9, constant.DESC_F9 + "" + strconv.Itoa(allowedFirstTransferDateMin))
		}
		if rangeFirstTransferDate > allowedFirstTransferDateMax {
			//return "first transfer date not allowed, maximum first transfer date is H+" + strconv.Itoa(allowedFirstTransferDateMax)
			return library.GetValidationResult(constant.RC_F8, constant.DESC_F8 + "" + strconv.Itoa(allowedFirstTransferDateMax))
		}
	}

	//only BF that can be open in this phase
	//validate source account, must be in array of parameter.allowed_sc_code [britama (50) & Britama bisnis (56)]
	if flexyParameter["allowed_sc_code"] == nil {
		return library.GetValidationResult(constant.RC_J9, constant.DESC_J9)
	}
	allowedSCCodeInterface := flexyParameter["allowed_sc_code"].([]interface{})
	allowedSCCode := make([]string, len(allowedSCCodeInterface))
	for i, v := range allowedSCCodeInterface {
		allowedSCCode[i] = v.(string)
	}
	isSCCodeAllowed := library.StringInSlice(productType, allowedSCCode)
	if(!isSCCodeAllowed){
		//return "sc code not oke, it must be BF"
		return library.GetValidationResult(constant.RC_FA, constant.DESC_FA)
	}

	//only IDR that can be open in this phase
	if flexyParameter["allowed_currency"] == nil {
		return library.GetValidationResult(constant.RC_JA, constant.DESC_JA)
	}
	allowedCurrencyInterface := flexyParameter["allowed_currency"].([]interface{})
	allowedCurrency := make([]string, len(allowedCurrencyInterface))
	for i, v := range allowedCurrencyInterface {
		allowedCurrency[i] = v.(string)
	}
	isCurrencyAllowed := library.StringInSlice(currency, allowedCurrency)
	if(!isCurrencyAllowed){
		//return "currency not oke, it must be IDR"
		return library.GetValidationResult(constant.RC_FB, constant.DESC_FB)
	}

	return library.GetValidationResult(constant.RC_SUCCESS, constant.DESC_SUCCESS_FLEXY)
}