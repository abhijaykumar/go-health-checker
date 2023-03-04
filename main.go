package main

import (
	"fmt"
	"net/http"
	"time"
)

type app struct {
	name           string
	url            string
	httpMethod     string
	timeoutSeconds int
	retryCount     int
	request        string
	response       string
	statusCode     int
	up             bool
	waitSeconds    int
}

func main() {
	sites := []app{
		{
			name:           "Google",
			url:            "https://google.com",
			httpMethod:     "GET",
			statusCode:     200,
			timeoutSeconds: 30,
			retryCount:     3,
			waitSeconds:    5,
		},
		{
			name:           "Facebook",
			url:            "https://facebook.com",
			httpMethod:     "GET",
			statusCode:     200,
			timeoutSeconds: 30,
			retryCount:     3,
			waitSeconds:    10,
		},
		{
			name:           "StackOverflow",
			url:            "https://stackoverflow.com",
			httpMethod:     "GET",
			statusCode:     200,
			timeoutSeconds: 30,
			retryCount:     3,
			waitSeconds:    1,
		},
	}
	c := make(chan app)
	for _, site := range sites {
		go checkSiteHealth(site, c)
	}
	for l := range c {
		go func(a app) {
			time.Sleep(time.Duration(a.waitSeconds) * time.Second)
			fmt.Println(a)
			go checkSiteHealth(a, c)
		}(l)
	}
	fmt.Println("Done amigo!")
}

func checkSiteHealth(a app, c chan app) {
	resp, err := http.Get(a.url)
	if err != nil || resp.StatusCode != a.statusCode {
		fmt.Println(a.name + " is DOWN !!!")
		a.up = false
	}
	fmt.Println(a.name + " is up")
	a.up = true
	c <- a
}
