package env

import "os"

var DbHost = os.Getenv("MYSQL_SERVER")
var DbPort = os.Getenv("MYSQL_PORT")
var DbDatabase = os.Getenv("MYSQL_DATABASE")
var DbUsername = os.Getenv("MYSQL_USER")
var DbPassword = os.Getenv("MYSQL_PASSWORD")
