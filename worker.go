package main

import (
	"fmt"
	"time"

	"io/ioutil"
	"os"
	"os/exec"

	"gopkg.in/yaml.v2"
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

type ArchciConfig struct {
	Image  string   `yaml:"image"`
	Script []string `yaml:"script"`
}

func ParseYaml(filename string) ArchciConfig {
	// fmt.Println("Start parse yaml") // TODO: Make it as debug log

	var archciConfig ArchciConfig
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(file, &archciConfig)
	if err != nil {
		panic(err)
	}

	// fmt.Printf("Value: %#v\n", config.Script[0])
	return archciConfig
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

			// 1. Clone the code in specified directory
			cloneCmd := exec.Command("git", "clone", task.Url)
			cloneOut, err := cloneCmd.Output()
			if err != nil {
				// TODO: Don't be so easy to exit
				os.Exit(1)
			}
			fmt.Println(string(cloneOut)) // Nothing to output if success
			fmt.Println("Success to clone the code")

			// 2. Parse archci.yaml file for base image and test scripts
			archciConfig := ParseYaml(task.Project + "/archci.yml")
			// fmt.Printf("Value: %#v\n", archciConfig.Image)
			dockerImage := archciConfig.Image
			fmt.Printf("Docker image: %#v\n", dockerImage)

			// 3. Generate archci.sh which will run test and push metrics to redis

			// 4. Docker run and mv the code into container's directory

			// 5. Docker commit with container name into a new docker image

			// 6. Generate pod.json for this task
			// Refer to http://stackoverflow.com/questions/28976455/start-kubernetes-container-with-specific-command

			// 7. kubectl create -f pod.json or request API

			// 8. Delete the code
			rmCmd := exec.Command("rm", "-rf", task.Project)
			rmOut, err := rmCmd.Output()
			if err != nil {
				// TODO: Don't be so easy to exit
				os.Exit(1)
			}
			fmt.Println(string(rmOut))
			fmt.Println("Success to delete the code")

			// Sleep for next task
			time.Sleep(100 * time.Second)
		}

		// Sleep for next task
		time.Sleep(100 * time.Second)
	}


	fmt.Println("Kubernetes worker exists")
}
