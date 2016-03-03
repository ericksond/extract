package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

// CreateAppdynamicsFile function
func (a *appdynamics) CreateAppdynamicsFile() {
	appdURL := a.BaseURL + a.Path
	fmt.Printf("[%s] Processing: %s\n", time.Now().Format("2016-03-02 15:04:05"), a.Name)

	// GET Request
	req, err := http.NewRequest("GET", appdURL, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth(a.User, a.Pass)
	req.Header.Add("Accept", "application/json")

	// HTTP Client
	cli := &http.Client{}

	if a.Proxy == true {
		fmt.Println("Using proxy...")
		proxyURL, err := url.Parse(a.ProxyURL)
		if err != nil {
			log.Fatal(err)
		}
		cli = &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyURL)}}
	}

	res, err := cli.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	// Generate JSON file
	file := a.BaseFilePath + a.Name + ".json"
	jsonfile, err := os.Create(file)
	if err != nil {
		log.Fatal(err)
	}
	defer jsonfile.Close()

	// Copy res.Body to JSON file
	if _, err := io.Copy(jsonfile, res.Body); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("[%s] Completed: %s\n", time.Now().Format("2016-03-02 15:04:05"), a.Name)
}
