package entities

type AddressInfo struct {
	addressline string
	city        string
	pincode     string
	province    string
	country     string
}

func NewAddressInfo(addressline string, city string, pincode string, province string, country string) AddressInfo {
	return AddressInfo{
		addressline: addressline,
		city:        city,
		pincode:     pincode,
		province:    province,
		country:     country,
	}
}
