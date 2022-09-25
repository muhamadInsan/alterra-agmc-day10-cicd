package database

import (
	"day4/tugas-crud-dinamis/config"
	"day4/tugas-crud-dinamis/middlewares"
	"day4/tugas-crud-dinamis/models"
)

func GetUser() (interface{}, error) {
	var users []models.User
	// u := users{
	// 	"name":  Name,
	// 	"email": Email,
	// }

	if err := config.DB.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil

}

func GetUserId(id uint) (interface{}, error) {
	var users []models.User

	if err := config.DB.Find(&users, id).Error; err != nil {
		return nil, err
	}

	return users, nil
}

// func UpdateUsers(id int, users models.User) (interface{}, error) {
// var users []models.User

// DB.Save for save to db
// if err := config.DB.Where("id = ?", id).Updates(&users).Error; err != nil {
// 	return nil, err
// }
// if err := config.DB.Update(id, name, email).Error; err != nil {
// 	return nil, err
// }

// 	return users, nil
// }

func DeleteUser(id uint) (interface{}, error) {
	var users []models.User

	if err := config.DB.Delete(&users, id).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func LoginUser(user *models.User) (interface{}, error) {

	if err := config.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(user).Error; err != nil {
		return nil, err
	}

	token, err := middlewares.CreateToken(int(user.ID))
	if err != nil {
		return nil, err
	}

	return &token, nil
}

// func GetDetailUser(userId uint) (interface{}, error) {
// 	var user models.User

// 	if err := config.DB.Find(&user, userId).Error; err != nil {
// 		return nil, err
// 	}
// 	return user, nil
// }
