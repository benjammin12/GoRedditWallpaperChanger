package main

import (
	"fmt"
	"github.com/turnage/graw/reddit"
	"net/http"
	"log"
	"os"
	"io"
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
	/*
	for _, post := range harvest.Posts[:10] {
		fmt.Printf("[%s] posted [%s], total comments %d:::: subreddit is [%s]\n",
			post.Author, post.Title, post.NumComments, post.Subreddit)
	}
	*/

	//most recent post image
	imageURL := harvest.Posts[0].URL


	response, err := http.Get(imageURL) //get the image you just called

	if err != nil {
		log.Fatalln("Error retrieving image thumbnail", err)//try for second image
	}

	defer response.Body.Close()

	file, err := os.Create("pic/temp.jpg")

	if err!= nil {
		log.Fatalln("Error creating files", err)
	}
	defer file.Close()
	fmt.Println("Created file!")
	_, err = io.Copy(file, response.Body)
	if err != nil {
		log.Fatal(err)
	}

}
