package Database

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Register(username, password, email, token string) bool {
	extinguser := Find(email) 
	if extinguser == false {
		err := addUser(username, password, email, token)
		if err == nil {
			return true
		}
		fmt.Print(err)
	} else {
	 return false
	}
	return true
}

func Find(email string) bool {
	query := "SELECT email FROM user WHERE email = ?"
	var result string
	err := db.QueryRow(query, email).Scan(&result)
	if err != nil {
		fmt.Println("\x1b[31m", err)
		return false
	}

	fmt.Print("User true")
	return true
}

func addUser(username, password, email, token string) error {
	query := "INSERT INTO user (username, password, email, token) VALUES (?, ?, ?, ?)"
	result, err := db.Exec(query, username, password, email, token)
	if err != nil {
		return err
	}

	rowCount, err := result.RowsAffected()
	if err != nil {
		return err
	}

	fmt.Printf("%d Added!\n", rowCount)
	return nil
}