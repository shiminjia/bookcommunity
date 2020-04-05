package config

import "os"

var MODE = os.Getenv("MODE")
var PORT = os.Getenv("PORT")

// about JWT
var JWT_SECRET = os.Getenv("JWT_SECRET")
var ISSUER = os.Getenv("ISSUER")
var EXPIRATION_TIME_LOGIN = os.Getenv("ONE_DAY")
var EXPIRATION_TIME_EMAIL = os.Getenv("TEN_MINS")

//about encrypt
var SALT = os.Getenv("SALT")

//about db
var DB_USER = os.Getenv("DB_USER")
var DB_PASSWORD = os.Getenv("DB_PASSWORD")
var DB_DATABASE = os.Getenv("DB_DATABASE")

// about redirect
var INDEX = os.Getenv("INDEX")