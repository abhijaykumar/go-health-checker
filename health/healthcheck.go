package health

import (
	"fmt"
	"net/http"
	"time"
)

type Application struct {
	Name           string
	Url            string
	HttpMethod     string
	TimeoutSeconds int
	RetryCount     int
	Request        string
	Response       string
	StatusCode     int
	Up             bool
	WaitSeconds    int64
}

var Sites = []Application{
	{
		Name:           "Google",
		Url:            "https://google.com",
		HttpMethod:     "GET",
		StatusCode:     200,
		TimeoutSeconds: 30,
		RetryCount:     3,
		WaitSeconds:    5,
	},
	{
		Name:           "Facebook",
		Url:            "https://facebook.com",
		HttpMethod:     "GET",
		StatusCode:     200,
		TimeoutSeconds: 30,
		RetryCount:     3,
		WaitSeconds:    10,
	},
	{
		Name:           "StackOverflow",
		Url:            "https://stackoverflow.com",
		HttpMethod:     "GET",
		StatusCode:     200,
		TimeoutSeconds: 30,
		RetryCount:     3,
		WaitSeconds:    15,
	},
}

func doHealthCheck() {
	c := make(chan Application)
	for _, site := range Sites {
		go checkSiteHealth(site, c)
	}
	for l := range c {
		go func(a Application) {
			fmt.Println("Waiting", a.WaitSeconds, "seconds before retrying for", a.Name)
			time.Sleep(time.Duration(a.WaitSeconds) * time.Second)
			go checkSiteHealth(a, c)
		}(l)
	}
	fmt.Println("Done amigo!")
}

func checkSiteHealth(a Application, c chan Application) {
	resp, err := http.Get(a.Url)
	if err != nil || resp.StatusCode != a.StatusCode {
		fmt.Println(a.Name + " is DOWN !!!")
		a.Up = false
	}
	fmt.Println(a.Name + " is up")
	a.Up = true
	c <- a
}
