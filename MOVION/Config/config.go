package config

import (
	"log"
	models "movion/Models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Jwtkey = []byte("secretkey")

var DB *gorm.DB

func InitDB(){
	dsn := "root:jishnu2004@tcp(127.0.0.1:3306)/movion?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("database connection failed")
	}

	err = autoMigrate()
	// err = DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Migration failed:", err)
	}
	log.Println("database connected...!!")
}
func autoMigrate()error{
	err := DB.AutoMigrate(
		&models.User{},
		&models.Movie{},
		&models.Screen{},
		&models.Seat{},
		&models.Show{},
		&models.Booking{},
		&models.BookedSeat{},
	)
	return err
}