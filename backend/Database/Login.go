package Database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)


func Login(email, rpassword string) (string, bool) {
	exmail := Find(email)
	if (exmail == true) {
		password, err := Finds(email)
		fmt.Println(password, err, rpassword)
		if password == rpassword {
			token, err := FindToken(email)
			if err != nil {
				return ".", false
			} else {
				return token, true
			}
		}
	}
	return "", false
}


func Finds(email string) (string, error) {
	query := "SELECT password FROM user WHERE email = ?"
	row := db.QueryRow(query, email)
	var password string
	err := row.Scan(&password)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("Kullanıcı bulunamadı")
		}
		return "", err
	}

	return password, nil
}

func FindToken(email string) (string, error)  {
	query := "SELECT token FROM user WHERE email = ?"
	row := db.QueryRow(query, email)
	var token string
	err := row.Scan(&token)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("Kullanıcı bulunamadı")
		}
		return "", err
	}

	return token, nil
}