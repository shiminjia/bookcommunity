package config

import "os"

var MODE = os.Getenv("MODE")
var ADDR = os.Getenv("ADDR")

// about JWT
var JWT_SECRET = os.Getenv("JWT_SECRET")
var ISSUER = os.Getenv("ISSUER")
var JWT_EFFECTIVE_TIME = os.Getenv("JWT_EFFECTIVE_TIME")
//var JWT_EFFECTIVE_TIME = "30"
var SALT = os.Getenv("SALT")
//var SALT = "SALT"
