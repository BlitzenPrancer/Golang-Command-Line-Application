// this file will be managing videos like listing them and updating them
package main

// storing the videos as json, so i am importing json package
import (
	"encoding/json"
	"io/ioutil"
)

// designing what the video might look like
type video struct {
	Id          string
	Title       string
	Description string
	Imageurl    string
	Url         string
}

// function to make changes and save videos back to the file
func saveVideos(videos []video) {
	// converting it back to byte array using json.Marshal
	videoBytes, err := json.Marshal(videos)
	if err != nil {
		panic(err)
	}
	// writing it back to the file
	err = ioutil.WriteFile("./videos.json", videoBytes, 0644)
	if err != nil {
		panic(err)
	}
}

// function that will return all the videos from a file
func getVideos() (videos []video) {
	// reading the json file using ioutil package
	fileBytes, err := ioutil.ReadFile("./videos.json")
	// checking for errors
	if err != nil {
		panic(err)
	}
	// converting byte array(filebytes) into videos slice
	err = json.Unmarshal(fileBytes, &videos)
	// checking for errors
	if err != nil {
		panic(err)
	}
	return videos
}
