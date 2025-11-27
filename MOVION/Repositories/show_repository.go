package repositories

import (
	config "movion/Config"
	models "movion/Models"
)

func CreateShow(show *models.Show) error {
	return config.DB.Create(show).Error
}

func CountShowsForScreenOnDate(screenID uint, date string) (int64, error) {
	var count int64
	err := config.DB.Model(&models.Show{}).Where("screen_id = ? AND DATE(show_time) = ?", screenID, date).
		Count(&count).Error
	return count, err
}
func CountDistinctMoviesOnDate(date string) (int64, error) {
	var count int64
	err := config.DB.Model(&models.Show{}).Where("DATE(show_time) = ?", date).Distinct("movie_id").Count(&count).Error
	return count, err
}
func MovieScheduledOnDate(movieID uint, date string) (bool, error) {
	var count int64
	err := config.DB.Model(&models.Show{}).Where("movie_id = ? AND DATE(show_time) = ?", movieID, date).Count(&count).Error
	return count > 0, err
}
func GetShowWithRelations(id uint) (*models.Show, error) {
	var show models.Show
	err := config.DB.Preload("Movie").Preload("Screen").First(&show, id).Error
	return &show, err
}

func LoadShowRelations(show *models.Show) (*models.Show, error) {
	err := config.DB.Preload("Movie").Preload("Screen").First(show, show.ID).Error
	return show, err
}

func GetShowByIdWithPreload(id uint) (*models.Show, error) {
	var show models.Show
	err := config.DB.Preload("Movie").Preload("Screen").First(&show, id).Error
	// err := config.DB.First(&show, id).Error
	return &show, err
}

func GetShowsByMovie(movieID uint) ([]models.Show, error) {
	var shows []models.Show
	err := config.DB.Where("movie_id = ?", movieID).Find(&shows).Error
	return shows, err
}

func GetShowsByScreen(screenID uint) ([]models.Show, error) {
	var shows []models.Show
	err := config.DB.Where("screen_id = ?", screenID).Find(&shows).Error
	return shows, err
}

func GetShowById(id uint) (*models.Show, error) {
	var show models.Show
	// err := config.DB.Preload("Movie").Preload("Screen").First(&show, id).Error
	err := config.DB.First(&show, id).Error
	return &show, err
}

func GetAllShows() ([]models.Show, error) {
	var shows []models.Show
	err := config.DB.Preload("Movie").Preload("Screen").Find(&shows).Error
	return shows, err
}

func SaveShow(show *models.Show) error {
	return config.DB.Save(show).Error
}

func DeleteShow(id uint) error {
	return config.DB.Unscoped().Delete(&models.Show{}, id).Error
}
