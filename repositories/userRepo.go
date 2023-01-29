package repositories

import (
	"simple-catalog-v2/connect"
	"simple-catalog-v2/models"

	"golang.org/x/crypto/bcrypt"
)

var db = connect.MySqlConnect()

func GetUserInfoByEmailOrId(email string, id int) (models.User, error) {
	var user models.User
	// err := db.QueryRow("SELECT * FROM user where email = ? OR id = ?", email, id).Scan(&user.Id, &user.Fullname, &user.Username, &user.Email, &user.Gender, &user.Password, &user.Status, &user.Avatar)
	err := db.Debug().Where("email = ?", email).Or("id = ?", id).Find(&user).Error
	return user, err
} 

func CreateUser(data *models.User) error {
	// insert, err := db.Prepare("INSERT INTO user (fullname,username,email,gender,password,status) values (?,?,?,?,?,?)")
	// if err != nil {
	// 	return err
	// }
	// defer insert.Close()

	// _, err = insert.Exec(data.Fullname, data.Username, data.Email, data.Gender, data.Password, data.Status)
	err := db.Debug().Create(&data).Error
	return err
}

func UpdateUser(id int, data *models.DisplayUserData) error {
	// update, err := db.Prepare("UPDATE user SET fullname=?, username=?, email=?, gender=?, status=? WHERE id=?")
	// if err != nil {
	// 	return err
	// }
	// defer update.Close()

	// _, err = update.Exec(data.Fullname, data.Username, data.Email, data.Gender, data.Status, id)
	err := db.Debug().Model(&models.User{}).Select("fullname", "username", "gender", "status").Where("id = ?", id).Updates(data).Error
	return err
}

func SetUserRefreshToken(email string, token string) error {
	err := db.Debug().Model(&models.User{}).Where("Email = ?", email).Update("refresh_token", token).Error
	return err
}

func GetUserRefreshToken(email string) (string, error) {
	var token string
	err := db.Debug().Model(&models.User{}).Select("refresh_token").Where("Email = ?", email).Find(&token).Error
	return token, err
}

func SetUserPassword(email string, password string) error {
	pass, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	if err != nil { 
		return err
	}

	err = db.Debug().Model(&models.User{}).Where("Email = ?", email).Update("password", string(pass)).Error
	return err
}