package services

import (
	"errors"
	"fmt"
	models "movion/Models"
	repositories "movion/Repositories"
)

func CreateScreen(screenInput *models.Screen, rows, cols int)(*models.Screen, error){
	if screenInput.Name == "" {
		return nil, errors.New("screen name is required")
	}
	if rows <= 0 || cols <= 0 {
		return  nil, errors.New("invalid seat layout")
	}
	screenInput.TotalSeats = rows * cols
	if err := repositories.CreateScreen(screenInput); err != nil {
		return nil, err
	}

	// generate seats
	seats := make([]models.Seat, 0, screenInput.TotalSeats)
	for r := 0; r<= rows; r++ {
		rowChar := rune('A' + r)
		for c := 1; c<=cols; c++ {
			seats = append(seats, models.Seat{
				ScreenID: screenInput.ID,
				SeatNumber: fmt.Sprintf("%c%d", rowChar, c),
			})
		}
	}
	// preload seats
	if err := repositories.CreateSeat(seats); err != nil {
		return nil, err
	}
	screen, err := repositories.GetScreenById(screenInput.ID)
	if err != nil {
		return  nil, err
	}
	return screen, nil
}
func GetAllScreens()([]models.Screen, error){
	return repositories.GetAllScreens()
}
func GetScreen(id uint)(*models.Screen, error){
	return repositories.GetScreenById(id)
}
func EditScreen(id uint, input *models.Screen) (*models.Screen, error){
	screen, err := repositories.GetScreenById(id)
	if err != nil {
		return  nil, errors.New("screen not found")
	}
	if input.Name != "" {
		screen.Name = input.Name
	}
	if input.TotalSeats != 0 {
		screen.TotalSeats = input.TotalSeats
	}
	if err := repositories.EditScreen(screen); err != nil {
		return nil, errors.New("screen editing faild")
	}
	return screen, err
}
func DeleteScreen(id uint) error{
	return repositories.DeleteScreen(id)
}