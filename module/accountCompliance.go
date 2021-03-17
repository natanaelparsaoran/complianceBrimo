package module

func GetProductTypeByAccnum(accountNumber string) string {
	productType := ""

	//if accountNumber[13:14]
	switch accountNumber[12:14] {
	case "30", "31", "32":
		productType = "GiroBRI"
	case "50":
		productType = "BritAma"
	case "53":
		productType = "Simpedes"
	case "51":
		productType = "Tab.Haji"
	case "56":
		productType = "BritAmaBisnis"
	default:
		productType = "BritAma"
	}

	return productType
}
