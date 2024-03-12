package knowledge

import (
	imageprocessing "demo/image_processing"
	"fmt"
	"image"
	"strings"
)

type Job1 struct {
	InputPath string
	Image     image.Image
	OutPath   string
}

func loadImage1(paths []string) <-chan Job1 {
	out := make(chan Job1)
	go func() {
		// For each input path create a job and add it to
		// the out channel
		for _, p := range paths {
			job := Job1{InputPath: p,
				OutPath: strings.Replace(p, "images/", "images/output/", 1)}
			job.Image = imageprocessing.ReadImage(p)
			out <- job
		}
		close(out)
	}()
	return out
}

func resize1(input <-chan Job1) <-chan Job1 {
	out := make(chan Job1)
	go func() {
		// For each input job, create a new job after resize and add it to
		// the out channel
		for job := range input { // Read from the channel
			job.Image = imageprocessing.Resize(job.Image)
			out <- job
		}
		close(out)
	}()
	return out
}

func convertToGrayscale(input <-chan Job1) <-chan Job1 {
	out := make(chan Job1)
	go func() {
		for job := range input { // Read from the channel
			job.Image = imageprocessing.Grayscale(job.Image)
			out <- job
		}
		close(out)
	}()
	return out
}

func saveImage(input <-chan Job1) <-chan bool {
	out := make(chan bool)
	go func() {
		for job := range input { // Read from the channel
			imageprocessing.WriteImage(job.OutPath, job.Image)
			out <- true
		}
		close(out)
	}()
	return out
}

// https://www.youtube.com/watch?v=8Rn8yOQH62k
func PipeLinePattern() {
	imagePaths := []string{"images/image1.jpg",
		"images/image2.jpg",
		"images/image3.jpg",
		"images/image4.jpg",
	}

	channel1 := loadImage1(imagePaths)
	channel2 := resize1(channel1)
	channel3 := convertToGrayscale(channel2)
	writeResults := saveImage(channel3)

	//The loop continues read from the writeResult channel
	//The loop automatically terminate when the writeResult channel closed in line 65 close(out)
	for success := range writeResults {
		if success {
			fmt.Println("PipeLinePattern Success!")
		} else {
			fmt.Println("PipeLinePattern Failed!")
		}
	}
}
