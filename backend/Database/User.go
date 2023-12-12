package Database

import (

	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func Users(tkn string) ([]User, error) {
	results, err := db.Query("SELECT * FROM user WHERE token=?", tkn)
	if err != nil {
		return nil, err
	}
	defer results.Close() 

	var users []User
	for results.Next() {
		var u User
		err = results.Scan(&u.Username, &u.Password, &u.Email, &u.Token)
		if err != nil {
			return nil, err 
		}

		fmt.Println(u)
		users = append(users, u)
	}
  fmt.Println(users)
	return users, nil
}