package services

import (
	"errors"
	config "movion/Config"
	models "movion/Models"
	repositories "movion/Repositories"
)

func CreateMovie(inputMovie *models.Movie)(*models.Movie, error){
	if inputMovie.Title == "" {
		return nil, errors.New("title is required")
	}
	err := repositories.CreateMovie(inputMovie)
	if err != nil {
		return nil, err
	}
	return inputMovie, err
}
func GetAllMovies() ([]models.Movie, error) {
    var movies []models.Movie
    err := config.DB.Find(&movies).Error
    return movies, err
}


func GetMovie(id uint)(*models.Movie, error){
	movie, err := repositories.GetMovie(id)
	if err != nil {
		return nil, errors.New("movie not found")
	}
	return movie,err
}

func EditMovie(id uint, input *models.Movie)(*models.Movie, error) {
	movie, err := repositories.GetMovie(id)
	if err != nil {
		return nil, errors.New("movie not found")
	}
	if input.Title != "" {
		movie.Title = input.Title
	}
	if input.Description != "" {
		movie.Description = input.Description
	}
	if input.Language != "" {
		movie.Language = input.Language
	}
	if input.DurationMin != 0 {
		movie.DurationMin = input.DurationMin
	}
	if input.ReleaseDate != "" {
		movie.ReleaseDate = input.ReleaseDate
	}
	if err := repositories.EditMovie(movie); err != nil {
		return nil, err
	}
	return movie, err
}

func DeleteMovie(id uint) error{
	return repositories.DeleteMovie(id)
}