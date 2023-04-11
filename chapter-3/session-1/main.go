package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type WindAndWater struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

func main() {
	url := "https://jsonplaceholder.typicode.com/posts"

	for {
		valueWater := rand.Intn(100) + 1
		valueWind := rand.Intn(100) + 1

		var statusWater string
		if valueWater < 5 {
			statusWater = "Aman"
		} else if valueWater >= 6 && valueWater <= 8 {
			statusWater = "Siaga"
		} else {
			statusWater = "Bahaya"
		}

		var statusWind string
		if valueWind < 6 {
			statusWind = "Aman"
		} else if valueWind >= 7 && valueWind <= 15 {
			statusWind = "Siaga"
		} else {
			statusWind = "Bahaya"
		}

		windAndWater := WindAndWater{
			Water: valueWater,
			Wind:  valueWind,
		}

		jsonData, err := json.Marshal(windAndWater)
		if err != nil {
			log.Fatalf("Error: %v", err)
			return
		}

		resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			log.Fatalf("Error: %v", err)
			return
		}
		defer resp.Body.Close()

		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("Error: %v", err)
			return
		}

		fmt.Println(string(respBody))
		fmt.Println("Status Water: ", statusWater)
		fmt.Println("Status Wind: ", statusWind)

		time.Sleep(15 * time.Second)
	}
}
