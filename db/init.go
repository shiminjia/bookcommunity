package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

//global DB object
var DB *gorm.DB

func Init() {
	DB = DBopen()
}

func DBopen() *gorm.DB {
	dbconfig := "root:11111111@/book_community?charset=utf8&parseTime=True&loc=Local"

	db, err := gorm.Open("mysql", dbconfig)
	if err != nil {
		log.Println("DB access ERROR.")
		panic("DB access ERROR.")
	}
	log.Println("DB access OK.")

	//create db if not exists
	CreateTable(db)

	return db
}

func CreateTable(db *gorm.DB){

	//check users table exists or not.
	//if users table does not exist, create it by User{}
	bool :=db.HasTable(&User{})
	if !bool {
		db.Table("users").CreateTable(&User{})
		log.Println("Table users created.")
	}

}

func Close(DB *gorm.DB) {
	DB.Close()
}