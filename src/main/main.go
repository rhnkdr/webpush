package main

import (
//	"bytes"
//	"encoding/json"
	"log"
	"net/http"

	//webpush "github.com/sherclockholmes/webpush-go"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"util"
)

//const (
//	vapidPrivateKey = "<YOUR VAPID PRIVATE KEY>"
//)

func main() {
	
	router := mux.NewRouter().StrictSlash(true) 
	
	router.HandleFunc("/getVapidKeys", util.GetVapidKey).Methods("POST") // get public vapid key
	router.HandleFunc("/pushSubscription", util.GetSubscription).Methods("POST") // get public vapid key
	c := cors.New(cors.Options{
        AllowedOrigins: []string{"*"},
        AllowCredentials: true,
        AllowedMethods: []string{"GET","POST","HEAD","OPTIONS","PUT"},
        AllowedHeaders: []string{"Content-Type","X-Requested-With","accept","Origin","Access-Control-Request-Method","Access-Control-Request-Headers","Authorization","AccessKey"},
        ExposedHeaders: []string{"Access-Control-Allow-Origin","Access-Control-Allow-Credentials"},
        MaxAge: 10,
    })
	
	handler := c.Handler(router)
	log.Fatal(http.ListenAndServe(":3000", handler)) 
	
}