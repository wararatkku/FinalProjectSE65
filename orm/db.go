package orm

import (
	// เชื่อมต่อ mysql
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm" // framwork ต่อกับ database ภาษา GO
)

var Db *gorm.DB
var err error

func InitDB() {
	// ติดต่อ mysql
	dsn := os.Getenv("MYSQL_DNS")
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("fail to connect database")
	}

	// Migrate the schema
	Db.AutoMigrate(&User{})    // เอาโครงสร้าง structure User ทั้งหมดลง mysql
	Db.AutoMigrate(&Car{})     // เอาโครงสร้าง structure Car ทั้งหมดลง mysql
	Db.AutoMigrate(&Booking{}) // เอาโครงสร้าง structure Booking ทั้งหมดลง mysql
}
