package functions

import (
	"crypto/rand"
	"demo/db"
	"fmt"
	"io"
	"os"
)

func GetUserAutho(phone string) (db.AuthoData, bool) {
	var users db.AuthoData
	dbb := db.DBConn()
	defer dbb.Close()
	db_name := os.Getenv("DB_NAME")
	query := "SELECT phone, otp FROM " + db_name + ".AuthoData WHERE phone = ?"
	err := dbb.QueryRow(query, phone).Scan(&users.Phone, &users.Otp)
	if err != nil {
		fmt.Println(err)
		return users, false
	}
	return users, true
}

func InsertAutoData(auth db.AuthoData) bool {
	db := db.DBConn()
	defer db.Close()
	db_name := os.Getenv("DB_NAME")
	in := "INSERT INTO " + db_name + ".AuthoData VALUES(?,?)"
	insert, err := db.Prepare(in)
	insert.Exec(auth.Phone, auth.Otp)
	if err != nil {
		panic(err.Error())
	}
	return true
}

func CreateOTP() string {
	var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
	max := 6
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}
