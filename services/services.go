package services

import (
	"fmt"
	"os/exec"

	"github.com/delivercodes/bikemessenger/models"
	"github.com/delivercodes/bikemessenger/utils"
)

//RunService runs the docker image and outputs the cmd
func RunService(service models.Service, name string) *exec.Cmd {
	image := service.Image
	args := []string{"run", "-d"}
	KillService(image)

	nameString := "--name=" + name
	args = append(args, nameString)
	fmt.Println(service.Ports)
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
func PullService(config models.Config) {
	for k := range config.Service {
		service := config.Service[k]

		args := []string{"pull", service.Image}
		out, _ := exec.Command("docker", args...).Output()
		fmt.Printf("%s", out)
		KillService(service.Image)
		RunService(service, k).Start()
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

	}
	exec.Command("docker", rmArgs...).Start()
	return out, err
}

//RestartService restarts the service .. holy shit dude
func RestartService(container string) *exec.Cmd {
	config, _ := utils.LoadConfigToModel(models.ConfigFile)
	out, _ := KillService(container)
	fmt.Printf("Restarting Service %s", out)

	cmd := RunService(config.Service[container], container)
	return cmd
}
