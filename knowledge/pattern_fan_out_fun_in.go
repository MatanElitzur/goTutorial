package knowledge

import (
	imageprocessing "demo/image_processing"
	"image"
	"strings"
)

type Job struct {
	InputPath string
	Image     image.Image
	OutPath   string
}

func loadImage(paths []string) []Job {
	var jobs []Job
	//for each input path create job struct and add it  to the list
	for _, p := range paths {
		job := Job{InputPath: p,
			OutPath: strings.Replace(p, "images/", "images/output/", 1)}
		job.Image = imageprocessing.ReadImage(p)
		jobs = append(jobs, job)
	}
	return jobs
}

func resize(jobs *[]Job) <-chan Job {
	out := make(chan Job, len(*jobs)) //The lenght of the channel buffer is like the lenght of the number of jobs
	for _, job := range *jobs {
		go func(job Job) {
			job.Image = imageprocessing.Resize(job.Image)
			out <- job //Place the job at the channel out
		}(job)
	}
	return out
}

// Collects the jobs back from the channel
func collectJobs(input <-chan Job, imageCnt int) []Job {
	var resizedJobs []Job
	for i := 0; i < imageCnt; i++ {
		job := <-input //Get the job from the channel
		resizedJobs = append(resizedJobs, job)
	}
	return resizedJobs
}

func saveImages(jobs *[]Job) {
	for _, job := range *jobs {
		imageprocessing.WriteImage(job.OutPath, job.Image)
	}
}

// https://www.youtube.com/watch?v=pQ2BCyFKKaY
func FunOutFunInPattern() {
	imagePaths := []string{"images/image1.jpg",
		"images/image2.jpg",
		"images/image3.jpg",
		"images/image4.jpg"}
	jobs := loadImage(imagePaths)
	//Fan out this function to multiple goroutines
	out := resize(&jobs)
	//Collect / Fan in
	resizedJobs := collectJobs(out, len(imagePaths))
	saveImages(&resizedJobs)
}
