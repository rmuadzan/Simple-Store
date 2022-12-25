package repositories

import (
	"simple-catalog-v2/connect"
	"simple-catalog-v2/models"
)

var db = connect.MySqlConnect()

func GetUserInfoByEmailOrId(email string, id int) (models.User, error) {
	var user models.User
	err := db.QueryRow("SELECT * FROM user where email = ? OR id = ?", email, id).Scan(&user.Id, &user.Fullname, &user.Username, &user.Email, &user.Gender, &user.Password, &user.Status, &user.Avatar)
	return user, err
} 

func CreateUser(data *models.User) error {
	insert, err := db.Prepare("INSERT INTO user (fullname,username,email,gender,password,status) values (?,?,?,?,?,?)")
	if err != nil {
		return err
	}
	defer insert.Close()

	_, err = insert.Exec(data.Fullname, data.Username, data.Email, data.Gender, data.Password, data.Status)
	return err
}

func UpdateUser(id int, data *models.DisplayUserData) error {
	update, err := db.Prepare("UPDATE user SET fullname=?, username=?, email=?, gender=?, status=? WHERE id=?")
	if err != nil {
		return err
	}
	defer update.Close()

	_, err = update.Exec(data.Fullname, data.Username, data.Email, data.Gender, data.Status, id)
	return err
}
