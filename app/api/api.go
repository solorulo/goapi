package api

import (
	"upper.io/db"
	"upper.io/db/postgresql"
	"log"
)

var settings = db.Settings{
	Host:     "localhost",  // PostgreSQL server IP or name.
	Database: "goapi",    // Database name.
	Port: 	  5432,
	User:     "postgres",     // Optional user name.
	Password: "rulo",     // Optional user password.
}

// var settings = db.Settings{
// 	Host:     "ec2-54-204-36-244.compute-1.amazonaws.com",  // PostgreSQL server IP or name.
// 	Database: "d6dpto2avd599u",    // Database name.
// 	Port: 	  5432,
// 	User:     "dwdjidvoaxlrjb",     // Optional user name.
// 	Password: "udKmOLPLpL4g-Dpw0kXuuIMIbV",     // Optional user password.
// }
type User struct {
	Id			int64  `db:"user_id"`
	Nombre 		string `db:"nombre"`
	Apellido  	string `db:"apellido"`
	Email       string `db:"email"`
	Pass 		string `db:"pass"`
}

func MakeMap(code int, msg string) map[string]interface{} {
	data := map[string]interface{} {
		"code"	:	code,
		"msg"	:	msg,
	}
	return data
}

func MakeMapData(code int, msg string, data interface{}) map[string]interface{} {
	res := map[string]interface{} {
		"code"	:	code,
		"msg"	:	msg,
		"data"	:	data,
	}
	return res
}

func NotLogged() map[string]interface{} {
	return MakeMap(0, "not logged")
}

func Error() map[string]interface{} {
	return MakeMap(0, "error")
}

func UserAuth(email, pass string) (User, bool) {
	error := false
	sess, err := db.Open(postgresql.Adapter, settings)
  	defer sess.Close()
	if err != nil {
    	log.Fatalf("db.Open(): %q\n", err)
    	error = true
  	}
  	col, err := sess.Collection("users")
  	if err != nil {
    	log.Fatalf("cols Users\n", err)
    	error = true
  	}
  	res := col.Find(db.Cond{"email": email, "pass":pass})
  	defer res.Close()
  	var user User
  	err = res.One(&user)
  	if err != nil {
    	log.Fatalf("one User\n", err)
    	error = true
  	}
	return  user, error
}

func GetUser(id int64) (User, bool) {
	error := false
	sess, err := db.Open(postgresql.Adapter, settings)
  	defer sess.Close()
	if err != nil {
    	log.Fatalf("db.Open(): %q\n", err)
    	error = true
  	}
  	col, err := sess.Collection("users")
  	if err != nil {
    	log.Fatalf("cols Users\n", err)
    	error = true
  	}
  	res := col.Find(db.Cond{"user_id": id})
  	defer res.Close()
  	var user User
  	err = res.One(&user)
  	if err != nil {
    	log.Fatalf("one User\n", err)
    	error = true
  	}
	return  user, error
}

func GetUsers() ([]User, bool) {
	error := false
	sess, err := db.Open(postgresql.Adapter, settings)
  	defer sess.Close()
	if err != nil {
    	log.Fatalf("db.Open(): %q\n", err)
    	error = true
  	}
  	col, err := sess.Collection("users")
  	if err != nil {
    	log.Fatalf("cols Users\n", err)
    	error = true
  	}
  	res := col.Find()
  	defer res.Close()
  	var users []User
  	err = res.All(&users)
  	if err != nil {
    	log.Fatalf("list users\n", err)
    	error = true
  	}
	return  users, error
}

func EditUser(id int64, nombre, apellido, email, password string) bool {
	error := false
	sess, err := db.Open(postgresql.Adapter, settings)
  	defer sess.Close()
	if err != nil {
    	log.Fatalf("db.Open(): %q\n", err)
    	error = true
  	}
  	col, err := sess.Collection("users")
  	if err != nil {
    	log.Fatalf("cols Users\n", err)
    	error = true
  	}
  	res := col.Find(db.Cond{"user_id": id})
  	defer res.Close()
  	err = res.Update(map[string]interface{}{
		"nombre": nombre,
		"apellido": apellido,
		"email": email,
		"password": password,
	})
  	if err != nil {
    	log.Fatalf("edit User\n", err)
    	error = true
  	}
	return  error
}

func AddUser(nombre, apellido, email, password string) bool {
	error := false
	sess, err := db.Open(postgresql.Adapter, settings)
  	defer sess.Close()
	if err != nil {
    	log.Fatalf("db.Open(): %q\n", err)
    	error = true
  	}
  	col, err := sess.Collection("users")
  	if err != nil {
    	log.Fatalf("cols Users\n", err)
    	error = true
  	}
  	_, err = col.Append(map[string]interface{}{
		"nombre": nombre,
		"apellido": apellido,
		"email": email,
		"password": password,
	})
  	if err != nil {
    	log.Fatalf("add User\n", err)
    	error = true
  	}
	return  error
}
