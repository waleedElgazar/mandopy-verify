package main

import (
	"demo/functions"
	"log"
	"net/http"
	"os"
)

func main() {
	os.Setenv("DB_HOST", "us-cdbr-east-04.cleardb.com")
	os.Setenv("DB_PORT", "(us-cdbr-east-04.cleardb.com)")
	os.Setenv("DB_DRIVER", "mysql")
	os.Setenv("DB_ROOT", "b7b41cd66ae593")
	os.Setenv("DB_PASSWORD", "ca3b0054")
	os.Setenv("DB_NAME", "heroku_31c814737f81a30")
	//os.Setenv("PORT", "8081")
	port := os.Getenv("PORT")
	http.HandleFunc("/", functions.Welcome)
	http.HandleFunc("/verify", functions.Signin)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
