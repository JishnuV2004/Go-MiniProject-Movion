package config

import (
	"log"
	models "movion/Models"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Jwtkey []byte

var DB *gorm.DB

func InitDB(){

	err := godotenv.Load()
	if err != nil {
		log.Println("env file not found, using system environment variables")
	}

	root := os.Getenv("DB_ROOT")
	jwtscret := os.Getenv("JWT_SECRET")
	Jwtkey = []byte(jwtscret)

	// var err error
	DB, err = gorm.Open(mysql.Open(root), &gorm.Config{})
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