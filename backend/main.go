package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	byte, err := w.Write([]byte("halo"))
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	log.Println("byte", byte, "written")
	//
	//call the API here
	accountKey := os.Getenv("accountkey")
	callBusAPI(accountKey)
	//callBusStopAPI(accountKey)
}

func callBusAPI(accountKey string) {
	url := "http://datamall2.mytransport.sg/ltaodataservice/BusArrivalv2?BusStopCode=83139"

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil || req == nil {
		log.Println("err: ", err)
		return
	}
	req.Header.Add("AccountKey", accountKey)
	req.Header.Add("accept", "application/json")
	res, err := client.Do(req)
	//res, err := http.Post(url, "application/json", bytes.NewReader(jsonBody))
	if err != nil {
		log.Println("error http get, err: ", err)
		return
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("fail to read all, err: ", err)
		return
	}
	log.Println("body:", string(body))
	defer res.Body.Close()
	fmt.Println("res:", res)
}

func callBusStopAPI(accountKey string) {
	url := "http://datamall2.mytransport.sg/ltaodataservice/BusStops"

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil || req == nil {
		log.Println("err: ", err)
		return
	}
	req.Header.Add("AccountKey", accountKey)
	req.Header.Add("accept", "application/json")
	res, err := client.Do(req)
	if err != nil {
		log.Println("error http get, err:", err)
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("error read all, err:", err)
		return
	}
	log.Println("body:", string(body))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
