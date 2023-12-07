package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/BKoch74/Image_Ascii_Art/GetPic"
)

// Sequential version of the image downloader.
func downloadImagesSequential(urls []string) {
	for _, url := range urls {
		err := GetPic.Pic(url)
		if err != nil {
			fmt.Println("Error downloading:", url, "-", err)
		}
	}
}

// Concurrent version of the image downloader.
func downloadImagesConcurrent(urls []string) {
	var (
		done   = make(chan struct{})
		errors = make(chan error)
	)

	for _, url := range urls {
		go func(url string) {
			errors <- GetPic.Pic(url)
			done <- struct{}{}
		}(url)
	}

	// Wait for all downloads to complete
	for i := 0; i < len(urls); i++ {
		select {
		case <-done:
		case err := <-errors:
			if err != nil {
				fmt.Println("Error downloading:", err)
			}
		}
	}
}

func extractFilename(url string) string {
	fileParts := strings.Split(url, "/")
	return fileParts[len(fileParts)-1]
}

func main() {
	urls := []string{
		"https://www.pexels.com/photo/aigle-19361893/",
		"https://www.pexels.com/photo/black-and-white-photo-of-a-snowy-mountain-19329508/",
		"https://pixabay.com/photos/cow-highland-cattle-highland-cow-8349729/",
		"https://pixabay.com/illustrations/polar-bear-snow-christmas-ice-8416167/",
		"https://pixabay.com/photos/mountain-landscape-nature-outdoors-8380938/"}
	// Add more image URLs

	// Sequential download
	start := time.Now()
	downloadImagesSequential(urls)
	fmt.Printf("Sequential download took: %v\n", time.Since(start))

	// Concurrent download
	start = time.Now()
	downloadImagesConcurrent(urls)
	fmt.Printf("Concurrent download took: %v\n", time.Since(start))
}

// Helper function to download and save a single image.
func downloadImage(url, filename string) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}
