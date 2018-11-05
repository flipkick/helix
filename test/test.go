package main

import (
	"fmt"
	"log"

	"github.com/nicklaw5/helix"
)

func main() {
	c, err := helix.NewClient(&helix.Options{
		ClientID:     "0htio5mjh3kd2espt44aldk1emkeuk",
		ClientSecret: "3u0996ko6bdvrrr0ubmh32u6sfqeik",
		RedirectURI:  "http://localhost:8888/auth/callback",
		// UserAccessToken: "",
		// AppAccessToken: "",
	})
	if err != nil {
		log.Fatal(err)
	}

	// GET AUTH URL
	// authURL := c.GetAuthorizationURL("some-state", false)
	// fmt.Printf("%+v\n", authURL)

	// GET USER ACCESS TOKEN
	code := "t2frg3cfl23pbke7llsdhc9vvnbba4"
	resp, err := c.GetUserAccessToken(code)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)

	// resp, err := c.PostWebhookSubscription(&helix.WebhookPayload{
	// 	Mode:         "subscribe",
	// 	Topic:        "https://api.twitch.tv/helix/users/follows?first=1&to_id=1337",
	// 	Callback:     "https://yourwebsite.com/path/to/callback/handler",
	// 	LeaseSeconds: 0,
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%+v\n", resp)
}
