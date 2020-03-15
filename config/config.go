package config

import "os"

var MODE = os.Getenv("MODE")
var ADDR = os.Getenv("ADDR")
var JWT_SECRET = os.Getenv("JWT_SECRET")