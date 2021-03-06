package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"github.com/joho/godotenv"
)

var accountKey string

func handler(w http.ResponseWriter, r *http.Request) {
	busStopCode := r.FormValue("busStopCode")
	log.Println("busStopCode: ", busStopCode)
	w.Header().Set("Access-Control-Allow-Origin", "*")

	url := fmt.Sprintf("http://datamall2.mytransport.sg/ltaodataservice/BusArrivalv2?BusStopCode=%v", busStopCode)
	res := callBusAPI(url, accountKey)
	//callBusAPI("http://datamall2.mytransport.sg/ltaodataservice/BusStops", accountKey)
	bytes, err := w.Write(res)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	log.Println(bytes, "bytes written")
}

func getAccountKey() string {
	godotenv.Load(".env")
	return os.Getenv("accountkey")
}

func callBusAPI(url, accountKey string) []byte {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil || req == nil {
		log.Println("err: ", err)
		return []byte{}
	}
	req.Header.Add("AccountKey", accountKey)
	req.Header.Add("accept", "application/json")
	res, err := client.Do(req)
	if err != nil {
		log.Println("error http get, err: ", err)
		return []byte{}
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("fail to read all, err: ", err)
		return []byte{}
	}
	log.Println("body:", string(body))
	defer res.Body.Close()
	fmt.Println("res:", res)
	return body
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/nextBusStop", handler)

	accountKey = getAccountKey()
	log.Println("serving at 8082")
	log.Fatal(http.ListenAndServe(":8082", mux))
}
