package model

type User struct {
	//gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`

}

func (User) TableName() string {
	return "user"
}

