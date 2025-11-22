package repositories

import (
	config "movion/Config"
	models "movion/Models"
)

func CreateMovie(movie *models.Movie) error{
	return  config.DB.Create(movie).Error
}
func GetAllMovies()([]models.Movie, error){
	var movies []models.Movie
	err := config.DB.Find(&movies).Error
	return movies, err
}
func GetMovie(id uint)(*models.Movie, error){
	var movie models.Movie
	err := config.DB.First(&movie, id).Error
	return &movie, err
}
func EditMovie(movie *models.Movie) error{
	return config.DB.Save(movie).Error
}
func DeleteMovie(id uint) error{
	return config.DB.Unscoped().Delete(&models.Movie{}, id).Error
}