package models

import (
	"crud_demo/app/helpers"
	"crud_demo/config"
	"time"
)

//var DB *gorm.DB

type User struct {
	//gorm.Model
	Id        int       `gorm:"primary_key;auto_increment" json:"id" `
	Name      string    `gorm:"not null;size:60" validate:"required" json:"name"`
	Username  string    `gorm:"not null;size:30" validate:"required" json:"username"`
	Password  string    `gorm:"not null;size:100" validate:"required" json:"password"`
	Status    int       `validate:"required" json:"status" `
	Role      int       `validate:"required" json:"role" `
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func AuthLogin(username string, password string) *User {
	user := new(User)
	res := config.DB.Debug().Where("username = ?", username).First(&user)

	if res.Error == nil {
		err := helpers.VerifyPassword(user.Password, password)
		if err != nil {
			return user
		}
		return nil
	}

	return user
}

func UserList() []User {
	var user []User
	res := config.DB.Find(&user)
	if res.Error == nil {
		return user
	}

	return nil
}

func NewUserStore(user *User) bool {
	res := config.DB.Debug().Create(&user)
	if res != nil {
		return true
	}

	return false
}

func UserShowById(id int) *User {
	user := new(User)
	res := config.DB.First(&user, id)

	if res.Error == nil {
		return user
	}

	return nil
}

func UserDelete(id int) bool {
	user := new(User)

	res := config.DB.Where("id = ?", id).Delete(&user)
	if res.Error == nil {
		return true
	}

	return false
}
