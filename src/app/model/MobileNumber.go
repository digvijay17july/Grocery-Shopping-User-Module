package model

type MobileNumber struct{
	Model
	Number string `json:"mobileNo"`
	NumberType priorityType `json:"number_type"`
	UserId uint `json:"user_id"`
}

type JSONMobileNumber struct{
	Id uint  `json:"id"`
	Number string `json:"mobileNo"`
	NumberType priorityType `json:"number_type"`
	UserId uint `json:"user_id"`
}

func NewJSONMobileNumber(mobileNumber MobileNumber) JSONMobileNumber{
	return JSONMobileNumber{
		Id:mobileNumber.Id,
		Number:mobileNumber.Number,
		NumberType:mobileNumber.NumberType,
		UserId:mobileNumber.UserId,
	}
}