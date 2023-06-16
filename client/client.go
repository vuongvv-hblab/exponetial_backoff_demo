package main

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/cenkalti/backoff/v4"
)

func main() {
	bo := backoff.NewExponentialBackOff()
	bo.InitialInterval = 500 * time.Millisecond
	bo.RandomizationFactor = 0.4
	bo.MaxInterval = 20 * time.Second
	bo.Multiplier = 2
	bo.MaxElapsedTime = 2 * time.Minute

	client := http.Client{
		//for case 5
		//Timeout: 5 * time.Second,
	}
	// Operation that make request to server
	f := func() error {
		resp, _ := client.Get("http://localhost:8080/hello1")
		log.Printf("Response status code : %d", resp.StatusCode)
		if resp.StatusCode == 500 {
			log.Printf("Elapsed time: %d", bo.GetElapsedTime()/time.Second)
			log.Println("-------------------------------------")

			log.Printf("Request failed, wait %d seconds to retry", bo.NextBackOff()/time.Second)
			log.Println("-------------------------------------")
			return errors.New("error")
		}

		//Read the response body
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}

		//Convert the body to type string
		sb := string(body)
		log.Printf("Hello from server: '%s'", sb)
		return nil
	}

	// Retry operation f until get success response
	err := backoff.Retry(f, bo)
	if err != nil {
		log.Fatalf("BackOff stopped retrying with Error '%s'", err)
	}

}
