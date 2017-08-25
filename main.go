package main

import (
	"fmt"
	"github.com/turnage/graw/reddit"

)

func main() {
	bot, err := reddit.NewBotFromAgentFile("login.agent", 0)
	if err != nil {
		fmt.Println("Failed to create bot handle: ", err)
		return
	}

	//listing returns the object you need to get information from the subreddit
	harvest, err := bot.Listing("/r/memes", "")
	if err != nil {
		fmt.Println("Failed to fetch /r/memes: ", err)
		return
	}

	//subreddit has posts
	//get the last 10 posts of this subreddit
	for _, post := range harvest.Posts[:3] {
		fmt.Printf("[%s] posted [%s], total posts %d:::: subreddit is [%s]\n",
			post.Author, post.Title, post.NumComments, post.Subreddit)
	}




	//mess := harvest.Messages[1]
	//fmt.Printf("[%s] wrote:: [%s]", mess.Author, mess.Body)

}