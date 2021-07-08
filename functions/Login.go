package functions

import (
	"bytes"
	"demo/db"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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
	foundUser, userData := getUser(w, r, phone)
	authData, foundAuth := GetUserAutho(phone)
	if foundAuth && foundUser {
		json.NewEncoder(w).Encode(userData)
		if creds.Otp == userData.Otp {
			w.WriteHeader(http.StatusAccepted)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
	}

	if foundAuth && !foundUser {
		json.NewEncoder(w).Encode(authData)
		if creds.Otp == authData.Otp {
			addUser(w, r, creds.Name, authData)
			w.WriteHeader(http.StatusAccepted)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}
	if !foundAuth && creds.Otp != "" {
		autho := db.AuthoData{
			Otp:   creds.Otp,
			Phone: creds.Phone,
		}
		InsertAutoData(autho)
		w.WriteHeader(http.StatusCreated)
	}

	/*
		if found {
			json.NewEncoder(w).Encode(userData)
			if userData.Otp == creds.Otp {
				w.WriteHeader(http.StatusAccepted)
			} else {
				w.WriteHeader(http.StatusUnauthorized)
			}
		} else {
			userAuthoData, found := GetUserAutho(phone)
			if found {
				addUser(w, r, creds.Name, userAuthoData)
			} else {
				fmt.Println(phone)
				InsertAutoData(phone)
				userAuthoData, _ = GetUserAutho(phone)
				json.NewEncoder(w).Encode(userAuthoData)
			}
			w.WriteHeader(http.StatusCreated)
		}*/

}
func Welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello")
}

func getUser(w http.ResponseWriter, r *http.Request, phone string) (bool, db.User) {
	url := "https://gp-mandoob-users.herokuapp.com/getUser/" + phone
	response, err := http.Get(url)

	if err != nil {
		fmt.Println("error", err.Error())
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("error", err.Error())
	}
	_, err = ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var user db.User
	err = json.Unmarshal(responseData, &user)
	if err != nil {
		fmt.Println("error", err.Error())
	}
	return user.Phone == phone, user
}

func addUser(w http.ResponseWriter, r *http.Request, name string, auth db.AuthoData) {
	values := map[string]string{"name": name, "phone": auth.Phone, "otp": auth.Otp}
	json_data, err := json.Marshal(values)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := http.Post("https://gp-mandoob-users.herokuapp.com/addUser", "application/json",
		bytes.NewBuffer(json_data))

	if err != nil {
		log.Fatal(err)
	}
	var res map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&res)
}

/*
	if found {
		json.NewEncoder(w).Encode(users)
		if creds.Otp == users.Otp {
			name := creds.Name
			var userData db.User
			UserFounded,userData:=getUser(w, r, phone)
			if  UserFounded{
				json.NewEncoder(w).Encode(userData)
				w.WriteHeader(http.StatusFound)
			} else {
				w.WriteHeader(http.StatusNotFound)
				addUser(w, r, name, users)
			}
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusAccepted)
			json.NewEncoder(w)
			return
		} else {
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusUnauthorized)
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
*/
