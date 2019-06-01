package model

type Address struct {
	Model
	State string `json:"state"`
	City string `json:"city"`
	PinCode int `json:"pinCode"`
	Country string `json:"country"`
	Latitude float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
	AddressType priorityType `json:"address_type"`
	UserId uint `json:"user_id"`
}


type JSONAddress struct {
	Id uint `json:"id"`
	State string `json:"state"`
	City string `json:"city"`
	PinCode int `json:"pinCode"`
	Country string `json:"country"`
	Latitude float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
	AddressType priorityType `json:"address_type"`
	UserId uint `json:"user_id"`
}

func NewJSONAddress(address Address) JSONAddress{
	return JSONAddress{
		Id:address.Id,
		State:address.State,
		City:address.City,
		PinCode:address.PinCode,
		Country:address.Country,
		Latitude:address.Latitude,
		Longitude:address.Longitude,
		AddressType:address.AddressType,
		UserId:address.UserId,
	}
}