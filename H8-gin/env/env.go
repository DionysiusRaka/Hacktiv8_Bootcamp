package env

import "gorm.io/gorm"

var (
	Host   = "localhost"
	User   = "postgres"
	Pass   = "password"
	Port   = "5432"
	Dbname = "golang-swagger"
	Db     *gorm.DB
	Err    error
)
