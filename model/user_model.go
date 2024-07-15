package model

import (
	"BlogAPI/pkg/database"
	"errors"
)

type User struct {
	ID       *int    `json:"id"`
	GoogleID *string `json:"google_id"`
	UserName string  `json:"username"`
	Password string  `json:"password"`
	Email    *string `json:"email"`
}

func (u *User) Create(username string, password string, email string) error {
	res, err := u.CheckIfExist("username", username)
	if err != nil {
		return err
	}
	if res {
		err2 := errors.New("username already exist")
		return err2
	}
	res, err = u.CheckIfExist("email", email)
	if err != nil {
		return err
	}
	if res {
		err2 := errors.New("email already exist")
		return err2
	}
	// Default value for user
	*u = User{
		ID:       nil,
		GoogleID: nil,
		UserName: username,
		Password: password,
		Email:    &email,
	}
	err = u.InsertToDB()
	if err != nil {
		return err
	}
	return nil
}

func (u *User) CheckIfExist(field string, value string) (bool, error) {
	db := database.GetInstance()
	var count bool
	err := db.QueryRow(`select count(*) from Users where ` + field + ` = ` + `"` + value + `";`).Scan(&count)
	if err != nil {
		return false, err
	}
	if count {
		return true, nil
	}
	return false, nil
}
func (u *User) InsertToDB() error {
	db := database.GetInstance()
	strQuery := "INSERT INTO `Users` (username, password, email) Values(?,?,?)"
	_, err := db.Query(strQuery, u.UserName, u.Password, u.Email)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) GetUserByUsername(username string) error {
	db := database.GetInstance()
	row, err := db.Query(`SELECT id, google_id, username, password, email FROM Users Where username = "` + username + `";`)
	if err != nil {
		return err
	}
	for row.Next() {
		err2 := row.Scan(&u.ID, &u.GoogleID, &u.UserName, &u.Password, &u.Email)
		if err2 != nil {
			return err2
		}
	}
	return nil
}
