package siap

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"

	"github.com/whatsauth/watoken"
	"github.com/whatsauth/whatsauth"
)

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func GetLoginInfofromPhoneNumber(phonenumber string, db *sql.DB) (response whatsauth.LoginInfo) {
	fmt.Println("phonenumber : " + phonenumber)
	if phonenumber != "" {
		response.Username = GetUsernamefromPhonenumber(phonenumber, db)
		fmt.Println("username : " + response.Username)
		if response.Username != "" {
			response.Password = UpdatePasswordfromUsername(response.Username, db)
			fmt.Println("password : " + response.Password)
			if response.Password != "" {
				response.Login = "Login"
				response.Userid = GetUserIdfromUsername(response.Username, db)
			}

		}
	}
	return response
}

func GetUsernamefromPhonenumber(phone_number string, db *sql.DB) (username string) {
	username = GetUsernamefromPhonenumberInTable(phone_number, "Pass", db)
	fmt.Println(username)
	if username == "" {
		username = GetUsernamefromPhonenumberInTable(phone_number, "Pass", db)

	}
	return username
}

func GetUsernamefromPhonenumberInTable(phone_number string, tabel string, db *sql.DB) (username string) {
	err := db.QueryRow("select Nama from dbo." + tabel + " where Phone = " + phone_number).Scan(&username)
	if err != nil {
		fmt.Printf("GetUsernamefromPhonenumberInTable %v: %v\n", tabel, err)
	}
	return username
}

func GetHashPasswordfromUsername(username string, db *sql.DB) (hashpassword string) {
	err := db.QueryRow("select Password from dbo.Pass where Nama = " + username).Scan(&hashpassword)
	if err != nil {
		fmt.Printf("GetHashPasswordfromUsername: %v\n", err)
	}
	return hashpassword
}

func UpdatePasswordfromUsername(username string, db *sql.DB) (newPassword string) {
	newPassword = watoken.RandomString(10)
	hashpass := GetMD5Hash(newPassword)
	var temp interface{}
	err := db.QueryRow("update dbo.Pass set Password = '" + hashpass + "' where Nama = '" + username + "'").Scan(&temp)
	if err != nil {
		fmt.Printf("UpdatePasswordfromUsername: %v\n", err)
	}
	return newPassword
}

func GetUserIdfromUsername(username string, db *sql.DB) (userid string) {
	err := db.QueryRow("select id from dbo.pass where Nama = '" + username + "'").Scan(&userid)
	if err != nil {
		fmt.Printf("GetUserIdfromUsername: %v\n", err)
	}
	return userid
}
