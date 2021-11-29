package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	authToken := os.Args[1]

	for _, a := range os.Args[2:] {
		request, err := http.NewRequest("GET", "https://members.tcmaker.org/api/v1/persons/"+a+"/?format=json", nil)
		request.Header.Add("Authorization", "Bearer "+authToken)
		request.Header.Add("Accept", "application/json")

		if err != nil {
			log.Fatalln(err)
		}

		client := http.Client{
			Timeout: time.Duration(30 * time.Second),
		}

		resp, err := client.Do(request)

		if err != nil {
			log.Fatalln(err)
		}

		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			log.Fatalln(err)
		}

		personData := make(map[string]string)
		json.Unmarshal([]byte(body), &personData)

		log.Println(personData["given_name"] + " " + personData["family_name"] + " - " + personData["email"] + " - " + personData["phone_number"])
	}
}
