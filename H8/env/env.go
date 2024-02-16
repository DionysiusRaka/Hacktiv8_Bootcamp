package env

import "gorm.io/gorm"

var (
	Host   = "localhost"
	User   = "postgres"
	Pass   = "password"
	Port   = "5432"
	Dbname = "learning-gorm"
	Db     *gorm.DB
	Err    error
)
