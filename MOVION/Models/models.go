package models

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
	"gorm.io/datatypes"
)

type Admin struct{
		Email string `json:"email"`
		Password string `json:"password"`
		Role string 
	}

type User struct {
	gorm.Model
	Username  string `json:"username" gorm:"not null"`
	Email     string `json:"email" gorm:"unique;not null"`
	Password  string `json:"password"`
	Role      string `json:"role" gorm:"default:user"`
	IsBlocked bool   `json:"is_blocked" gorm:"default:false"`
}

// type User struct {
// 	gorm.Model
// 	Username string `json:"username"`
// 	Email    string `gorm:"unique" json:"email"`
// 	Password string `json:"password"`
// 	Role     string `json:"role"`
// }
type Claims struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	Block bool
	jwt.RegisteredClaims
}
type EditUserInput struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
type Movie struct {
	gorm.Model
	Title       string `json:"title" gorm:"not null"`
	Description string `json:"description"`
	Language    string `json:"language"`
	DurationMin int    `json:"duration_min"`
	ReleaseDate string `json:"release_date"`

	Shows []Show `json:"shows"` // relation
}
type Screen struct {
    gorm.Model
    Name       string     `json:"name"`
    TotalSeats int        `json:"totalseats"`
    Seats      []Seat     `json:"seats" gorm:"constraint:OnDelete:CASCADE;"`
    Shows      []Show     `json:"shows"`
}
type Seat struct {
    gorm.Model
    ScreenID   uint   `json:"screen_id"`
    SeatNumber string `json:"seat_number"`
}

type Show struct {
    gorm.Model
    MovieID        uint      `json:"movie_id"`
    ScreenID       uint      `json:"screen_id"`
    ShowTime       time.Time `json:"show_time"`
    Price          float64   `json:"price"`
    AvailableSeats int       `json:"available_seats"`

	Movie  Movie  `gorm:"foreignKey:MovieID" json:"movie"`
    Screen Screen `gorm:"foreignKey:ScreenID" json:"screen"`

    // Movie  Movie  `json:"movie"`
    // Screen Screen `json:"screen"`
}

type Booking struct {
    gorm.Model
    UserID           uint      `json:"user_id"`
    ShowID           uint      `json:"show_id"`
    SeatNumbers      datatypes.JSON `json:"seat_numbers"` // ["A1", "A2"]
    TotalSeatsBooked int       `json:"total_seats_booked"`
    Amount           float64   `json:"amount"`
    Status           string    `json:"status" gorm:"default:'booked'"`

    User User `json:"user"`
    Show Show `json:"show"`
}
type BookedSeat struct {
    gorm.Model
    ShowID     uint   `json:"show_id"`
    SeatNumber string `json:"seat_number"`
}

type EditShow struct {
    MovieID  uint      `json:"movie_id"`
    ScreenID uint      `json:"screen_id"`
    ShowTime time.Time `json:"show_time"`
    Price    float64   `json:"price"`
}



// Optional
// type Seat struct {
// 	gorm.Model
// 	ShowID     uint   `json:"show_id"`
// 	SeatNumber string `json:"seat_number"`
// 	IsBooked   bool   `json:"is_booked" gorm:"default:false"`

// 	Show Show `json:"show"`
// }

