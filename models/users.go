package models

import (
	"context"
	"fazz/backend/lib"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type User struct {
	Id       int    `json:"id" db:"id"`
	Email    string `json:"email" db:"email" form:"email"`
	Password string `json:"-" db:"password" form:"password"`
	Username string `json:"username" db:"username" form:"username"`
}

func FindAllUsers() []User {
	db := lib.DB()
	defer db.Close(context.Background())

	rows, _ := db.Query(context.Background(),
		`select * from users`)

	users, err := pgx.CollectRows(rows, pgx.RowToStructByPos[User])

	if err != nil {
		fmt.Println(err)
	}

	return users
}

func FindOneUser(id int) User {
	db := lib.DB()
	defer db.Close(context.Background())

	rows, _ := db.Query(context.Background(),
		`select * from users`)

	users, err := pgx.CollectRows(rows, pgx.RowToStructByPos[User])

	if err != nil {
		fmt.Println(err)
	}
	var data User
	for _, v := range users {
		if v.Id == id {
			data = v
		}
	}
	return data
}

func InsertUser(email string, username string, password string) {
	db := lib.DB()
	defer db.Close(context.Background())

	dataSql := `insert into "users" (email, username, password) values ($1, $2, $3)`

	db.Exec(context.Background(), dataSql, email, username, password)
}

func EditUser(email string, username string, password string, id string) {
	db := lib.DB()
	defer db.Close(context.Background())

	dataSql := `update "users" set (email , username, password) = ($1, $2, $3) where id=$4`

	db.Exec(context.Background(), dataSql, email, username, password, id)

}

func RemoveData(id string) {
	db := lib.DB()
	defer db.Close(context.Background())

	dataSql := `delete from users where id =$1`

	db.Exec(context.Background(), dataSql, id)

}
