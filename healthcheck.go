package main

import (
	"fmt"
	"net/http"
	"time"
)

type application struct {
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

func doHealthCheck() {
	sites := []application{
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
			waitSeconds:    15,
		},
	}
	c := make(chan application)
	for _, site := range sites {
		go checkSiteHealth(site, c)
	}
	for l := range c {
		go func(a application) {
			fmt.Println("Waiting", a.waitSeconds, "seconds before retrying for", a.name)
			time.Sleep(time.Duration(a.waitSeconds) * time.Second)
			go checkSiteHealth(a, c)
		}(l)
	}
	fmt.Println("Done amigo!")
}

func checkSiteHealth(a application, c chan application) {
	resp, err := http.Get(a.url)
	if err != nil || resp.StatusCode != a.statusCode {
		fmt.Println(a.name + " is DOWN !!!")
		a.up = false
	}
	fmt.Println(a.name + " is up")
	a.up = true
	c <- a
}
