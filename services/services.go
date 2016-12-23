package services

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/delivercodes/bikemessenger/models"
	yaml "gopkg.in/yaml.v2"
)

func readfile(file string) models.Config {
	t := models.Config{}
	dat, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	ymlerr := yaml.Unmarshal([]byte(dat), &t)
	if ymlerr != nil {
		log.Fatalf("error: %v", err)
	}
	return t
}

func execService(cmdName string, cmdArgs []string) {
	cmd := exec.Command(cmdName, cmdArgs...)
	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating StdoutPipe for Cmd", err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(cmdReader)
	go func() {
		for scanner.Scan() {
			fmt.Printf("%s\n", scanner.Text())
		}
	}()

	err = cmd.Start()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error starting Cmd", err)
		os.Exit(1)
	}

	err = cmd.Wait()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error waiting for Cmd", err)
		os.Exit(1)
	}
}

//
// func runService() {
//
// 	config := readfile("data.yml")
// 	image := config.Service.Image
// 	ports := ""
// 	envs := ""
// 	for _, port := range config.Service.Ports {
// 		ports += " -p " + port
// 	}
// 	for _, env := range config.Service.Env {
// 		envs += " -e " + env
// 	}
// 	pull := []string{"pull", image}
// 	run := []string{"run", image}
//
// 	execService("docker", run)
// 	execService("docker", pull)
// 	// fmt.Fprintf(w, "docker pull %v \n\n", image)
// 	// fmt.Fprintf(w, "docker run %v%v%v", image, ports, envs)
// }

func runService() {
	config := readfile("data.yml")
	image := config.Service.Image
	name := "--name=fuck"
	args := []string{"run", name, image}

	cmd := exec.Command("docker", args...)
	cmd.Stdout = os.Stdout
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Starting Docker instance on pid %d\n", cmd.Process.Pid)
}

//PullService ...
func PullService() {
	config := readfile("data.yml")
	image := config.Service.Image
	args := []string{"pull", image}
	out, err := exec.Command("docker", args...).Output()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	fmt.Printf("%s", out)
	runService()
}

//CheckService ...
func CheckService() []byte {
	args := []string{"--unix-socket", "/var/run/docker.sock", "http://localhost/containers/json"}
	out, err := exec.Command("curl", args...).Output()
	if err != nil {
		log.Fatal(err)
	}
	return out
}
