package entities

type Venue struct {
	addressInfo AddressInfo
	stadiumName string
	stadiumId   uint16
}

func NewVenue(addressInfo AddressInfo, stadiumName string, stadiumId uint16) Venue {
	return Venue{
		addressInfo: addressInfo,
		stadiumName: stadiumName,
		stadiumId:   stadiumId,
	}
}
