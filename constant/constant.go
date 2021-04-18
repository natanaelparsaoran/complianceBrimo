package constant 

const (
	RC_SUCCESS = "00"
	
	/*
	 * britama flexy
	 */
	DESC_SUCCESS_FLEXY = "seluruh parameter pembukaan rekening flexy valid"	
	//----
	RC_FB = "FB"	
	DESC_FB = "currency tidak diperbolehkan. hanya IDR bisa digunakan utk buka rekening flexy"
	//----
	RC_JA = "JA"	
	DESC_JA = "invalid json for parameter -> allowed_currency broken"
	//----
	RC_FA = "FA"	
	DESC_FA = "sc code tidak diperbolehkan. hanya BF bisa digunakan utk buka rekening flexy"
	//----
	RC_J9 = "J9"
	DESC_J9 = "sc code tidak diperbolehkan. hanya BF bisa digunakan utk buka rekening flexy"
	//----
	RC_F8 = "F8"
	DESC_F8 = "tanggal transfer salah, maksimum tanggal adalah H+"
	//----
	RC_F9 = "F9"
	DESC_F9 = "tanggal transfer salah, minimum tanggal adalah H+"
	//----
	RC_F7 = "F7"
	DESC_F7 = "format tanggal transfer salah, format yg benar 'ddmmyy'"
	//----
	RC_J8 = "J8"
	DESC_J8 = "invalid json for parameter -> first_transfer_date_max broken"
	//----
	RC_J7 = "J7"
	DESC_J7 = "invalid json for parameter -> first_transfer_date_min broken"
	//----
	RC_F6 = "F6"
	DESC_F6 = "Target dana salah, maksimum target dana adalah "// + strconv.Itoa(allowedAmountMax),
	//----
	RC_F5 = "F5"
	DESC_F5 = "Target dana salah, minimum target dana adalah "// + strconv.Itoa(allowedAmountMin),
	//----
	RC_J6 = "J6"
	DESC_J6 = "invalid json for parameter -> amount_max broken"
	//----
	RC_J5 = "J5"
	DESC_J5 = "invalid json for parameter -> amount_min broken"
	//----
	RC_F4 = "F4"
	DESC_F4 = "jangka waktu salah, maksimum jangka waktu tabungan adalah "// + strconv.Itoa(allowedMonthMax) + " bulan"
	//----
	RC_F3 = "F3"
	DESC_F3 = "jangka waktu salah, minimum jangka waktu tabungan adalah "// + strconv.Itoa(allowedMonthMin) + " bulan"
	//----
	RC_J4 = "J4"
	DESC_J4 = "invalid json for parameter -> month_max broken"
	//----
	RC_J3 = "J3"
	DESC_J3 = "invalid json for parameter -> month_min broken"
	//----
	RC_F2 = "F2"
	DESC_F2 = "belum cukup umur. umur yg diperbolehkan 17 tahun ke atas"
	//----
	RC_F1 = "F1"
	DESC_F1 = "format tanggal lahir salah, format yg benar ddmmyy"
	//----
	RC_J2 = "J2"
	DESC_J2 = "invalid json for parameter -> borndate_min broken"
	//----
	RC_F0 = "F0"
	DESC_F0 = "saving RC_ tidak diperbolehkan. hanya rekening britama dan britama bisnis yg bisa digunakan utk buka rekening flexy"
	//----
	RC_J1 = "J1"
	DESC_J1 = "invalid json for parameter -> allowed_saving_code broken"
	//----
	RC_J0 = "J0"
	DESC_J0 = "invalid json for parameter"
)