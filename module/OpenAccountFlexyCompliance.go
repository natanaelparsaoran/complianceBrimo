package module
import (
	"encoding/json"
	"time"
	"strconv"

	"github.com/addonrizky/complianceBrimo/library"
	"github.com/addonrizky/complianceBrimo/model"
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
	var validationObject model.Validation	
	err := json.Unmarshal([]byte(parameter), &flexyParameter)

	if err != nil {
		validationObject = model.Validation{
			Code: "J0",
			Desc: "invalid json for parameter",
		}
		return validationObject
	}

	//validate source account, must be in array of parameter.allowed_saving_code [britama (50) & Britama bisnis (56)]
	if flexyParameter["allowed_saving_code"] == nil {
		validationObject = model.Validation{
			Code: "J1",
			Desc: "invalid json for parameter -> allowed_saving_code broken",
		}
		return validationObject
	}
	allowedSavingCodeInterface := flexyParameter["allowed_saving_code"].([]interface{})
	allowedSavingCode := make([]string, len(allowedSavingCodeInterface))
	for i, v := range allowedSavingCodeInterface {
		allowedSavingCode[i] = v.(string)
	}
	isSavingCodeAllowed := library.StringInSlice(accountNumber[12:14], allowedSavingCode)
	if(!isSavingCodeAllowed){
		//return "saving code not oke, it is neither britama nor britamabisnis"
		validationObject = model.Validation{
			Code: "F0",
			Desc: "saving code tidak diperbolehkan. hanya rekening britama dan britama bisnis yg bisa digunakan utk buka rekening flexy",
		}
		return validationObject
	}

	//borndate->age must be >= 17yold
	if flexyParameter["borndate_min"] == nil {
		validationObject = model.Validation{
			Code: "J2",
			Desc: "invalid json for parameter -> borndate_min broken",
		}
		return validationObject
	}
	allowedAge, _ := strconv.Atoi(flexyParameter["borndate_min"].(string))
	formatDate := "020106"
	formattedTime, err := time.Parse(formatDate, bornDate)
	if err != nil {
		//return "woy yg bener tanggalnyaz, format yg bener ini ...."
		validationObject = model.Validation{
			Code: "F1",
			Desc: "format tanggal lahir salah, format yg benar ddmmyy",
		}
		return validationObject
	}
	ageActual := library.Age(formattedTime, time.Now())
	if ageActual < allowedAge {
		//return "age notzz allowed"
		validationObject = model.Validation{
			Code: "F2",
			Desc: "belum cukup umur. umur yg diperbolehkan 17 tahun ke atas",
		}
		return validationObject
	}

	//month must be >= 9month && <= 240
	if flexyParameter["month_min"] == nil {
		validationObject = model.Validation{
			Code: "J3",
			Desc: "invalid json for parameter -> month_min broken",
		}
		return validationObject
	}
	if flexyParameter["month_max"] == nil {
		validationObject = model.Validation{
			Code: "J4",
			Desc: "invalid json for parameter -> month_max broken",
		}
		return validationObject
	}
	allowedMonthMin, _ := strconv.Atoi(flexyParameter["month_min"].(string))
	allowedMonthMax, _ := strconv.Atoi(flexyParameter["month_max"].(string))
	monthInt, _ := strconv.Atoi(month)
	if monthInt < allowedMonthMin {
		//return "month not allowed, minimum tenor is " + strconv.Itoa(allowedMonthMin) + " month"
		validationObject = model.Validation{
			Code: "F3",
			Desc: "jangka waktu salah, minimum jangka waktu tabungan adalah " + strconv.Itoa(allowedMonthMin) + " bulan",
		}
		return validationObject
	}
	if monthInt > allowedMonthMax {
		//return "month not allowed, maximumt tenor is " + strconv.Itoa(allowedMonthMax) + " month"
		validationObject = model.Validation{
			Code: "F4",
			Desc: "jangka waktu salah, maksimum jangka waktu tabungan adalah " + strconv.Itoa(allowedMonthMax) + " bulan",
		}
		return validationObject
	}

	//Target amount must be >= 500.000
	if flexyParameter["amount_min"] == nil {
		validationObject = model.Validation{
			Code: "J5",
			Desc: "invalid json for parameter -> amount_min broken",
		}
		return validationObject
	}
	if flexyParameter["amount_max"] == nil {
		validationObject = model.Validation{
			Code: "J6",
			Desc: "invalid json for parameter -> amount_max broken",
		}
		return validationObject
	}
	allowedAmountMin, _ := strconv.Atoi(flexyParameter["amount_min"].(string))
	allowedAmountMax, _ := strconv.Atoi(flexyParameter["amount_max"].(string))
	targetAmountInt, _ := strconv.Atoi(targetAmount)
	if targetAmountInt < allowedAmountMin {
		//return "target amount not allowed, minimum target is " + strconv.Itoa(allowedAmountMin)
		validationObject = model.Validation{
			Code: "F5",
			Desc: "Target dana salah, minimum target dana adalah " + strconv.Itoa(allowedAmountMin),
		}
		return validationObject
	}
	if targetAmountInt > allowedAmountMax {
		//return "target amount not allowed, maximum target is " + strconv.Itoa(allowedAmountMax)
		validationObject = model.Validation{
			Code: "F6",
			Desc: "Target dana salah, maksimum target dana adalah " + strconv.Itoa(allowedAmountMax),
		}
		return validationObject
	}

	//first transfer date must be > today && <= today+30
	if flexyParameter["first_transfer_date_min"] == nil {
		validationObject = model.Validation{
			Code: "J7",
			Desc: "invalid json for parameter -> first_transfer_date_min broken",
		}
		return validationObject
	}
	if flexyParameter["first_transfer_date_max"] == nil {
		validationObject = model.Validation{
			Code: "J8",
			Desc: "invalid json for parameter -> first_transfer_date_max broken",
		}
		return validationObject
	}
	allowedFirstTransferDateMin, _ := strconv.Atoi(flexyParameter["first_transfer_date_min"].(string))
	allowedFirstTransferDateMax, _ := strconv.Atoi(flexyParameter["first_transfer_date_max"].(string))
	formatDate = "020106"
	formattedTime, _ = time.Parse(formatDate, firstTransferDate)
	if err != nil {
		//return "woy yg bener tanggalnya, format yg bener ini ...."
		validationObject = model.Validation{
			Code: "F7",
			Desc: "format tanggal transfer salah, format yg benar 'ddmmyy'",
		}
		return validationObject
	}
	rangeFirstTransferDate := library.Diffday(time.Now(), formattedTime)
	if rangeFirstTransferDate < allowedFirstTransferDateMin {
		//return "first transfer date not allowed, minimum first transfer date is H+" + strconv.Itoa(allowedFirstTransferDateMin)
		validationObject = model.Validation{
			Code: "F9",
			Desc: "tanggal transfer salah, minimum tanggal adalah H+" + strconv.Itoa(allowedFirstTransferDateMin),
		}
		return validationObject
	}
	if rangeFirstTransferDate > allowedFirstTransferDateMax {
		//return "first transfer date not allowed, maximum first transfer date is H+" + strconv.Itoa(allowedFirstTransferDateMax)
		validationObject = model.Validation{
			Code: "F8",
			Desc: "tanggal transfer salah, maksimum tanggal adalah H+" + strconv.Itoa(allowedFirstTransferDateMax),
		}
		return validationObject
	}

	//only BF that can be open in this phase
	//validate source account, must be in array of parameter.allowed_sc_code [britama (50) & Britama bisnis (56)]
	if flexyParameter["allowed_sc_code"] == nil {
		validationObject = model.Validation{
			Code: "J9",
			Desc: "invalid json for parameter -> allowed_sc_code broken",
		}
		return validationObject
	}
	allowedSCCodeInterface := flexyParameter["allowed_sc_code"].([]interface{})
	allowedSCCode := make([]string, len(allowedSCCodeInterface))
	for i, v := range allowedSCCodeInterface {
		allowedSCCode[i] = v.(string)
	}
	isSCCodeAllowed := library.StringInSlice(productType, allowedSCCode)
	if(!isSCCodeAllowed){
		//return "sc code not oke, it must be BF"
		validationObject = model.Validation{
			Code: "FA",
			Desc: "sc code tidak diperbolehkan. hanya BF bisa digunakan utk buka rekening flexy",
		}
		return validationObject
	}

	//only IDR that can be open in this phase
	if flexyParameter["allowed_currency"] == nil {
		validationObject = model.Validation{
			Code: "JA",
			Desc: "invalid json for parameter -> allowed_currency broken",
		}
		return validationObject
	}
	allowedCurrencyInterface := flexyParameter["allowed_currency"].([]interface{})
	allowedCurrency := make([]string, len(allowedCurrencyInterface))
	for i, v := range allowedCurrencyInterface {
		allowedCurrency[i] = v.(string)
	}
	isCurrencyAllowed := library.StringInSlice(currency, allowedCurrency)
	if(!isCurrencyAllowed){
		//return "currency not oke, it must be IDR"
		validationObject = model.Validation{
			Code: "FB",
			Desc: "currency tidak diperbolehkan. hanya IDR bisa digunakan utk buka rekening flexy",
		}
		return validationObject
	}

	validationObject = model.Validation{
		Code: "00",
		Desc: "seluruh parameter pembukaan rekening flexy valid",
	}
	return validationObject
}