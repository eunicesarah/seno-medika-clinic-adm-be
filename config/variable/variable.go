package variable

import (
	"os"
)

var DbHost = os.Getenv("POSTGRES_HOST")
var DbUser = os.Getenv("POSTGRES_USER")
var DBName = os.Getenv("POSTGRES_DB")
var DBPass = os.Getenv("POSTGRES_PASSWORD")
var JWTSecret = os.Getenv("JWT_SECRET")
