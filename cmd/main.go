package main

import (
	"flag"
	"fmt"
	"os"
)

// defining sub commands as well as arguments or flags that this application is going to accept.

func main() {
	// definition of videos get subcommand
	getCmd := flag.NewFlagSet("get", flag.ExitOnError)
	// input for "videos get" command
	getAll := getCmd.Bool("all", false, "Get all of the videos")
	getID := getCmd.String("id", "", "YouTube Video ID")

	// add command
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	// adding all the fields
	addId := addCmd.String("id", "", "Youtube video ID")
	addTitle := addCmd.String("title", "", "Youtube video Title")
	addUrl := addCmd.String("url", "", "Youtube vider URL")
	addImageUrl := addCmd.String("imageurl", "", "Youtube Video Image URL")
	addDesc := addCmd.String("desc", "", "Youtube video description")

	// doing validation of commands entered by user
	// checking the length of the argument to check the user had entered the sub command
	if len(os.Args) < 2 {
		fmt.Println("expected 'add' or 'get' subcommands")
		os.Exit(1)
	}

	// i am invoking those hander functions using switch statement
	switch os.Args[1] {
	case "get":
		HandleGet(getCmd, getAll, getID)
	case "add":
		HandleAdd(addCmd, addId, addTitle, addUrl, addImageUrl, addDesc)
	default:
	}
}

// handler function to handle ech of the commands
func HandleGet(getCmd *flag.FlagSet, all *bool, id *string) {
	// parsing the cli input
	getCmd.Parse(os.Args[2:])
	// checking the input and validating it
	if *all == false && *id == "" {
		fmt.Print("id is required or specify --all flag for showing all videos")
		// printing defaults
		getCmd.PrintDefaults()
		os.Exit(1)
	}

	// if the user enters all flag
	if *all {
		// returning all videos
		videos := getVideos()
		fmt.Printf("ID \t Title \t URL \t ImageURL \t Description \n")
		// for loop that ranges over the videos to return
		for _, video := range videos {
			fmt.Printf("%v \t %v \t %v \t %v \t %v \n", video.Id, video.Title, video.Url, video.Imageurl, video.Description)
		}
		return
	}

	// if the user enters id flag
	if *id != "" {
		// getting videos
		videos := getVideos()
		// reading the id passed in
		id := *id
		// looping through videos and checking if saved id is equal to passed id
		for _, video := range videos {
			if id == video.Id {
				fmt.Printf("ID \t Title \t URL \t ImageURL \t Description \n")
				// printing out the video data
				fmt.Printf("%v \t %v \t %v \t %v \t %v \n", video.Id, video.Title, video.Url, video.Imageurl, video.Description)
			}
		}
	}

}

// for get function
func HandleAdd(addCmd *flag.FlagSet, id *string, title *string, url *string, imageUrl *string, description *string) {
	ValidateVideo(addCmd, id, title, url, imageUrl, description)
	// to add the new video i am forming a new struct
	video := video{
		Id:          *id,
		Title:       *title,
		Description: *description,
		Imageurl:    *imageUrl,
		Url:         *url,
	}
	// appending new videos to existing videos
	videos := getVideos()
	videos = append(videos, video)
	// saving the videos back to file
	saveVideos(videos)
}

// validate video function
func ValidateVideo(addCmd *flag.FlagSet, id *string, title *string, url *string, imageUrl *string, description *string) {
	if *id == "" || *title == "" || *url == "" || *imageUrl == "" || *description == "" {
		fmt.Print("all fields are required for adding a video!!")
		addCmd.PrintDefaults()
		os.Exit(1)
	}
}
