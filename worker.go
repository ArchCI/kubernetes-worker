package main

import (
	"fmt"
	"time"
)

// TODO: duplicated with simple-worker
type Task struct {
	Id      float64 `json:"id"`
	Commit  string  `json:"commit"`
	Public  bool    `json:"is_public"`
	Type    string  `json:"type"`
	Project string  `json:"project`
	Url     string  `json:"url"`
}

func main() {
	fmt.Println("Start kubernete worker")

	// Loop to pull task and run test
	for {

		task := Task{Id: 123, Commit: "commit", Public: true, Type: "github", Project: "test-project", Url: "https://github.com/tobegit3hub/test-project.git"}
		tasks := []Task{task}

		// If no task, sleep and wait for next
		if len(tasks) == 0 {
			fmt.Println("Sleep 5 seconds then pull task again")
			time.Sleep(5 * time.Second)
			continue
		}

		// TODO: goroutine to process in concurrent
		for _, task := range tasks {
			// 1. Build the image

			fmt.Println(task.Url)

			// 2. Build the pod file


			// 3. kubectl create -f xxx.json
		}

		// Sleep for next task
		time.Sleep(100 * time.Second)
	}


	fmt.Println("Kubernetes worker exists")
}
