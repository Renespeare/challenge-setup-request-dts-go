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

func getRandomNumber(min int, max int) (result int) {
	rand.New(rand.NewSource(10))
	result = rand.Intn(max - min + 1) + min
	return
}

func main()  {
	for {
    min := 1
    max := 100
    water := getRandomNumber(min, max)
		wind := getRandomNumber(min, max)
		data := map[string]interface{}{
			"water": water,
			"wind": wind,
		}
		requestJson, err := json.Marshal(data)
		client := &http.Client{}
		if err != nil {
			log.Fatalln(err)
		}

		req, err := http.NewRequest("POST", "https://jsonplaceholder.typicode.com/posts", bytes.NewBuffer(requestJson))
		req.Header.Set("Content-type", "application/json")
		if err != nil {
			log.Fatalln(err)
		}

		res, err := client.Do(req)
		if err != nil {
			log.Fatalln(err)
		}
		// defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatalln(err)
		}

		waterStatus := "aman"
		windStatus := "aman"

		switch {
		case water >= 6 && water <= 8:
			waterStatus = "siaga"
		case water > 8:
			waterStatus = "bahaya"
		default:
			waterStatus = "aman"
		}

		switch {
		case wind >= 7 && wind <= 15:
			windStatus = "siaga"
		case wind > 15:
			windStatus = "bahaya"
		default:
			windStatus = "aman"
		}

		fmt.Println(string(body))
		fmt.Println("status water:", waterStatus)
		fmt.Println("status wind:", windStatus)
		

		time.Sleep(time.Second * 15)
	}
}