package repositories

import (
	"encoding/json"
	// "errors"
	"movion/Config"
	"movion/Models"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// ---------------------- SHOW ----------------------

func GetShowForUpdate(tx *gorm.DB, showID uint) (*models.Show, error) {
	var show models.Show
	err := tx.Set("gorm:query_option", "FOR UPDATE").First(&show, showID).Error
	return &show, err
}

// ---------------------- BOOKED SEATS ----------------------

func IsSeatBooked(tx *gorm.DB, showID uint, seat string) (bool, error) {
	var count int64
	err := tx.Model(&models.BookedSeat{}).
		Where("show_id = ? AND seat_number = ?", showID, seat).
		Count(&count).Error
	return count > 0, err
}

func InsertBookedSeats(tx *gorm.DB, showID uint, seats []string) error {
	items := make([]models.BookedSeat, 0, len(seats))
	for _, s := range seats {
		items = append(items, models.BookedSeat{ShowID: showID, SeatNumber: s})
	}
	return tx.Create(&items).Error
}

func DeleteBookedSeats(tx *gorm.DB, showID uint, seats []string) error {
	return tx.Where("show_id = ? AND seat_number IN ?", showID, seats).
		Delete(&models.BookedSeat{}).Error
}

// ---------------------- BOOKING ----------------------

func CreateBookingRecord(tx *gorm.DB, userID, showID uint, seats []string, price float64) (*models.Booking, error) {
	seatJSON, _ := json.Marshal(seats)

	booking := models.Booking{
		UserID:           userID,
		ShowID:           showID,
		SeatNumbers:      datatypes.JSON(seatJSON),
		TotalSeatsBooked: len(seats),
		Amount:           float64(len(seats)) * price,
		Status:           "booked",
	}

	err := tx.Create(&booking).Error
	return &booking, err
}
func GetLatestBooking(userID, showID uint) (*models.Booking, error) {
    var booking models.Booking

    // 1. Load booking only
    err := config.DB.
        Where("user_id = ? AND show_id = ?", userID, showID).
        Order("created_at DESC").
        First(&booking).Error
    if err != nil {
        return nil, err
    }

    // 2. FORCE RELOAD SHOW (fresh updated data)
    var freshShow models.Show
    if err := config.DB.
        Preload("Movie").
        Preload("Screen").
        Preload("Screen.Seats").
        First(&freshShow, booking.ShowID).Error; err != nil {
        return nil, err
    }

    // override old show
    booking.Show = freshShow

    // 3. reload user
    config.DB.First(&booking.User, booking.UserID)

    return &booking, nil
}



func GetBookingByID(tx *gorm.DB, bookingID uint) (*models.Booking, error) {
	var booking models.Booking
	err := tx.First(&booking, bookingID).Error
	return &booking, err
}

func UpdateShowAvailableSeats(tx *gorm.DB, showID uint, amount int) error {
	return tx.Model(&models.Show{}).
		Where("id = ?", showID).
		Update("available_seats", gorm.Expr("available_seats + ?", amount)).
		Error
}

func MarkBookingCancelled(tx *gorm.DB, booking *models.Booking) error {
	booking.Status = "cancelled"
	return tx.Save(booking).Error
}
