package main

import "os"

func getCultCookie() string {
	cookie := os.Getenv("CULT_COOKIE")
	if len(cookie) == 0 {
		panic("CULT_COOKIE env var not passed")
	}
	return cookie
}

func getCultAPIKey() string {
	apiKey := os.Getenv("CULT_API_KEY")
	if len(apiKey) == 0 {
		panic("CULT_API_KEY env var not passed")
	}
	return apiKey
}
