package functions

import (
	"demo/db"
	"encoding/json"
	"fmt"
	"net/http"
)

func Signin(w http.ResponseWriter, r *http.Request) {
	var creds db.User
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		fmt.Println("error")
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w)
		return
	}

	phone := creds.Phone
	users, found := GetUserAutho(phone)
	if found {
		json.NewEncoder(w).Encode(users)
		if creds.Otp == users.Otp {
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusAccepted)
			json.NewEncoder(w)
			return
		} else {
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w)
			return
		}
	} else {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusNotFound)
		otp := CreateOTP()
		auth := db.AuthoData{
			Phone: phone,
			Otp:   otp,
		}
		json.NewEncoder(w).Encode(auth)
		InsertAutoData(auth)
		return
	}

}
func Welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello")
}
