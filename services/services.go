package services

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/delivercodes/bikemessenger/utils"
)

//RunService runs the docker image and outputs the cmd
func RunService() *exec.Cmd {
	config := utils.Readfile("data.yml")
	image := config.Service.Image

	KillService(image)

	name := "--name=" + image
	args := []string{"run", name, image}

	cmd := exec.Command("docker", args...)
	return cmd
}

//PullService ...
func PullService() {
	config := utils.Readfile("data.yml")
	image := config.Service.Image
	args := []string{"pull", image}
	out, err := exec.Command("docker", args...).Output()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	fmt.Printf("%s", out)
	KillService(image)
	cmd := RunService()
	cmd.Stdout = os.Stdout
	runErr := cmd.Start()
	if runErr != nil {
		log.Fatal(runErr)
		os.Exit(1)
	}

}

//CheckService ...
func CheckService() ([]byte, error) {
	args := []string{"--unix-socket", "/var/run/docker.sock", "http://localhost/containers/json"}
	out, err := exec.Command("curl", args...).Output()
	return out, err
}

//KillService kills a docker service, it takes the id of the container to kill
func KillService(container string) ([]byte, error) {
	killArgs := []string{"kill", container}
	rmArgs := []string{"rm", container}
	out, err := exec.Command("docker", killArgs...).Output()
	if out != nil {
		exec.Command("docker", rmArgs...).Run()
	}
	return out, err
}

//RestartService restarts the service .. holy shit dude
func RestartService(container string) *exec.Cmd {
	out, err := KillService(container)
	fmt.Printf("Restarting Service %s", out)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	cmd := RunService()
	return cmd
}
