package services

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/delivercodes/bikemessenger/models"
	"github.com/delivercodes/bikemessenger/utils"
)

//RunService runs the docker image and outputs the cmd
func RunService(service models.Service, name string) *exec.Cmd {
	image := service.Image
	args := []string{"run"}
	KillService(image)

	nameString := "--name=" + name
	args = append(args, nameString)
	//
	if service.Ports != nil {
		for _, port := range service.Ports {
			start := "-p"
			out := start + port
			args = append(args, out)
		}
	}

	if service.Env != nil {
		for _, env := range service.Env {
			out := "-e" + env
			args = append(args, out)
		}
	}

	args = append(args, image)
	cmd := exec.Command("docker", args...)
	fmt.Println(args)
	return cmd
}

//PullService ...
func PullService() {
	config := utils.LoadConfigToModel(models.ConfigFile())
	for k := range config.Service {
		service := config.Service[k]

		args := []string{"pull", service.Image}
		out, err := exec.Command("docker", args...).Output()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		fmt.Printf("%s", out)
		KillService(service.Image)
		runErr := RunService(service, k).Start()
		if runErr != nil {
			fmt.Printf("%s", runErr)
		}
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
	config := utils.LoadConfigToModel(models.ConfigFile())
	out, err := KillService(container)
	fmt.Printf("Restarting Service %s", out)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	cmd := RunService(config.Service[container], container)
	return cmd
}
