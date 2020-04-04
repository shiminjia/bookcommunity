package db

import (
	"time"
)

type User struct {
	Id
	Email        string     `gorm:"type:varchar(100);unique;not null;unique_index"`
	Password     string     `gorm:"type:varchar(100);not null"`
	Username     string     `gorm:"type:varchar(100);unique;not null;unique_index"`
	Icon_url     string     `gorm:"type:varchar(100)"`
	First_name   string     `gorm:"type:varchar(100)"`
	Middle_name  string     `gorm:"type:varchar(100)"`
	Last_name    string     `gorm:"type:varchar(100)"`
	Gender       int64      `gorm:"type:bigint(2);not null;default:0"`
	Birthday     *time.Time `gorm:"type:datetime;"`
	Location     string     `gorm:"type:varchar(100)"`
	Carerr       string     `gorm:"type:varchar(100)"`
	Interest     string     `gorm:"type:varchar(100)"`
	Introduction string     `gorm:"type:varchar(1000)"`
	Level        int64      `gorm:"type:bigint(2);not null;default:10"`
	Status       int64      `gorm:"type:bigint(2);not null;default:0"`
	Time
}

// set User table name to users
func (u *User) TableName() string {
	return "users"
}

// Create a new user.
func (u *User) Create() (*User,error) {
	err := DB.Create(&u).Error
	if err != nil {
		return nil, err
	}
	return u, nil
}

//check email or username is unique or not.
func (u *User) Verify() (*User, error) {
	user := &User{}
	d := DB.Where("email = ? or username = ?", u.Email, u.Username).Find(user)
	return user, d.Error
}

//Get user info by id
func (u *User) GetUserById() (*User, error) {
	user := &User{}
	d := DB.Where("id = ?", u.Id.Id).Find(user)
	return user, d.Error
}

func (u *User) UpdateStatus() (*User, error) {
	d := DB.Model(&u).Where("id = ?", u.Id.Id).Update("status", u.Status)
	return u, d.Error
}

// Create creates a new user account.

// Update an user.

// Get an user info.
