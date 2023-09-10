package main

import (
	"example/server/coinapp"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Web Client")

	// http.DefaultClient - empty client, same as &http.Client{}
	// response, err := http.DefaultClient.Get("https://dummyjson.com/products/1?jane=qwerty")

	// client := &http.Client{
	// 	// func to handle redirects
	// 	CheckRedirect: func(req *http.Request, via []*http.Request) error {
	// 		fmt.Println("REDIRECT")
	// 		return nil
	// 	},
	// 	Timeout: time.Second * 5,
	// 	// function to read/modify req before sending
	// 	Transport: &loggerRoundTripper{
	// 		next: http.DefaultTransport,
	// 	},
	// }

	client, _ := coinapp.NewHTTPClient(time.Minute)

	client.GetAssets().Data[0].PrintCoinappInfo()

}
