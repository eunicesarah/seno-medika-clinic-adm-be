package superadmin

import (
	"seno-medika.com/config/db"
	"seno-medika.com/model/person"
)

func FindAll() ([]person.UserWithoutPassword, error) {
	var users []person.UserWithoutPassword
	rows, err := db.DB.Query("SELECT user_id, user_uuid, nama, email, role FROM users")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user person.UserWithoutPassword
		err = rows.Scan(&user.UserID, &user.UserUUID, &user.Nama, &user.Email, &user.Role)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func FindById(id int) (person.UserWithoutPassword, error) {
	var user person.UserWithoutPassword

	err := db.DB.QueryRow("SELECT user_id, user_uuid, nama, email, role FROM users WHERE user_id = $1", id).Scan(&user.UserID, &user.UserUUID, &user.Nama, &user.Email, &user.Role)
	if err != nil {
		return person.UserWithoutPassword{}, err
	}

	return user, nil
}

func FindByUuid(uid string) (person.UserWithoutPassword, error) {
	var user person.UserWithoutPassword

	err := db.DB.QueryRow("SELECT user_id, user_uuid, nama, email, role FROM users WHERE user_uuid = $1", uid).Scan(&user.UserID, &user.UserUUID, &user.Nama, &user.Email, &user.Role)
	if err != nil {
		return person.UserWithoutPassword{}, err
	}

	return user, nil
}

func FindByName(name string) ([]person.UserWithoutPassword, error) {
	var users []person.UserWithoutPassword
	rows, err := db.DB.Query("SELECT user_id, user_uuid, nama, email, role FROM users WHERE nama = $1", name)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user person.UserWithoutPassword
		err = rows.Scan(&user.UserID, &user.UserUUID, &user.Nama, &user.Email, &user.Role)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func FindByEmail(email string) ([]person.UserWithoutPassword, error) {
	var users []person.UserWithoutPassword
	rows, err := db.DB.Query("SELECT user_id, user_uuid, nama, email, role FROM users WHERE email = $1", email)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user person.UserWithoutPassword
		err = rows.Scan(&user.UserID, &user.UserUUID, &user.Nama, &user.Email, &user.Role)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func FindByRole(role string) ([]person.UserWithoutPassword, error) {
	var users []person.UserWithoutPassword
	rows, err := db.DB.Query("SELECT user_id, user_uuid, nama, email, role FROM users WHERE role = $1", role)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user person.UserWithoutPassword
		err = rows.Scan(&user.UserID, &user.UserUUID, &user.Nama, &user.Email, &user.Role)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}
