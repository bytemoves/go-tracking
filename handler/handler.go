package handler

import (
	"log"
	"encoding/json"
	"net/http"
	"github.com/bytemoves/tracking-service/storages"
)

func NewHandler()  *http.ServeMux{
	mux := http.NewServeMux()
	mux.HandleFunc("/tracking",tracking)
	mux.HandleFunc("/search",search)

	return mux
}


func tracking (w http.ResponseWriter, r*http.Request){
	var driver = struct{
		ID string `json:"id"`
		Lat float64 `json:"lat"`
		Lng float64 `json:"lng"`
	} {}

	rClient := storages.GetRedisClient()

	if err := json.NewDecoder(r.Body).Decode(&driver); err != nil{
		log.Printf("could not decode request %v",err)
		http.Error(w,"could not decode request", http.StatusInternalServerError)
		return
	}

	rClient.AddDriverLocation(driver.Lng,driver.Lat,driver.ID)
	w.WriteHeader(http.StatusOK)
	return
}


func search (w http.ResponseWriter, r*http.Request){
	rClient  := storages.GetRedisClient()
	body := struct {
		Lat float64 `json:"lat"`
		Lng float64 `json:"lng"`
		Limit int `json:"limit"`
	}{}

	
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil{
		log.Printf("could not decode request %v",err)
		http.Error(w,"could not decode request", http.StatusInternalServerError)
		return
	}

	

	drivers := rClient.SearchDrivers(body.Limit, body.Lng, body.Lng,15)
	 data , err  := json.Marshal(drivers)
	 if err != nil{
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	 }
	 w.Header().Set( "Content-Type","application/json")
	 w.WriteHeader(http.StatusOK)
	 w.Write(data)
	 return
}








