package db

import "time"

type Id struct {
	Id       uint   `gorm:"primary_key;not null;AUTO_INCREMENT"`
}

type Time struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
