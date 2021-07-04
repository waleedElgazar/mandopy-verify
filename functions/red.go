package functions

import (
	"fmt"

	"github.com/go-redis/redis/v8"
)

var client = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "",
	DB:       0,
})

func SetPhoneRedis(phone string) {
	err := client.Set(client.Context(), "phone", phone, 0).Err()
	if err != nil {
		fmt.Println(err)
	}
}

func SetOtpRedis(otp string) {
	err := client.Set(client.Context(), "otp", otp, 0).Err()
	if err != nil {
		fmt.Println(err)
	}
}

func GetPhoneRedis()string{
	phone, err := client.Get(client.Context(), "phone").Result()
	if err != nil {
		fmt.Println(err)
	}
	return phone
}

func GetOtpRedis()string{
	otp, err := client.Get(client.Context(), "otp").Result()
	if err != nil {
		fmt.Println(err)
	}
	return otp
}



