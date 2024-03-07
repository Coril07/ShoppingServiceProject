package user

import (
	"time"

	"gorm.io/gorm"
)

// 用户模型
type User struct {
	gorm.Model
	Username  string    `gorm:"type:varchar(30)"`
	Password  string    `gorm:"type:varchar(100)"`
	Password2 string    `gorm:"-"`
	Salt      string    `gorm:"type:varchar(100)"`
	Token     string    `gorm:"type:varchar(500)"`
	Gender    string    `gorm:"type:char(2)"`
	Birth     time.Time `gorm:"type:datetime"`
	Url       string    `gorm:"type:varchar(200)"`
	Email     string    `gorm:"type:varchar(200)"`
	Address   string    `gorm:"type:varchar(400)"`
	IsDeleted bool
	IsAdmin   bool
}

// 新建用户实例
func NewUser(username string, password string, password2 string, gender string, birth time.Time, url string, email string, address string) *User {

	return &User{
		Username:  username,
		Password:  password,
		Password2: password2,
		Gender:    gender,
		Birth:     birth,
		Url:       url,
		Email:     email,
		Address:   address,
		IsDeleted: false,
		IsAdmin:   false,
	}
}
