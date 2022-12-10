package repository

import (
	"modulo/database"
	"modulo/models"
)

func CreateUser(user models.UserModel) (models.UserModel, error) {
	db, err := database.Connect()
	if err != nil {
		return models.UserModel{}, err
	}
	defer db.Close()

	statement, err := db.Prepare("insert into users (name, email) values (?, ?)")
	if err != nil {
		return models.UserModel{}, err
	}
	defer statement.Close()

	insert, err := statement.Exec(user.Name, user.Email)
	if err != nil {
		return models.UserModel{}, err
	}

	insertedId, err := insert.LastInsertId()
	if err != nil {
		return models.UserModel{}, err
	}

	user.ID = uint32(insertedId)
	return user, nil
}

func GetAllusers() ([]models.UserModel, error) {
	db, err := database.Connect()
	if err != nil {
		return []models.UserModel{}, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT * from users")
	if err != nil {
		return []models.UserModel{}, err
	}
	defer rows.Close()
	var users []models.UserModel

	for rows.Next() {
		var user models.UserModel
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return []models.UserModel{}, err
		}

		users = append(users, user)
	}

	return users, nil
}

func GetUser(id uint) (models.UserModel, error) {
	db, err := database.Connect()
	if err != nil {
		return models.UserModel{}, err
	}
	defer db.Close()

	row, err := db.Query("SELECT * FROM users WHERE id = ?", id)
	if err != nil {
		return models.UserModel{}, err
	}
	defer row.Close()

	var user models.UserModel
	if row.Next() {
		if err := row.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return models.UserModel{}, err
		}
	}

	return user, nil
}

func UpdateUser(user models.UserModel, ID uint32) error {
	db, err := database.Connect()
	if err != nil {
		return err
	}
	defer db.Close()

	statement, err := db.Prepare("UPDATE users SET name = ?, email = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(user.Name, user.Email, ID)
	if err != nil {
		return err
	}

	return nil
}

func DeleteUser(ID uint64) error {
	db, err := database.Connect()
	if err != nil {
		return err
	}
	defer db.Close()

	statement, err := db.Prepare("DELETE FROM users WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(ID)
	if err != nil {
		return err
	}

	return nil
}
