package util

import (
	"fmt"
	"encoding/json"
	"net/http"
	"io/ioutil"
	"log"
	"bytes"
		
	webpush "github.com/sherclockholmes/webpush-go"
)

type VapidKeys struct {
	publicKeys string `json:"publicKey`
}

type SubscriptionData struct {
	Endpoint string `json:"endpoint"`
	Auth   string `json:"auth"`
	P256dh string `json:"p256dh"` 
}

//generate vapid keys and send the public key to client
func GetVapidKey(w http.ResponseWriter, r *http.Request) {

	privateKey, publicKey, err := webpush.GenerateVAPIDKeys()
	if err != nil {
		// TODO: Handle failure!
	}
	
	fmt.Println("Private Key", privateKey)
	fmt.Println("Public Key", publicKey)
	
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(publicKey)
	
}

//decode subscription object from client
func GetSubscription(w http.ResponseWriter, r *http.Request){
	
	fmt.Println("getSubs")
	
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	// Unmarshal
	var data SubscriptionData
	err = json.Unmarshal(b, &data)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	
	newData := new(SubscriptionData)
	newData.Endpoint = data.Endpoint
	newData.Auth = data.Auth
	newData.P256dh = data.P256dh
	
	output, err := json.Marshal(newData)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	
	stringOutput := string(output)
	subJSON := stringOutput

	// Decode subscription
	s := webpush.Subscription{}
	if err := json.NewDecoder(bytes.NewBufferString(subJSON)).Decode(&s); err != nil {
		log.Fatal(err)
	}
	
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(s)
}

// Send Notifaction to the subscribe users
//func SendNotification (w http.ResponseWriter, r *http.Request){
//		b, err := ioutil.ReadAll(r.Body)
//	defer r.Body.Close()
//	if err != nil {
//		http.Error(w, err.Error(), 500)
//		return
//	}
//	// Unmarshal
//	var data SubscriptionData
//	err = json.Unmarshal(b, &data)
//	if err != nil {
//		http.Error(w, err.Error(), 500)
//		return
//	}
//	
//	newData := new(SubscriptionData)
//	newData.Endpoint = data.Endpoint
//	newData.Auth = data.Auth
//	newData.P256dh = data.P256dh
//	
//	output, err := json.Marshal(newData)
//	if err != nil {
//		http.Error(w, err.Error(), 500)
//		return
//	}
//	
//	stringOutput := string(output)
//	subJSON := stringOutput
//
//	// Decode subscription
//	s := webpush.Subscription{}
//	if err := json.NewDecoder(bytes.NewBufferString(subJSON)).Decode(&s); err != nil {
//		log.Fatal(err)
//	}
//	_, err := webpush.SendNotification([]byte("Test"), &s, &webpush.Options{
//		VAPIDPrivateKey: vapidPrivateKey,
//	})
//	if err != nil {
//		log.Fatal(err)
//	}
//}