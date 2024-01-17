package usermodel

type ReqUpdateUser struct {
	Name string `json:"name" gorm:"column:name"`
}

func (ReqUpdateUser) TableName() string {
	return User{}.TableName()
}
