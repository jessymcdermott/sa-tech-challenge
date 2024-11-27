package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// ImageData is a struct to map to list of Images available
type ImageData struct {
	Images []string `json:"images"`
}

// ImageUrl is a struct to map the JSON output
type ImageUrl struct {
	ImageUrl string `json:"imageUrl"`
}

var bucketName string
var imageUrls []string

// init is special function that gets called before main
func init() {
	// Open the JSON file
	file, err := os.Open("images.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Decode the JSON data
	var data ImageData
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&data)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	// Get the bucket name from the environment variable, with a default value
	bucketName = os.Getenv("BUCKET_NAME")
	if bucketName == "" {
		bucketName = "random-pictures"
	}

	for _, image := range data.Images {
		url := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", bucketName, image)
		imageUrls = append(imageUrls, url)
	}
}

func main() {

	// create a new echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Health check endpoint
	e.GET("/health", healthCheckHandler)

	// define a route '/imageUrl'
	e.GET("/imageUrl", imageUrlHandler)

	// start the server on the specified port
	e.Logger.Fatal(e.Start(":10116"))
}

func imageUrlHandler(c echo.Context) error {

	// select a random image url
	randomIndex := rand.Intn(len(imageUrls))
	selectedUrl := imageUrls[randomIndex]

	// create a image url struct with the selected image url
	response := ImageUrl{ImageUrl: selectedUrl}

	// return the response
	return c.JSON(http.StatusOK, response)
}

func healthCheckHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"status": "healthy"})
}
