package services

import (
	"errors"
	models "movion/Models"
	repositories "movion/Repositories"
)

func CreteShow(input *models.Show) (*models.Show, error) {
	if input.MovieID == 0 || input.ScreenID == 0 || input.ShowTime.IsZero() {
		return nil, errors.New("movie_id, screen_id and show_time are required")
	}
	showdata := input.ShowTime.Format("2006-01-02")
	count, err := repositories.CountShowsForScreenOnDate(input.ScreenID, showdata)
	if err != nil {
		return nil, errors.New("failed to check screen schedule")
	}
	if count >= 3 {
		return nil, errors.New("maximum 3 shows allowed per screen per day")
	}
	movieCount, err := repositories.CountDistinctMoviesOnDate(showdata)
	if err != nil {
		return nil, errors.New("failed to check daily movie count")
	}
	exists, err := repositories.MovieScheduledOnDate(input.MovieID, showdata)
	if err != nil {
		return nil, errors.New("failed to check movie schedule")
	}
	if movieCount >= 3 && !exists {
		return nil, errors.New("only 3 different movies allowed per day")
	}

	screen, err := repositories.GetScreenById(input.ScreenID)
	if err != nil {
		return nil, errors.New("screen not found")
	}

	input.AvailableSeats = screen.TotalSeats

	if err := repositories.CreateShow(input); err != nil {
		return nil, err
	}

	// LOAD RELATED MOVIE + SCREEN
	result, err := repositories.GetShowWithRelations(input.ID)
	if err != nil {
		return nil, errors.New("failed to load show relations")
	}

	return result, nil
}

func GetShowByID(id uint) (*models.Show, error) {
	return repositories.GetShowByIdWithPreload(id)
}

func GetAllShows() ([]models.Show, error) {
	return repositories.GetAllShows()
}

// func GetShowsByScreen(screenID uint) ([]models.Show, error) {
// return repositories.GetShowsByScreen(screenID)
// }

func EditShow(id uint, data *models.EditShow) (*models.Show, error) {
	show, err := repositories.GetShowById(id)
	if err != nil {
		return nil, errors.New("show not found")
	}

	if !data.ShowTime.IsZero() {
		show.ShowTime = data.ShowTime
	}
	if data.Price != 0 {
		show.Price = data.Price
	}
	// if data.Screen.TotalSeats != 0 {
	// show.Screen.TotalSeats = data.Screen.TotalSeats
	// // optionally update available seats
	// show.AvailableSeats = data.Screen.TotalSeats
	// }
	// Update MovieID
	if data.MovieID != 0 {
		show.MovieID = data.MovieID
	}

	// Update ScreenID
	if data.ScreenID != 0 {
		show.ScreenID = data.ScreenID
	}
	if err := repositories.SaveShow(show); err != nil {
		return nil, err
	}
	updated, err := repositories.LoadShowRelations(show)
	if err != nil {
		return nil, err
	}

	return updated, nil
}
func DeleteShow(id uint) error {
	return repositories.DeleteShow(id)
}
