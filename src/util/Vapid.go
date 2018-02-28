package util

import (
	"fmt"
		
	webpush "github.com/sherclockholmes/webpush-go"
)

func GetVapidKey(w http.ResponseWriter, r *http.Request) {

	privateKey, publicKey, err := webpush.GenerateVAPIDKeys()
	if err != nil {
		// TODO: Handle failure!
	}
	fmt.Println("privateKey:", privateKey)
	fmt.Println("publicKey:", publicKey)
	
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(vapidKeys)
	
}
