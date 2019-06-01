package model



type User struct {
	Model
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Status bool   `json:"status"`
	Type   string `json:"user"`
	Password string `json:"password"`
	MobileNumbers []MobileNumber `gorm:"foreignkey:UserId"  json:"mobileNumbers"`
	Addresses [] Address `gorm:"foreignkey:UserId" json:"addresses"`
	Email string `json:"email"`
	Username string `json:"username"`
}


type JSONUser struct {
Id     uint   `json:"id"`
Name   string `json:"name"`
Age    int    `json:"age"`
Status bool   `json:"status"`
Type   string `json:"user"`
	MobileNumbers []MobileNumber `json:mobileNumbers`
	Addresses [] Address `json:"addresses"`
	Email string `json:"email"`
	Username string `json:"username"`
}

func (u *User) Disable(){
	u.Status=false
}
func (u *User) Enable(){
	u.Status=false
}


func NewJSONUser(user User) JSONUser{
	return JSONUser{
      user.Id,
		user.Name,
		user.Age,
		user.Status,
		user.Type,
		user.MobileNumbers,
		user.Addresses,
		user.Email,
		user.Username,
	}
}