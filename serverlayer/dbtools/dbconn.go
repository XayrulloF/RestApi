package dbtools

import (
	"database/sql"
	"log"
	"restProject/serverlayer/model"

	_ "github.com/go-sql-driver/mysql"
)

var driverName, dataSourceName string

func DbInit(dn, dsn string) {
	driverName, dataSourceName = dn, dsn
}
func connect() *sql.DB {
	dbconn, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatal(err.Error())
	}
	return dbconn
}

func SelectAllUsers() []model.User {
	db := connect()
	rows, err := db.Query("SELECT * FROM example.users;")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	users := []model.User{}
	for rows.Next() {
		user := model.User{}
		err := rows.Scan(&user.Id, &user.Name, &user.Age)
		if err != nil {
			log.Fatal(err.Error())
			continue
		}
		users = append(users, user)
	}
	return users
}

func SelectUserByName(name string) model.User {
	db := connect()
	row := db.QueryRow("SELECT * FROM example.users WHERE name=?;", name)
	defer db.Close()

	user := model.User{}

	err := row.Scan(&user.Id, &user.Name, &user.Age)
	if err != nil {
		log.Fatal(err.Error())
	}
	return user
}

func SelectUserByAge(age int32) model.User {
	db := connect()
	row := db.QueryRow("SELECT * FROM example.users WHERE age = ?;", age)
	defer db.Close()
	user := model.User{}

	err := row.Scan(&user.Id, &user.Name, &user.Age)
	if err != nil {
		log.Fatal(err.Error())
	}
	return user
}

func AddUser(user model.User) int64 {
	db := connect()
	addPrepare, err := db.Prepare("INSERT INTO example.users(name, age) VALUES (?, ?);")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()
	addExec, err := addPrepare.Exec(user.Name, user.Age)
	if err != nil {
		log.Fatal(err.Error())
	}
	lastId, err := addExec.LastInsertId()
	if err != nil {
		log.Fatal(err.Error())
	}
	return lastId
}

func UpdateUser(user model.User) int64 {
	db := connect()
	updPrepare, err := db.Prepare("UPDATE example.users SET name=?, age=? WHERE id=?;")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()
	updExec, err := updPrepare.Exec(&user.Name, &user.Age, &user.Id)
	if err != nil {
		log.Fatal(err.Error())
	}
	rowsAffected, err := updExec.LastInsertId()
	if err != nil {
		log.Fatal(err.Error())
	}
	return rowsAffected
}

func DeleteUser(user model.User) int64 {
	db := connect()
	delPrepare, err := db.Prepare("DELETE FROM example.users WHERE id=?;")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()
	delExec, err := delPrepare.Exec(user.Id)
	if err != nil {
		log.Fatal(err.Error())
	}
	rowsAffected, err := delExec.RowsAffected()
	if err != nil {
		log.Fatal(err.Error())
	}
	return rowsAffected
}
