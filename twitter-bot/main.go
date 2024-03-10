package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ChimeraCoder/anaconda"
)

func main() {

	anaconda.SetConsumerKey(os.Getenv(""))
	anaconda.SetConsumerSecret(os.Getenv(""))
	api := anaconda.NewTwitterApi(os.Getenv(""), os.Getenv(""))

	tweet, err := api.PostTweet("Hello from my Golang Twitter bot!", nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Tweeted:", tweet.Text)

	retweet, err := api.Retweet(tweet.Id, false)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Retweeted:", retweet.Text)
}
