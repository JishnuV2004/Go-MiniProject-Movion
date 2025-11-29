package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"movion/Config"
	"movion/Models"
	"movion/Repositories"

	"gorm.io/gorm"
)

// --------------------------------------
// BOOK SEATS
// --------------------------------------
func BookSeats(userID, showID uint, seatNumbers []string) (*models.Booking, error) {
	err := config.DB.Transaction(func(tx *gorm.DB) error {

		// 1. Lock show row FOR UPDATE
		show, err := repositories.GetShowForUpdate(tx, showID)
		if err != nil {
			return errors.New("show not found")
		}

		if show.AvailableSeats < len(seatNumbers) {
			return errors.New("not enough seats available")
		}

		// 2. Check each seat
		for _, seat := range seatNumbers {
			booked, err := repositories.IsSeatBooked(tx, showID, seat)
			if err != nil {
				return err
			}
			if booked {
				return fmt.Errorf("seat %s already booked", seat)
			}
		}

		// 3. Insert booked seats
		if err := repositories.InsertBookedSeats(tx, showID, seatNumbers); err != nil {
			return err
		}

		// 4. Create booking
		_, err = repositories.CreateBookingRecord(tx, userID, showID, seatNumbers, show.Price)
		if err != nil {
			return err
		}

		// 5. Reduce available seats
		show.AvailableSeats -= len(seatNumbers)
		return tx.Save(show).Error
		//old
	})

	if err != nil {
		return nil, err
	}

	return repositories.GetLatestBooking(userID, showID)
}

// --------------------------------------
// CANCEL BOOKING
// --------------------------------------
func CancelBooking(bookingID, userID uint) error {
	return config.DB.Transaction(func(tx *gorm.DB) error {

		// 1. Get Booking
		booking, err := repositories.GetBookingByID(tx, bookingID)
		if err != nil {
			return errors.New("booking not found")
		}

		if booking.UserID != userID {
			return errors.New("not authorized")
		}

		if booking.Status != "booked" {
			return errors.New("booking already cancelled")
		}

		// 2. Read seat list
		var seats []string
		if err := json.Unmarshal(booking.SeatNumbers, &seats); err != nil {
			seats = []string{}
		}

		// 3. Delete booked seats
		if len(seats) > 0 {
			if err := repositories.DeleteBookedSeats(tx, booking.ShowID, seats); err != nil {
				return err
			}
		}

		// 4. Update seats back for show
		if err := repositories.UpdateShowAvailableSeats(tx, booking.ShowID, booking.TotalSeatsBooked); err != nil {
			return err
		}

		// 5. Mark booking cancelled
		return repositories.MarkBookingCancelled(tx, booking)
	})
}
