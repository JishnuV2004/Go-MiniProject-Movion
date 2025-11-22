package repositories

import (
	config "movion/Config"
	models "movion/Models"
)

func CreateScreen(screen *models.Screen) error{
	return config.DB.Create(screen).Error
}
func CreateSeat(seats []models.Seat) error{
	return config.DB.Create(&seats).Error
}
func GetScreenById(id uint)(*models.Screen, error){
	var screen models.Screen
	err := config.DB.Preload("Seats").First(&screen, id).Error
	return &screen, err
}
func GetAllScreens()([]models.Screen, error){
	var screens []models.Screen
	err := config.DB.Find(&screens).Error
	return screens, err
}
func EditScreen(editscreen *models.Screen) error{
	return config.DB.Save(editscreen).Error
}
func DeleteScreen(id uint) error{
    return config.DB.Delete(&models.Screen{}, id).Error
}
