package main

import (
	"fmt"

	"google.golang.org/api/youtube/v3"
)

func main() {
	client := getClient()
	service, err := youtube.New(client)
	if err != nil {
		panic(err)
	}

	call := service.LiveBroadcasts.List([]string{"status"}).Mine(true)
	res, err := call.Do()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", res.Items)
	for _, broadcast := range res.Items {
		if broadcast.Status.LifeCycleStatus == "complete" {
			continue
		}
		if broadcast.Status.PrivacyStatus == "private" {
			continue
		}
		println("https://youtu.be/"+broadcast.Id, broadcast.Status.LifeCycleStatus, broadcast.Status.PrivacyStatus)
		broadcast.Status.PrivacyStatus = "private"
		_, err := service.LiveBroadcasts.Update([]string{"status"}, broadcast).Do()
		if err != nil {
			panic(err)
		}
	}
}
