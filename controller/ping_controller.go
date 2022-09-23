package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Book struct {
	TotalItem int `json:"totalItems"`
	Items     []Items
}

type Items struct {
	VolumeInfo VolumeInfo `json:"volumeInfo"`
}

type VolumeInfo struct {
	Title         string  `json:"title"`
	Publisher     string  `json:"publisher"`
	PublishedDate string  `json:"publishedDate"`
	AverageRating float32 `json:"AverageRating"`
}

func Ping(context *gin.Context) {
	book, err := getJson("https://www.googleapis.com/books/v7/volumes?q=harry+potter")
	if err != nil {
		log.Println(err)
	}
	context.JSON(http.StatusOK, gin.H{"data": book})
}

func getJson(url string) (Book, error) {
	client := http.Client{
		Timeout: time.Second * 5, // Timeout after 2 seconds
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Println(err)
		return Book{}, err
	}

	res, getErr := client.Do(req)
	if getErr != nil {
		log.Println(getErr)
		return Book{}, getErr
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Println(readErr)
		return Book{}, readErr
	}

	book := Book{}
	jsonErr := json.Unmarshal(body, &book)
	if jsonErr != nil {
		log.Println(jsonErr)
		return Book{}, jsonErr
	}
	return book, nil
}
